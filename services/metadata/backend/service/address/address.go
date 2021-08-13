package address

import (
	"context"
	addressEntity "github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/entity/address"

	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
)

// implements Repository interface
type Service struct {
	repo Repository
}

// NewService creates new address service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateAddress creates new address
func (s *Service) CreateAddress(ctx context.Context, street valueobject.Street, postalCode valueobject.PostalCode, locality valueobject.Locality, country valueobject.Country, canton valueobject.Canton, additional valueobject.Additional) (valueobject.Identifier, error) {
	// generate new UUID
	id, _ := valueobject.NewIdentifier()

	// populate data
	a := addressEntity.NewAddress(id, street, postalCode, locality, country, canton, additional)

	if _, err := s.repo.Save(ctx, a); err != nil {
		return valueobject.Identifier{}, err
	}

	return id, nil
}

// UpdateAddress updates existing address with provided values
func (s *Service) UpdateAddress(ctx context.Context, id valueobject.Identifier, street valueobject.Street, postalCode valueobject.PostalCode, locality valueobject.Locality, country valueobject.Country, canton valueobject.Canton, additional valueobject.Additional) (*addressEntity.Address, error) {

	// get address to update
	a, err := s.repo.Load(ctx, id)
	if err != nil {
		return &addressEntity.Address{}, err
	}

	// throw error if address has been deleted
	if !a.DeletedAt.Time().IsZero() {
		return &addressEntity.Address{}, addressEntity.ErrAddressHasBeenDeleted
	}

	// throw error if none of passed values differ from current ones
	if isIdentical(*a, street, postalCode, locality, country, canton, additional) {
		return &addressEntity.Address{}, addressEntity.ErrNoPropertiesChanged
	}

	// update address
	if err := a.UpdateAddress(id, street, postalCode, locality, country, canton, additional); err != nil {
		return &addressEntity.Address{}, err
	}

	// save event
	if _, err := s.repo.Save(ctx, a); err != nil {
		return &addressEntity.Address{}, err
	}

	return a, nil
}

// DeleteAddress deletes address of provided ID
func (s *Service) DeleteAddress(ctx context.Context, id valueobject.Identifier) (*addressEntity.Address, error) {
	// get address to delete
	a, err := s.repo.Load(ctx, id)
	if err != nil {
		return &addressEntity.Address{}, err
	}

	// delete address
	a.DeleteAddress(id)

	// save event
	if _, err := s.repo.Save(ctx, a); err != nil {
		return &addressEntity.Address{}, err
	}

	return a, nil
}

// GetAddress gets address of provided ID
func (s *Service) GetAddress(ctx context.Context, id valueobject.Identifier) (*addressEntity.Address, error) {
	a, err := s.repo.Load(ctx, id)
	if err != nil {
		return &addressEntity.Address{}, err
	}
	return a, err
}

// GetAddresses get existing addresses, setting includeDeleted returns also addresses marked as deleted
func (s *Service) GetAddresses(ctx context.Context, includeDeleted bool) ([]addressEntity.Address, error) {
	var addresses []addressEntity.Address
	ids, err := s.repo.GetAddressIds(ctx, includeDeleted)
	if err != nil {
		return []addressEntity.Address{}, err
	}

	for _, id := range ids {
		a, err := s.GetAddress(ctx, id)
		if err != nil {
			return []addressEntity.Address{}, err
		}

		addresses = append(addresses, *a)
	}

	return addresses, nil

}

// checks itf there is a difference between provided addresses
func isIdentical(a addressEntity.Address, street valueobject.Street, postalCode valueobject.PostalCode, locality valueobject.Locality, country valueobject.Country, canton valueobject.Canton, additional valueobject.Additional) bool {
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
