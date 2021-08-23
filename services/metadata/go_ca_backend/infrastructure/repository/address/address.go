package address

import (
	"context"
	"encoding/json"
	"fmt"
	addressEntity "github.com/dasch-swiss/dsp-meta-svc/services/metadata/go_ca_backend/entity/address"
	"log"
	"time"

	"github.com/EventStore/EventStore-Client-Go/client"
	"github.com/EventStore/EventStore-Client-Go/direction"
	"github.com/EventStore/EventStore-Client-Go/messages"
	"github.com/EventStore/EventStore-Client-Go/position"
	"github.com/EventStore/EventStore-Client-Go/streamrevision"
	"github.com/dasch-swiss/dsp-meta-svc/services/metadata/go_ca_backend/event"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/gofrs/uuid"
)

// contains a pointer to the client
type addressRepository struct {
	c *client.Client
}

// NewAddressRepository creates new repository to store address events in
func NewAddressRepository(client *client.Client) *addressRepository {
	return &addressRepository{
		c: client,
	}
}

// Save saves address events in events store/repository
func (r *addressRepository) Save(ctx context.Context, a *addressEntity.Address) (valueobject.Identifier, error) {
	var proposedEvents []messages.ProposedEvent
	streamRevision := streamrevision.StreamRevisionStreamExists

	for _, events := range a.Events() {
		switch e := events.(type) {
		case *event.AddressCreated:
			j, err := json.Marshal(e)
			if err != nil {
				return e.ID, fmt.Errorf("problem sertializing '%T' event to JSON", e)
			}

			eventID, _ := uuid.NewV4()
			pe := messages.ProposedEvent{
				EventID:     eventID,
				EventType:   "AddressCreated",
				ContentType: "application/json",
				Data:        j,
			}

			proposedEvents = append(proposedEvents, pe)
			streamRevision = streamrevision.StreamRevisionNoStream

		case *event.AddressChanged:
			j, err := json.Marshal(e)
			if err != nil {
				return e.ID, fmt.Errorf("problem sertializing '%T' event to JSON", e)
			}

			eventID, _ := uuid.NewV4()
			pe := messages.ProposedEvent{
				EventID:     eventID,
				EventType:   "AddressChanged",
				ContentType: "application/json",
				Data:        j,
			}

			proposedEvents = append(proposedEvents, pe)

		case *event.AddressDeleted:
			j, err := json.Marshal(e)
			if err != nil {
				return e.ID, fmt.Errorf("problem sertializing '%T' event to JSON", e)
			}

			eventID, _ := uuid.NewV4()
			pe := messages.ProposedEvent{
				EventID:     eventID,
				EventType:   "AddressDeleted",
				ContentType: "application/json",
				Data:        j,
			}

			proposedEvents = append(proposedEvents, pe)
		}
	}

	streamID := "Address-" + a.ID.String()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	_, err := r.c.AppendToStream(ctx, streamID, streamRevision, proposedEvents)
	if err != nil {
		log.Fatalf("Unexpected failure: %+v", err)
	}

	return a.ID, nil
}

// Load loads addresses from event store
func (r *addressRepository) Load(ctx context.Context, id valueobject.Identifier) (*addressEntity.Address, error) {
	streamID := "Address-" + id.String()

	// TODO find better way to return all events, not only hardcoded number of 1000 events
	recordedEvents, err := r.c.ReadStreamEvents(ctx, direction.Forwards, streamID, streamrevision.StreamRevisionStart, 1000, false)
	if err != nil {
		log.Printf("Unexpected failure: %+v", err)
		return &addressEntity.Address{}, addressEntity.ErrAddressNotFound
	}

	var events []event.Event

	for _, record := range recordedEvents {
		switch eventType := record.EventType; eventType {
		case "AddressCreated":
			var e event.AddressCreated
			err := json.Unmarshal(record.Data, &e)
			if err != nil {
				return &addressEntity.Address{}, fmt.Errorf("problem desertializing '%T' event to JSON", e)
			}
			events = append(events, &e)

		case "AddressChanged":
			var e event.AddressChanged
			err := json.Unmarshal(record.Data, &e)
			if err != nil {
				return &addressEntity.Address{}, fmt.Errorf("problem desertializing '%T' event to JSON", e)
			}
			events = append(events, &e)

		case "AddressDeleted":
			var e event.AddressDeleted
			err := json.Unmarshal(record.Data, &e)
			if err != nil {
				return &addressEntity.Address{}, fmt.Errorf("problem desertializing '%T' event to JSON", e)
			}
			events = append(events, &e)

		default:
			log.Printf("Unexpected event type: %T", eventType)
		}
	}

	return addressEntity.NewAddressFromEvents(events), nil
}

// GetAddressIds returns list of all active addresses ids
func (r *addressRepository) GetAddressIds(ctx context.Context, includeDeleted bool) ([]valueobject.Identifier, error) {
	// TODO: improve it especially because of looping through deleted addresses isn't efficient https://stackoverflow.com/a/37335777/9338572
	eventsToRead := 1000
	numberOfEvents := uint64(eventsToRead)

	recordedEvents, err := r.c.ReadAllEvents(ctx, direction.Forwards, position.StartPosition, numberOfEvents, true)
	if err != nil {
		log.Printf("Unexpected failure %+v", err)
		return nil, err
	}

	var addressIds []valueobject.Identifier

	// filter to select only AddressCreated events
	for _, record := range recordedEvents {
		switch eventType := record.EventType; eventType {

		case "AddressCreated":
			var e event.AddressCreated
			err := json.Unmarshal(record.Data, &e)
			if err != nil {
				return []valueobject.Identifier{}, fmt.Errorf("problem deserializing '%s' event from JSON", record.EventType)
			}
			addressIds = append(addressIds, e.ID)

		case "AddressDeleted":
			var e event.AddressDeleted
			err := json.Unmarshal(record.Data, &e)
			if err != nil {
				return []valueobject.Identifier{}, fmt.Errorf("problem deserializing '%s' event from JSON", record.EventType)
			}
			// if deleted address should not be returned - loop through the address ids
			if !includeDeleted {
				// check if address ID is marked as deleted and remove it from results
				isNotDeleted := func(x valueobject.Identifier) bool { return x != e.ID }
				i := 0
				for _, x := range addressIds {
					if isNotDeleted(x) {
						addressIds[i] = x
						i++
					}
				}
				addressIds = addressIds[:i]
			}
		}
	}

	return addressIds, nil
}
