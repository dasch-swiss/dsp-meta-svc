package address_test

import (
	"context"
	addressEntity "github.com/dasch-swiss/dsp-meta-svc/services/metadata/go_ca_backend/entity/address"

	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/go_ca_backend/event"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/gofrs/uuid"
)

// in memory repository visible only inside this package
type inMemRepo struct {
	m map[uuid.UUID][]event.Event
}

// creates new in memory repsitory
func NewInMemRepo() *inMemRepo {
	var m = map[uuid.UUID][]event.Event{}

	return &inMemRepo{
		m: m,
	}
}

// saves address
func (r *inMemRepo) Save(ctx context.Context, e *addressEntity.Address) (valueobject.Identifier, error) {
	var events []event.Event

	// get previously stored events
	if r.m[e.ID.UUID()] != nil {
		events = append(events, r.m[e.ID.UUID()]...)
	}

	// append new events
	events = append(events, e.Events()...)

	// store new events
	r.m[e.ID.UUID()] = events

	return e.ID, nil
}

// loads address
func (r *inMemRepo) Load(ctx context.Context, id valueobject.Identifier) (*addressEntity.Address, error) {
	if r.m[id.UUID()] == nil {
		return nil, addressEntity.ErrAddressNotFound
	}

	return addressEntity.NewAddressFromEvents(r.m[id.UUID()]), nil

}

func (r *inMemRepo) GetAddressIds(ctx context.Context, includeDeleted bool) ([]valueobject.Identifier, error) {
	i := 0
	addressIds := make([]valueobject.Identifier, len(r.m))

	for id := range r.m {
		// create mepty identyfier
		uuid := valueobject.Identifier{}
		// create byte array from provided id string
		b := []byte(id.String())
		// assign value of uuid
		uuid.UnmarshalText(b)
		addressIds[i] = uuid
		i++
	}

	return addressIds, nil
}
