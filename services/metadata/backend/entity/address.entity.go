package main

import (
	// "errors"
	"fmt"

	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	// "github.com/opencontainers/runc/libcontainer/configs/validate"
)

const addresType = "http://ns.dasch.swiss/repository#Address"

type Address struct {
	ID         valueobject.Identifier  `json:"id"`
	Type       string                  `json:"type"`
	Street     valueobject.Street      `json:"street"`
	PostalCode valueobject.PostalCode  `json:"postalCode"`
	Locality   valueobject.Locality    `json:"locality"`
	Country    valueobject.Country     `json:"country"`
	Canton     *valueobject.Canton     `json:"canton"`
	Additional *valueobject.Additional `json:"additional"`
	CreatedAt  valueobject.Timestamp   `json:"createdAt"`
	CreatedBy  valueobject.Identifier  `json:"createdBy"`
	ChangedAt  valueobject.Timestamp   `json:"changedAt,omitempty"`
	ChangedBy  valueobject.Identifier  `json:"changedBy,omitempty"`
	DeletedAt  valueobject.Timestamp   `json:"deletedAt,omitempty"`
	DeletedBy  valueobject.Identifier  `json:"deletedBy,omitempty"`
}

func NewAddress(Street valueobject.Street, PostalCode valueobject.PostalCode, Locality valueobject.Locality, Country valueobject.Country, Canton *valueobject.Canton, Additional *valueobject.Additional) (*Address, error) {
	a := &Address{
		ID:         valueobject.Identifier{},
		Type:       addresType,
		Street:     Street,
		PostalCode: PostalCode,
		Locality:   Locality,
		Country:    Country,
		Canton:     Canton,
		Additional: Additional,
		CreatedAt:  valueobject.NewTimestamp(),
		CreatedBy:  valueobject.Identifier{},
		ChangedAt:  valueobject.NewTimestamp(),
		ChangedBy:  valueobject.Identifier{},
		DeletedAt:  valueobject.NewTimestamp(),
		DeletedBy:  valueobject.Identifier{},
	}

	// if a.Street == "" || a.PostalCode == "" || a.Locality == "" || a.Country == "" {
	// 	fmt.Print("Invalid address: street, postalCode, locality and country are required")
	// 	return nil, errors.New("Invalid address: street, postalCode, locality and country are required")
	// }

	fmt.Println(a)

	return a, nil
}

func (a *Address) AddType() error {
	a.Type = addresType
	a.ChangedAt = valueobject.NewTimestamp()
	// a.ChangedBy = valueobject.Identifier{}
	return nil
}

func (a *Address) AddStreet(s valueobject.Street) error {
	a.Street = s
	a.ChangedAt = valueobject.NewTimestamp()
	// a.ChangedBy = valueobject.Identifier{}
	return nil
}

func (a *Address) AddPostalCode(pc valueobject.PostalCode) error {
	a.PostalCode = pc
	a.ChangedAt = valueobject.NewTimestamp()
	// a.ChangedBy = valueobject.Identifier{}
	return nil
}

func (a *Address) AddLocality(l valueobject.Locality) error {
	a.Locality = l
	a.ChangedAt = valueobject.NewTimestamp()
	// a.ChangedBy = valueobject.Identifier{}
	return nil
}

func (a *Address) AddCountry(c valueobject.Country) error {
	a.Country = c
	a.ChangedAt = valueobject.NewTimestamp()
	// a.ChangedBy = valueobject.Identifier{}
	return nil
}

func (a *Address) AddCanton(c valueobject.Canton) error {
	a.Canton = &c
	a.ChangedAt = valueobject.NewTimestamp()
	// a.ChangedBy = valueobject.Identifier{}
	return nil
}

func (a *Address) AddAdditional(add valueobject.Additional) error {
	a.Additional = &add
	a.ChangedAt = valueobject.NewTimestamp()
	// a.ChangedBy = valueobject.Identifier{}
	return nil
}

func main() {
	street, _ := valueobject.NewStreet("street")
	code, _ := valueobject.NewPostalCode("00-000")
	locality, _ := valueobject.NewLocality("Poznan")
	coauntry, _ := valueobject.NewCountry("Poland")
	// canton, _ := valueobject.NewCanton("Basel-Stadt")
	// additional, _ := valueobject.NewAdditional("blablabla")
	NewAddress(street, code, locality, coauntry, nil, nil)
}
