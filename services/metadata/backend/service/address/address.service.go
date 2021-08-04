package address

import (
	"context"

	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/entity"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
)

// implements Repository interface
type Service struct {
	repo Repository
}

// creates new address usecase/sercice
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// creates new address
func (s *Service) CreateAddress(ctx context.Context, street valueobject.Street, postalCode valueobject.PostalCode, locality valueobject.Locality, country valueobject.Country, canton valueobject.Canton, additional valueobject.Additional) (valueobject.Identifier, error) {
	// generate new UUID
	id, _ := valueobject.NewIdentifier()

	// populate data
	a := address.NewAddress(id, street, postalCode, locality, country, canton, additional)

	if _, err := s.repo.Save(ctx, a); err != nil {
		return valueobject.Identifier{}, err
	}

	return id, nil
}

// updates existing address with provided values
func (s *Service) UpdateAddress(ctx context.Context, id valueobject.Identifier, street valueobject.Street, postalCode valueobject.PostalCode, locality valueobject.Locality, country valueobject.Country, canton valueobject.Canton, additional valueobject.Additional) (*address.Address, error) {

	// get address to update
	a, err := s.repo.Load(ctx, id)
	if err != nil {
		return &address.Address{}, err
	}

	// throw error if address has been deleted
	if !a.DeletedAt.Time().IsZero() {
		return &address.Address{}, address.ErrAddressHasBeenDeleted
	}

	// throw error if none of passed values differ from current ones
	if isIdentical(*a, street, postalCode, locality, country, canton, additional) {
		return &address.Address{}, address.ErrNoPropertiesChanged
	}

	// update address
	if err := a.UpdateAddress(id, street, postalCode, locality, country, canton, additional); err != nil {
		return &address.Address{}, err
	}

	// save event
	if _, err := s.repo.Save(ctx, a); err != nil {
		return &address.Address{}, err
	}

	return a, nil
}

func (s *Service) DeleteAddress(ctx context.Context, id valueobject.Identifier) (*address.Address, error) {
	// get address to delete
	a, err := s.repo.Load(ctx, id)
	if err != nil {
		return &address.Address{}, err
	}

	// delete address
	a.DeleteAddress(id)

	// save event
	if _, err := s.repo.Save(ctx, a); err != nil {
		return &address.Address{}, err
	}

	return a, nil
}

func (s *Service) GetAddress(ctx context.Context, id valueobject.Identifier) (*address.Address, error) {
	a, err := s.repo.Load(ctx, id)
	if err != nil {
		return &address.Address{}, err
	}
	return a, err
}

func (s *Service) GetAddresses(ctx context.Context, includeDeletedAddresses bool) ([]address.Address, error) {
	var addresses []address.Address
	ids, err := s.repo.GetAddressIds(ctx, includeDeletedAddresses)
	if err != nil {
		return []address.Address{}, err
	}

	for _, id := range ids {
		a, err := s.GetAddress(ctx, id)
		if err != nil {
			return []address.Address{}, err
		}

		addresses = append(addresses, *a)
	}

	return addresses, nil

}

func isIdentical(a address.Address, street valueobject.Street, postalCode valueobject.PostalCode, locality valueobject.Locality, country valueobject.Country, canton valueobject.Canton, additional valueobject.Additional) bool {
	if a.Street.Equals(street) &&
		a.PostalCode.Equals(postalCode) &&
		a.Locality.Equals(locality) &&
		a.Country.Equals(country) &&
		a.Canton.Equals(canton) &&
		a.Additional.Equals(additional) {
		return true
	}

	return false
}
