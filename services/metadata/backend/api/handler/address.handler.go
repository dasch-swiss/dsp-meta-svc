package handler

import (
	"encoding/json"
	"net/http"

	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/service/address"
)

// reusable struct to use when decoding the JSON request body
type AddressRequestBody struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Street     string `json:"street"`
	PostalCode string `json:"postalCode"`
	Locality   string `json:"locality"`
	Country    string `json:"country"`
	Canton     string `json:"canton"`
	Additional string `json:"additional"`
}

// creates addres with provided request body
func createAddress(service address.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// check JWT token to make sure user is authenticated
		// an object containing the users info is returned by ExtractTokenMetadata (currently not used, hence the underscore)
		_, tokenErr := ExtractTokenMetadata(r)
		if tokenErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(tokenErr.Error()))
			return
		}

		var input AddressRequestBody
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
	}
}
