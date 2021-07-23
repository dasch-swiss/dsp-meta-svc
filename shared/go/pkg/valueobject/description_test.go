package valueobject_test

import (
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDescription(t *testing.T) {
	a, _ := valueobject.NewDescription("abc")
	assert.Equal(t, a.String(), "abc")
}

func TestNewInvalidDescription(t *testing.T) {
	_, err := valueobject.NewDescription("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewDescription(" ")
	assert.NotNil(t, err2)
}

func TestDescription_Equals(t *testing.T) {
	a, _ := valueobject.NewDescription("abc")
	b, _ := valueobject.NewDescription("abc")
	assert.True(t, a.Equals(b))
}

func TestDescription_NotEquals(t *testing.T) {
	a, _ := valueobject.NewDescription("abc")
	b, _ := valueobject.NewDescription("def")
	assert.False(t, a.Equals(b))
}
