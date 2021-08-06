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

func (p *Address) NullifyJsonProps() Address {
	if p.ChangedAt == "0001-01-01 00:00:00 +0000 UTC" {
		p.ChangedAt = "null"
	}

	if p.ChangedBy == "00000000-0000-0000-0000-000000000000" {
		p.ChangedBy = "null"
	}

	if p.DeletedAt == "0001-01-01 00:00:00 +0000 UTC" {
		p.DeletedAt = "null"
	}

	if p.DeletedBy == "00000000-0000-0000-0000-000000000000" {
		p.DeletedBy = "null"
	}

	return *p
}
