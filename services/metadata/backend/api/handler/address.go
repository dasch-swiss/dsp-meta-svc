package handler

import (
	"context"
	"encoding/json"
	addressEntity "github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/entity/address"
	"github.com/gorilla/mux"
	"log"
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

// makes URL handlers for creating, updating, deleting and getting addresses
func HandleAddressRoutes(r *mux.Router, service address.UseCase)  {
	r.HandleFunc("/v1/addresses", createAddress(service)).Methods("POST", "OPTIONS")
	r.HandleFunc("/v1/addresses/{id}", updateAddress(service)).Methods("PUT", "OPTION")
	r.HandleFunc("/v1/addresses/{id}", deleteAddress(service)).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/v1/addresses/{id}", getAddress(service)).Methods("GET", "OPTION")
	r.HandleFunc("/v1/addresses", getAddresses(service)).Methods("GET", "OPTIONS")
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

		// convert input strings to value object
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

// updates address with the provided RequestBody.
// Updating address that has been marked as deleted is not possible.
// All fields of the RequestBody must be provided.
// At least one of the values of the provided RequestBody must differ from the current value of the corresponding address field.
// If a value of a field is identical to what it already is, the update will not be performed for that field.
func updateAddress(service address.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var input AddressRequestBody
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		// get variables from request url
		vars := mux.Vars(r)
		// create empty Identifier
		uuid := valueobject.Identifier{}
		// create byte array from provided id string
		b := []byte(vars["id"])
		// assign value of Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get address
		a, err := service.GetAddress(ctx, uuid)
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
		if a != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(addressEntity.ErrNoAddressDataReturned.Error()))
			return
		}

		// convert input strings to value object
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

		// update address
		updatedAddress, err := service.UpdateAddress(ctx, uuid, s, pc, l, c, canton, add)
		if err != nil && err == addressEntity.ErrAddressHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err == addressEntity.ErrNoPropertiesChanged {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(addressEntity.ErrServerNotResponding.Error()))
		}

		res := &presenter.Address{
			ID:         uuid,
			Type:       ADDRESS_TYPE,
			Street:     updatedAddress.Street.String(),
			PostalCode: updatedAddress.PostalCode.String(),
			Locality:   updatedAddress.Locality.String(),
			Country:    updatedAddress.Country.String(),
			Canton:     updatedAddress.Canton.String(),
			Additional: updatedAddress.Additional.String(),
			CreatedAt:  updatedAddress.CreatedAt.String(),
			CreatedBy:  updatedAddress.CreatedBy.String(),
			ChangedAt:  updatedAddress.ChangedAt.String(),
			ChangedBy:  updatedAddress.ChangedBy.String(),
			DeletedAt:  updatedAddress.DeletedAt.String(),
			DeletedBy:  updatedAddress.DeletedBy.String(),
		}

		// replace null-values with "null"
		*res = res.NullifyJsonProps()
		
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

// deletes address with provided UUID
func deleteAddress(service address.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// get variables from request url
		vars := mux.Vars(r)
		// create empty Identifier
		uuid := valueobject.Identifier{}
		// create byte array from provided id string
		b := []byte(vars["id"])
		// assign value of Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// delete address
		deletedAddress, err := service.DeleteAddress(ctx, uuid)
		w.Header().Set("Content-Type", "application/json")
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
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(addressEntity.ErrServerNotResponding.Error()))
			return
		}
		if deletedAddress == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(addressEntity.ErrNoAddressDataReturned.Error()))
			return
		}

		res := &presenter.Address{
			ID:         uuid,
			Type:       ADDRESS_TYPE,
			Street:     deletedAddress.Street.String(),
			PostalCode: deletedAddress.PostalCode.String(),
			Locality:   deletedAddress.Locality.String(),
			Country:    deletedAddress.Country.String(),
			Canton:     deletedAddress.Canton.String(),
			Additional: deletedAddress.Additional.String(),
			CreatedAt:  deletedAddress.CreatedAt.String(),
			CreatedBy:  deletedAddress.CreatedBy.String(),
			ChangedAt:  deletedAddress.ChangedAt.String(),
			ChangedBy:  deletedAddress.ChangedBy.String(),
			DeletedAt:  deletedAddress.DeletedAt.String(),
			DeletedBy:  deletedAddress.DeletedBy.String(),
		}

		// replace null-values with "null"
		*res = res.NullifyJsonProps()

		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

func getAddress(service address.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// get variables from request url
		vars := mux.Vars(r)
		uuid, err := valueobject.IdentifierFromBytes([]byte(vars["id"]))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get address by UUID
		a, err := service.GetAddress(ctx, uuid)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err == addressEntity.ErrAddressNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err != addressEntity.ErrAddressNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if a == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(addressEntity.ErrNoAddressDataReturned.Error()))
			return
		}

		res := &presenter.Address{
			ID:         uuid,
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

		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

// gets list of all addresses
// By default, this only returns active (not marked as delete) addresses.
// IncludeDeleted can be provided in the request body to also return addresses marked as deleted.
func getAddresses(service address.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct{
			IncludeDeleted bool `json:"includeDeleted"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			// default to false if decoding fails (likely because it wasn't provided)
			input.IncludeDeleted = false
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get all addresses
		addresses, err := service.GetAddresses(ctx, input.IncludeDeleted)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err == addressEntity.ErrAddressNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err != addressEntity.ErrAddressNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(addressEntity.ErrServerNotResponding.Error()))
			return
		}
		if addresses == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(addressEntity.ErrNoAddressDataReturned.Error()))
			return
		}

		var res []presenter.Address
		for _, a := range addresses {
			addressToAppend := presenter.Address{
				ID:         a.ID,
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
			addressToAppend = addressToAppend.NullifyJsonProps()
			res = append(res, addressToAppend)
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

