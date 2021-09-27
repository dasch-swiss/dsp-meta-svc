package address_test

import (
	"context"
	addressEntity "github.com/dasch-swiss/dsp-meta-svc/services/metadata/go_ca_backend/entity/address"
	"testing"
	"time"

	"github.com/EventStore/EventStore-Client-Go/direction"
	"github.com/EventStore/EventStore-Client-Go/streamrevision"
	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/go_ca_backend/infrastructure/repository/address"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
)

func Test_AddressRepository_Save(t *testing.T) {
	container := GetEmptyDatabase()
	defer container.Close()

	c := CreateTestClient(container, t)
	defer c.Close()

	r := address.NewAddressRepository(c)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	// create value objects
	id, err := valueobject.NewIdentifier()
	street, _ := valueobject.NewStreet("Banhofstrasse 1")
	postalCode, _ := valueobject.NewPostalCode("4053")
	locality, _ := valueobject.NewLocality("Basel")
	country, _ := valueobject.NewCountry("Switzerland")
	canton, _ := valueobject.NewCanton("Basel-Stadt")
	additional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")

	// create new address
	expectedAddress := addressEntity.NewAddress(id, street, postalCode, locality, country, canton, additional)

	// save event to the store
	_, err = r.Save(ctx, expectedAddress)
	assert.Nil(t, err)

	// retrieve events from the store
	streamID := "Address-" + id.String()
	recordedEvents, err := c.ReadStreamEvents(ctx, direction.Forwards, streamID, streamrevision.StreamRevisionStart, 1, false)
	if err != nil {
		t.Fatalf("Unexpected failure: %+v", err)
	}

	// check the recorded event type
	assert.Equal(t, "AddressCreated", recordedEvents[0].EventType)
}

func Test_AddressRepository_Load(t *testing.T) {
	container := GetEmptyDatabase()
	defer container.Close()

	c := CreateTestClient(container, t)
	defer c.Close()

	r := address.NewAddressRepository(c)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	// create value objects
	id, _ := valueobject.NewIdentifier()
	street, _ := valueobject.NewStreet("Banhofstrasse 1")
	postalCode, _ := valueobject.NewPostalCode("4053")
	locality, _ := valueobject.NewLocality("Basel")
	country, _ := valueobject.NewCountry("Switzerland")
	canton, _ := valueobject.NewCanton("Basel-Stadt")
	additional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")

	// create new address
	a := addressEntity.NewAddress(id, street, postalCode, locality, country, canton, additional)

	// save event to the store
	r.Save(ctx, a)

	// load address from store
	addressFromEvents, err := r.Load(ctx, id)
	if err != nil {
		t.Fatalf("Unexpected failure: %+v", err)
	}

	// check if initially created address is the same as one loaded from events
	assert.Equal(t, a.ID, addressFromEvents.ID)
	assert.Equal(t, a.Type, addressFromEvents.Type)
	assert.Equal(t, a.Street, addressFromEvents.Street)
	assert.Equal(t, a.PostalCode, addressFromEvents.PostalCode)
	assert.Equal(t, a.Locality, addressFromEvents.Locality)
	assert.Equal(t, a.Country, addressFromEvents.Country)
	assert.Equal(t, a.Canton, addressFromEvents.Canton)
	assert.Equal(t, a.Additional, addressFromEvents.Additional)
}

func Test_AddressRepository_GetAddressIds(t *testing.T) {
	container := GetEmptyDatabase()
	defer container.Close()

	c := CreateTestClient(container, t)
	defer c.Close()

	r := address.NewAddressRepository(c)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	// create value objects
	id, _ := valueobject.NewIdentifier()
	street, _ := valueobject.NewStreet("Banhofstrasse 1")
	postalCode, _ := valueobject.NewPostalCode("4053")
	locality, _ := valueobject.NewLocality("Basel")
	country, _ := valueobject.NewCountry("Switzerland")
	canton, _ := valueobject.NewCanton("Basel-Stadt")
	additional, _ := valueobject.NewAdditional("Es ist keine echte Adresse")

	// create new address
	a := addressEntity.NewAddress(id, street, postalCode, locality, country, canton, additional)

	// save event to the store
	r.Save(ctx, a)

	// get list of address ids
	addressIds, err := r.GetAddressIds(ctx, false)
	if err != nil {
		t.Fatalf("Unexpected failure: %+v", err)
	}

	assert.Len(t, addressIds, 1)
}
