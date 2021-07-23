package valueobject_test

import (
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIdentifier_String(t *testing.T) {
	identifier, err := valueobject.NewIdentifier()
	assert.Nil(t, err)
	assert.NotEmpty(t, identifier.String())
}

func TestIdentifierFromBytes(t *testing.T) {
	b := []byte("dc62dcd0-fb83-4488-8e5e-6d361ac79b6b")

	identifier, err := valueobject.IdentifierFromBytes(b)
	assert.Nil(t, err)
	assert.Equal(t, identifier.String(), "dc62dcd0-fb83-4488-8e5e-6d361ac79b6b")
}

func TestIdentifierFromBytes_InvalidUUID(t *testing.T) {
	b := []byte("d")

	_, err := valueobject.IdentifierFromBytes(b)
	assert.NotNil(t, err)
}
