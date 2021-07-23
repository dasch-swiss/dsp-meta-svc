package valueobject_test

import (
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShortName(t *testing.T) {
	a, _ := valueobject.NewShortName("abc")
	assert.Equal(t, a.String(), "abc")
}

func TestNewInvalidShortName(t *testing.T) {
	_, err := valueobject.NewShortName("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewShortName(" ")
	assert.NotNil(t, err2)
}

func TestShortName_Equals(t *testing.T) {
	a, _ := valueobject.NewShortName("abc")
	b, _ := valueobject.NewShortName("abc")
	assert.True(t, a.Equals(b))
}

func TestShortName_NotEquals(t *testing.T) {
	a, _ := valueobject.NewShortName("abc")
	b, _ := valueobject.NewShortName("def")
	assert.False(t, a.Equals(b))
}
