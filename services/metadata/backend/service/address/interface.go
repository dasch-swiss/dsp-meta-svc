package address

import (
	"context"

	address "github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/entity"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
)

type Reader interface {
	Load(ctx context.Context, id valueobject.Identifier) (*address.Address, error)
	GetAddressIds(ctx context.Context, includeDeletedAddresses bool) ([]valueobject.Identifier, error)
}

type Writer interface {
	Save(ctx context.Context, e *address.Address) (valueobject.Identifier, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetAddress(ctx context.Context, id valueobject.Identifier) (*address.Address, error)
	GetAddresses(ctx context.Context, includeDeletedAddresses bool) ([]address.Address, error)
	CreateAddress(ctx context.Context, s valueobject.Street, pc valueobject.PostalCode, l valueobject.Locality, c valueobject.Country, ca valueobject.Canton, a valueobject.Additional) (valueobject.Identifier, error)
	UpdateAddress(ctx context.Context, id valueobject.Identifier, s valueobject.Street, pc valueobject.PostalCode, l valueobject.Locality, c valueobject.Country, ca valueobject.Canton, a valueobject.Additional) (*address.Address, error)
	DeleteAddress(ctx context.Context, id valueobject.Identifier) (*address.Address, error)
}
