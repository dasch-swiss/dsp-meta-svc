package address

import (
	"context"
	addressEntity "github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/entity/address"

	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
)

type Reader interface {
	Load(ctx context.Context, id valueobject.Identifier) (*addressEntity.Address, error)
	GetAddressIds(ctx context.Context, includeDeleted bool) ([]valueobject.Identifier, error)
}

type Writer interface {
	Save(ctx context.Context, e *addressEntity.Address) (valueobject.Identifier, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetAddress(ctx context.Context, id valueobject.Identifier) (*addressEntity.Address, error)
	GetAddresses(ctx context.Context, includeDeleted bool) ([]addressEntity.Address, error)
	CreateAddress(ctx context.Context, s valueobject.Street, pc valueobject.PostalCode, l valueobject.Locality, c valueobject.Country, ca valueobject.Canton, a valueobject.Additional) (valueobject.Identifier, error)
	UpdateAddress(ctx context.Context, id valueobject.Identifier, s valueobject.Street, pc valueobject.PostalCode, l valueobject.Locality, c valueobject.Country, ca valueobject.Canton, a valueobject.Additional) (*addressEntity.Address, error)
	DeleteAddress(ctx context.Context, id valueobject.Identifier) (*addressEntity.Address, error)
}
