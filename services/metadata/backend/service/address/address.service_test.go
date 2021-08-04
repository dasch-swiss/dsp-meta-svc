package address_test

import (
	"context"
	"testing"
	"time"

	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/service/address"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateAddress(t *testing.T) {
	// create test environment
	repo := NewInMemRepo()
	service := address.NewService(repo)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	// create value objects
	expectedType := "http://ns.dasch.swiss/repository#Address"
	expectedStreet, _ := valueobject.NewStreet("Banhofstrasse 1")
	expectedPostalCode, _ := valueobject.NewPostalCode("4053")
	expectedLocality, _ := valueobject.NewLocality("Basel")
	expectedCountry, _ := valueobject.NewCountry("Switzerland")
	expectedCanton, _ := valueobject.NewCanton("Basel-Stadt")
	expectedAdditional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")

	// get id by creating new address
	addressId, err := service.CreateAddress(ctx, expectedStreet, expectedPostalCode, expectedLocality, expectedCountry, expectedCanton, expectedAdditional)

	// get address
	foundAddress, err := service.GetAddress(ctx, addressId)
	assert.Nil(t, err)
	assert.Equal(t, expectedType, foundAddress.Type)
	assert.Equal(t, expectedStreet, foundAddress.Street)
	assert.Equal(t, expectedPostalCode, foundAddress.PostalCode)
	assert.Equal(t, expectedLocality, foundAddress.Locality)
	assert.Equal(t, expectedCountry, foundAddress.Country)
	assert.Equal(t, expectedCanton, foundAddress.Canton)
	assert.Equal(t, expectedAdditional, foundAddress.Additional)
}

func TestService_GetAddresses(t *testing.T) {
	// create test environment
	repo := NewInMemRepo()
	service := address.NewService(repo)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	// create value objects
	expectedStreet, _ := valueobject.NewStreet("Banhofstrasse 1")
	expectedPostalCode, _ := valueobject.NewPostalCode("4053")
	expectedLocality, _ := valueobject.NewLocality("Basel")
	expectedCountry, _ := valueobject.NewCountry("Switzerland")
	expectedCanton, _ := valueobject.NewCanton("Basel-Stadt")
	expectedAdditional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")

	// get id by creating new address
	addressId, _ := service.CreateAddress(ctx, expectedStreet, expectedPostalCode, expectedLocality, expectedCountry, expectedCanton, expectedAdditional)

	// get a list of addresses
	addresses, err := service.GetAddresses(ctx, false)
	assert.Nil(t, err)
	assert.Len(t, addresses, 1)
	assert.Equal(t, addresses[0].ID, addressId)
}

func TestService_UpdateAddress(t *testing.T) {
	// create test environment
	repo := NewInMemRepo()
	service := address.NewService(repo)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	// create value objects
	expectedStreet, _ := valueobject.NewStreet("Banhofstrasse 1")
	expectedPostalCode, _ := valueobject.NewPostalCode("4053")
	expectedLocality, _ := valueobject.NewLocality("Basel")
	expectedCountry, _ := valueobject.NewCountry("Switzerland")
	expectedCanton, _ := valueobject.NewCanton("Basel-Stadt")
	expectedAdditional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")

	// get id by creating new address
	addressId, _ := service.CreateAddress(ctx, expectedStreet, expectedPostalCode, expectedLocality, expectedCountry, expectedCanton, expectedAdditional)

	// get address
	foundAddress, _ := service.GetAddress(ctx, addressId)

	newStreet, _ := valueobject.NewStreet("New Street")
	newPostalCode, _ := valueobject.NewPostalCode("0000")
	newLocality, _ := valueobject.NewLocality("New City")
	newCountry, _ := valueobject.NewCountry("New Country")
	newCanton, _ := valueobject.NewCanton("New Canton")
	newAdditional, _ := valueobject.NewAdditional("This is new additional info")

	// update address
	updatedAddress, err := service.UpdateAddress(ctx, addressId, newStreet, newPostalCode, newLocality, newCountry, newCanton, newAdditional)
	assert.Nil(t, err)
	assert.Equal(t, foundAddress.Type, updatedAddress.Type)
	assert.NotEqual(t, foundAddress.Street, updatedAddress.Street)
	assert.NotEqual(t, foundAddress.PostalCode, updatedAddress.PostalCode)
	assert.NotEqual(t, foundAddress.Locality, updatedAddress.Locality)
	assert.NotEqual(t, foundAddress.Country, updatedAddress.Country)
	assert.NotEqual(t, foundAddress.Canton, updatedAddress.Canton)
	// why?
	assert.NotEqual(t, foundAddress.Additional, updatedAddress.Additional)
}

func TestService_DeleteAddress(t *testing.T) {
	// create test environment
	repo := NewInMemRepo()
	service := address.NewService(repo)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	// create value objects
	expectedStreet, _ := valueobject.NewStreet("Banhofstrasse 1")
	expectedPostalCode, _ := valueobject.NewPostalCode("4053")
	expectedLocality, _ := valueobject.NewLocality("Basel")
	expectedCountry, _ := valueobject.NewCountry("Switzerland")
	expectedCanton, _ := valueobject.NewCanton("Basel-Stadt")
	expectedAdditional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")

	// get id by creating new address
	addressId, _ := service.CreateAddress(ctx, expectedStreet, expectedPostalCode, expectedLocality, expectedCountry, expectedCanton, expectedAdditional)

	deletedAddress, err := service.DeleteAddress(ctx, addressId)
	assert.Nil(t, err)
	assert.NotZero(t, deletedAddress.DeletedAt)
}
