package address_test

import (
	addressEntity "github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/entity/address"
	"testing"

	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/backend/event"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
)

const ADDRESS_TYPE = "http://ns.dasch.swiss/repository#Address"

func Test_NewAddress(t *testing.T) {
	expectedId, _ := valueobject.NewIdentifier()
	expectedType := ADDRESS_TYPE
	expectedStreet, _ := valueobject.NewStreet("Banhofstrasse 1")
	expectedPostalCode, _ := valueobject.NewPostalCode("4053")
	expectedLocality, _ := valueobject.NewLocality("Basel")
	expectedCountry, _ := valueobject.NewCountry("Switzerland")
	expectedCanton, _ := valueobject.NewCanton("Basel-Stadt")
	expectedAdditional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")

	a := addressEntity.NewAddress(expectedId, expectedStreet, expectedPostalCode, expectedLocality, expectedCountry, expectedCanton, expectedAdditional)
	assert.Equal(t, expectedId, a.ID)
	assert.Equal(t, expectedType, a.Type)
	assert.Equal(t, expectedStreet, a.Street)
	assert.Equal(t, expectedPostalCode, a.PostalCode)
	assert.Equal(t, expectedLocality, a.Locality)
	assert.Equal(t, expectedCountry, a.Country)
	assert.Equal(t, expectedAdditional, a.Additional)

	assert.False(t, a.CreatedAt.Time().IsZero())
	assert.True(t, a.ChangedAt.Time().IsZero())
	assert.True(t, a.DeletedAt.Time().IsZero())

	addressEvents := a.Changes
	createdEvent := addressEvents[0]

	switch e := createdEvent.(type) {
	case *event.AddressCreated:
		assert.Equal(t, expectedId, a.ID)
		assert.Equal(t, expectedType, a.Type)
		assert.Equal(t, expectedStreet, a.Street)
		assert.Equal(t, expectedPostalCode, a.PostalCode)
		assert.Equal(t, expectedLocality, a.Locality)
		assert.Equal(t, expectedCountry, a.Country)
		assert.Equal(t, expectedAdditional, a.Additional)
	default:
		t.Fatalf("unexpected event type: %T", e)
	}
}

func Test_NewAddressFromEvents(t *testing.T) {
	expectedId, _ := valueobject.NewIdentifier()
	expectedType := ADDRESS_TYPE
	expectedStreet, _ := valueobject.NewStreet("Banhofstrasse 1")
	expectedPostalCode, _ := valueobject.NewPostalCode("4053")
	expectedLocality, _ := valueobject.NewLocality("Basel")
	expectedCountry, _ := valueobject.NewCountry("Switzerland")
	expectedCanton, _ := valueobject.NewCanton("Basel-Stadt")
	expectedAdditional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")
	expectedCreatedAt := valueobject.NewTimestamp()
	expectedCreatedBy, _ := valueobject.NewIdentifier()

	createEvent := &event.AddressCreated{
		ID:         expectedId,
		Type:       expectedType,
		Street:     expectedStreet,
		PostalCode: expectedPostalCode,
		Locality:   expectedLocality,
		Country:    expectedCountry,
		Canton:     expectedCanton,
		Additional: expectedAdditional,
		CreatedAt:  expectedCreatedAt,
		CreatedBy:  expectedCreatedBy,
	}

	events := []event.Event{createEvent}
	print(events)

	a := addressEntity.NewAddressFromEvents(events)
	assert.Equal(t, expectedId, a.ID)
	assert.Equal(t, expectedType, a.Type)
	assert.Equal(t, expectedStreet, a.Street)
	assert.Equal(t, expectedPostalCode, a.PostalCode)
	assert.Equal(t, expectedLocality, a.Locality)
	assert.Equal(t, expectedCountry, a.Country)
	assert.Equal(t, expectedAdditional, a.Additional)

	assert.False(t, a.CreatedAt.Time().IsZero())
	assert.True(t, a.ChangedAt.Time().IsZero())
}

