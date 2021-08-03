package address

import (
	// "fmt"
	"log"

	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/event"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
)

const addressType = "http://ns.dasch.swiss/repository#Address"

// TODO canton and additonal should be optional
type Address struct {
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
	ChangedAt  valueobject.Timestamp  `json:"changedAt,omitempty"`
	ChangedBy  valueobject.Identifier `json:"changedBy,omitempty"`
	DeletedAt  valueobject.Timestamp  `json:"deletedAt,omitempty"`
	DeletedBy  valueobject.Identifier `json:"deletedBy,omitempty"`

	Changes []event.Event
}

// creates a new address entity
func NewAddress(id valueobject.Identifier, street valueobject.Street, postalCode valueobject.PostalCode, locality valueobject.Locality, country valueobject.Country, canton valueobject.Canton, additional valueobject.Additional) *Address {
	a := &Address{}

	a.raise(&event.AddressCreated{
		ID:         id,
		Type:       addressType,
		Street:     street,
		PostalCode: postalCode,
		Locality:   locality,
		Country:    country,
		Canton:     canton,
		Additional: additional,
		CreatedAt:  valueobject.NewTimestamp(),
		CreatedBy:  valueobject.Identifier{},
	})

	return a
}

// updates an address entity
func (a *Address) UpdateAddress(id valueobject.Identifier, street valueobject.Street, postalCode valueobject.PostalCode, locality valueobject.Locality, country valueobject.Country, canton valueobject.Canton, additional valueobject.Additional) error {
	if !a.DeletedAt.Time().IsZero() {
		return ErrAddressHasBeenDeleted
	}
	a.raise(&event.AddressChanged{
		ID:         id,
		Type:       addressType,
		Street:     street,
		PostalCode: postalCode,
		Locality:   locality,
		Country:    country,
		Canton:     canton,
		Additional: additional,
		ChangedAt:  valueobject.NewTimestamp(),
		ChangedBy:  valueobject.Identifier{},
	})

	return nil
}

// deletes an address entity
func (a *Address) DeleteAddress(id valueobject.Identifier) error {
	if !a.DeletedAt.Time().IsZero() {
		return ErrAddressHasBeenDeleted
	}

	a.raise(&event.AddressDeleted{
		ID:        id,
		DeletedAt: valueobject.NewTimestamp(),
		DeletedBy: valueobject.Identifier{},
	})

	return nil
}

// func main() {
// 	street, _ := valueobject.NewStreet("street")
// 	code, _ := valueobject.NewPostalCode("00-000")
// 	locality, _ := valueobject.NewLocality("Poznan")
// 	coauntry, _ := valueobject.NewCountry("Poland")
// 	// canton, _ := valueobject.NewCanton("Basel-Stadt")
// 	// additional, _ := valueobject.NewAdditional("blablabla")
// 	NewAddress(street, code, locality, coauntry, nil, nil)
// }

// The raise method does two things, it appends the event into our changes slice
// and calls the event handler On saying that this is a new event and we should
// not increment the version number. The version is an optimistic concurrency
// pattern used to help us avoid database locks to change our aggregate.
func (a *Address) raise(event event.Event) {
	a.Changes = append(a.Changes, event)
	a.On(event, true)
}

// On handles user events on the project aggregate.
// The On method first does a type switch on the event and selects the case for
// each event type. This is where state change happens. Once an event is emitted
// and saved we do not throw an error, we simply process the event and carry on.
// We can change here if we decide that an event is no longer relevant or if it
// means something different, but we can’t return an error and say an event is
// invalid. Then we check if this is a new event. If it isn’t we increment the
// version number of our aggregate.
func (a *Address) On(ev event.Event, new bool) {
	switch e := ev.(type) {
	case *event.AddressCreated:
		a.ID = e.ID
		a.Type = addressType
		a.Street = e.Street
		a.PostalCode = e.PostalCode
		a.Locality = e.Locality
		a.Country = e.Country
		a.Canton = e.Canton
		a.Additional = e.Additional
		a.CreatedAt = e.CreatedAt
		a.CreatedBy = e.CreatedBy

	case *event.AddressChanged:
		a.ID = e.ID
		a.Type = addressType
		a.Street = e.Street
		a.PostalCode = e.PostalCode
		a.Locality = e.Locality
		a.Country = e.Country
		a.Country = e.Country
		a.Canton = e.Canton
		a.ChangedAt = e.ChangedAt
		a.ChangedBy = e.ChangedBy

	case *event.AddressDeleted:
		a.ID = e.ID
		a.DeletedAt = e.DeletedAt
		a.DeletedBy = e.DeletedBy

	default:
		log.Printf("unknown event %T", e)
	}
}

// helper method that creates a new address from a series of events
func NewAddressFromEvents(e []event.Event) *Address {
	a := &Address{}

	for _, err := range e {
		a.On(err, false)
	}

	return a
}