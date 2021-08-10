package repository

import (
	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/entity/address"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
)

type inmemDB struct {
	m map[valueobject.Identifier]*address.Address
}

func NewInmemDB() *inmemDB {
	var m = map[valueobject.Identifier]*address.Address{}
	return &inmemDB{
		m: m,
	}
}

func (r *inmemDB) Create(e *address.Address) (valueobject.Identifier, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

func (r *inmemDB) Get(id valueobject.Identifier) (*address.Address, error) {
	if r.m[id] == nil {
		return nil, address.ErrAddressNotFound
	}
	return r.m[id], nil
}

