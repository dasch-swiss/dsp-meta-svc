package presenter

import "github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"

// data used as the result for any address operation
type Address struct {
	ID         valueobject.Identifier `json:"id"`
	Type       string                 `json:"type"`
	Street     string                 `json:"street"`
	PostalCode string                 `json:"postalCode"`
	Locality   string                 `json:"locality"`
	Country    string                 `json:"country"`
	Canton     string                 `json:"canton"`
	Additional string                 `json:"additional"`
	CreatedAt  string                 `json:"createdAt"`
	CreatedBy  string                 `json:"createdBy"`
	ChangedAt  string                 `json:"changedAt"`
	ChangedBy  string                 `json:"changedBy"`
	DeletedAt  string                 `json:"deletedAt"`
	DeletedBy  string                 `json:"deletedBy"`
}

func (a *Address) NullifyJsonProps() Address {
	if a.ChangedAt == "0001-01-01 00:00:00 +0000 UTC" {
		a.ChangedAt = ""
	}

	if a.ChangedBy == "00000000-0000-0000-0000-000000000000" {
		a.ChangedBy = ""
	}

	if a.DeletedAt == "0001-01-01 00:00:00 +0000 UTC" {
		a.DeletedAt = ""
	}

	if a.DeletedBy == "00000000-0000-0000-0000-000000000000" {
		a.DeletedBy = ""
	}

	return *a
}