func Test_UpdateAddress(t *testing.T) {
	expectedId, _ := valueobject.NewIdentifier()
	expectedType := ADDRESS_TYPE
	expectedStreet, _ := valueobject.NewStreet("Banhofstrasse 1")
	expectedPostalCode, _ := valueobject.NewPostalCode("4053")
	expectedLocality, _ := valueobject.NewLocality("Basel")
	expectedCountry, _ := valueobject.NewCountry("Switzerland")
	expectedCanton, _ := valueobject.NewCanton("Basel-Stadt")
	expectedAdditional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")

	a := addressEntity.NewAddress(expectedId, expectedStreet, expectedPostalCode, expectedLocality, expectedCountry, expectedCanton, expectedAdditional)
	assert.Equal(t, expectedId, a.ID)
	assert.Equal(t, expectedType, a.Type)
	assert.Equal(t, expectedStreet, a.Street)
	assert.Equal(t, expectedPostalCode, a.PostalCode)
	assert.Equal(t, expectedLocality, a.Locality)
	assert.Equal(t, expectedCountry, a.Country)
	assert.Equal(t, expectedAdditional, a.Additional)
	assert.False(t, a.CreatedAt.Time().IsZero())
	assert.True(t, a.ChangedAt.Time().IsZero())

	addressEvents := a.Changes
	createdEvent := addressEvents[0]

	switch e := createdEvent.(type) {
	case *event.AddressCreated:
		assert.Equal(t, expectedId, a.ID)
		assert.Equal(t, expectedType, a.Type)
		assert.Equal(t, expectedStreet, a.Street)
		assert.Equal(t, expectedPostalCode, a.PostalCode)
		assert.Equal(t, expectedLocality, a.Locality)
		assert.Equal(t, expectedCountry, a.Country)
		assert.Equal(t, expectedAdditional, a.Additional)
	default:
		t.Fatalf("unexpected event type: %T", e)
	}

	newStreet, _ := valueobject.NewStreet("New Street")
	newPostalCode, _ := valueobject.NewPostalCode("0000")
	newLocality, _ := valueobject.NewLocality("New City")
	newCountry, _ := valueobject.NewCountry("New Country")
	newCanton, _ := valueobject.NewCanton("New Canton")
	newAdditional, _ := valueobject.NewAdditional("This is new additional info")

	a.UpdateAddress(expectedId, newStreet, newPostalCode, newLocality, newCountry, newCanton, newAdditional)

	assert.Len(t, a.Changes, 2)

	addressUpdatedEvent := a.Changes[1]

	switch e := addressUpdatedEvent.(type) {
	case *event.AddressChanged:
		assert.Equal(t, expectedId, e.ID)
		assert.Equal(t, expectedType, e.Type)
		assert.Equal(t, newStreet, e.Street)
		assert.Equal(t, newPostalCode, e.PostalCode)
		assert.Equal(t, newLocality, e.Locality)
		assert.Equal(t, newCountry, e.Country)
		assert.Equal(t, newAdditional, e.Additional)

		assert.False(t, a.ChangedAt.Time().IsZero())
		assert.IsType(t, a.ChangedBy, valueobject.Identifier{})
		assert.True(t, a.DeletedAt.Time().IsZero())

	default:
		t.Fatalf("unexpected event type: %T", e)
	}
}

func Test_DeleteAddress(t *testing.T) {
	expectedId, _ := valueobject.NewIdentifier()
	expectedType := ADDRESS_TYPE
	expectedStreet, _ := valueobject.NewStreet("Banhofstrasse 1")
	expectedPostalCode, _ := valueobject.NewPostalCode("4053")
	expectedLocality, _ := valueobject.NewLocality("Basel")
	expectedCountry, _ := valueobject.NewCountry("Switzerland")
	expectedCanton, _ := valueobject.NewCanton("Basel-Stadt")
	expectedAdditional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")

	a := addressEntity.NewAddress(expectedId, expectedStreet, expectedPostalCode, expectedLocality, expectedCountry, expectedCanton, expectedAdditional)
	assert.Equal(t, expectedId, a.ID)
	assert.Equal(t, expectedType, a.Type)
	assert.Equal(t, expectedStreet, a.Street)
	assert.Equal(t, expectedPostalCode, a.PostalCode)
	assert.Equal(t, expectedLocality, a.Locality)
	assert.Equal(t, expectedCountry, a.Country)
	assert.Equal(t, expectedAdditional, a.Additional)
	assert.False(t, a.CreatedAt.Time().IsZero())
	assert.True(t, a.ChangedAt.Time().IsZero())

	addressEvents := a.Changes
	createdEvent := addressEvents[0]

	switch e := createdEvent.(type) {
	case *event.AddressCreated:
		assert.Equal(t, expectedId, a.ID)
		assert.Equal(t, expectedType, a.Type)
		assert.Equal(t, expectedStreet, a.Street)
		assert.Equal(t, expectedPostalCode, a.PostalCode)
		assert.Equal(t, expectedLocality, a.Locality)
		assert.Equal(t, expectedCountry, a.Country)
		assert.Equal(t, expectedAdditional, a.Additional)
	default:
		t.Fatalf("unexpected event type: %T", e)
	}

	// deletedAddress := a.DeleteAddress(a.ID)
	// assert.Nil(t, deletedAddress)

	a.DeleteAddress(a.ID)

	assert.Len(t, a.Changes, 2)

	addressDeletedEvent := a.Changes[1]

	switch e := addressDeletedEvent.(type) {
	case *event.AddressDeleted:
		assert.Equal(t, a.ID, e.ID)
		assert.False(t, a.DeletedAt.Time().IsZero())
		assert.IsType(t, a.DeletedBy, valueobject.Identifier{})
	default:
		t.Fatalf("unexpected event type: %T", e)
	}

	updatedAddtess := a.UpdateAddress(expectedId, expectedStreet, expectedPostalCode, expectedLocality, expectedCountry, expectedCanton, expectedAdditional)
	assert.NotNil(t, updatedAddtess)

	// while trying to update deleted address an error should be returned
	err := a.UpdateAddress(expectedId, expectedStreet, expectedPostalCode, expectedLocality, expectedCountry, expectedCanton, expectedAdditional)
	// amd no more events added to the event array
	assert.Len(t, a.Changes, 2)
	assert.Equal(t, err, addressEntity.ErrAddressHasBeenDeleted)
}
