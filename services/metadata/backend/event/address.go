package event

import "github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"

// implementation of marker interface to make sure that event structs can only
// come from this package. Go way of implementing a class hierarchy.
func (e AddressCreated) isEvent() {}
func (e AddressChanged) isEvent() {}
func (e AddressDeleted) isEvent() {}

type AddressCreated struct {
	ID         valueobject.Identifier `json:"id"`
	Type       string                 `json:"type"`
	Street     valueobject.Street     `json:"street"`
	PostalCode valueobject.PostalCode `json:"postalCode"`
	Locality   valueobject.Locality   `json:"locality"`
	Country    valueobject.Country    `json:"country"`
	Canton     valueobject.Canton     `json:"canton"`
	Additional valueobject.Additional `json:"additional"`
	CreatedAt  valueobject.Timestamp  `json:"createdAt"`
	CreatedBy  valueobject.Identifier `json:"createdBy"`
}

type AddressChanged struct {
	ID         valueobject.Identifier `json:"id"`
	Type       string                 `json:"type"`
	Street     valueobject.Street     `json:"street"`
	PostalCode valueobject.PostalCode `json:"postalCode"`
	Locality   valueobject.Locality   `json:"locality"`
	Country    valueobject.Country    `json:"country"`
	Canton     valueobject.Canton     `json:"canton"`
	Additional valueobject.Additional `json:"additional"`
	ChangedAt  valueobject.Timestamp  `json:"changedAt"`
	ChangedBy  valueobject.Identifier `json:"changedBy"`
}

type AddressDeleted struct {
	ID        valueobject.Identifier `json:"id"`
	DeletedAt valueobject.Timestamp  `json:"deletedAt"`
	DeletedBy valueobject.Identifier `json:"deletedBy"`
}
