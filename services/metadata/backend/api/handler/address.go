package handler

import (
	"context"
	"encoding/json"
	addressEntity "github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/entity/address"
	"net/http"
	"time"

	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/api/presenter"
	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/service/address"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
)

const ADDRESS_TYPE = "http://ns.dasch.swiss/repository#Address"

// reusable struct to use when decoding the JSON request body
type AddressRequestBody struct {
	// ID         string `json:"id"`
	// Type       string `json:"type"`
	Street     string `json:"street"`
	PostalCode string `json:"postalCode"`
	Locality   string `json:"locality"`
	Country    string `json:"country"`
	Canton     string `json:"canton"`
	Additional string `json:"additional"`
}

// creates address with provided request body
func createAddress(service address.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var input AddressRequestBody
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// convert input to value object
		s, err := valueobject.NewStreet(input.Street)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		pc, err := valueobject.NewPostalCode(input.PostalCode)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		l, err := valueobject.NewLocality(input.Locality)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		c, err := valueobject.NewCountry(input.Country)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		canton, err := valueobject.NewCanton(input.Canton)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		add, err := valueobject.NewAdditional(input.Additional)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// create address
		id, err := service.CreateAddress(ctx, s, pc, l, c, canton, add)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// get address
		a, err := service.GetAddress(ctx, id)
		if err != nil && err == addressEntity.ErrAddressNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err == addressEntity.ErrAddressHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err != addressEntity.ErrAddressNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(addressEntity.ErrServerNotResponding.Error()))
			return
		}
		if a == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(addressEntity.ErrNoAddressDataReturned.Error()))
			return
		}

		res := &presenter.Address{
			ID:         id,
			Type:       ADDRESS_TYPE,
			Street:     a.Street.String(),
			PostalCode: a.PostalCode.String(),
			Locality:   a.Locality.String(),
			Country:    a.Country.String(),
			Canton:     a.Canton.String(),
			Additional: a.Additional.String(),
			CreatedAt:  a.CreatedAt.String(),
			CreatedBy:  a.CreatedBy.String(),
			ChangedAt:  a.ChangedAt.String(),
			ChangedBy:  a.ChangedBy.String(),
			DeletedAt:  a.DeletedAt.String(),
			DeletedBy:  a.DeletedBy.String(),
		}

		// replace null-values with "null"
		*res = res.NullifyJsonProps()

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
