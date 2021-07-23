package valueobject_test

import (
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShortCode(t *testing.T) {
	a, _ := valueobject.NewShortCode("00FF")
	assert.Equal(t, a.String(), "00FF")
}

func TestNewInvalidShortCode(t *testing.T) {
	_, err := valueobject.NewShortCode("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewShortCode(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewShortCode("test")
	assert.NotNil(t, err3)
}

func TestShortCode_Equals(t *testing.T) {
	a, _ := valueobject.NewShortCode("00FF")
	b, _ := valueobject.NewShortCode("00FF")
	assert.True(t, a.Equals(b))
}

func TestShortCode_NotEquals(t *testing.T) {
	a, _ := valueobject.NewShortCode("00FF")
	b, _ := valueobject.NewShortCode("11AA")
	assert.False(t, a.Equals(b))
}
