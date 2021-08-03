package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCanton(t *testing.T) {
	v, _ := valueobject.NewCanton("Basel-Landschaft")
	assert.Equal(t, v.String(), "Basel-Landschaft")
}

func TestInvalidNewCanton(t *testing.T) {
	_, err := valueobject.NewCanton("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewCanton(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewCanton("LongerThan25CharactersCanton")
	assert.NotNil(t, err3)
}

func Test_Canton_Equals(t *testing.T) {
	a, _ := valueobject.NewCanton("abc")
	b, _ := valueobject.NewCanton("abc")
	assert.True(t, a.Equals(b))
}

func Test_Canton_NotEquals(t *testing.T) {
	a, _ := valueobject.NewCanton("abc")
	b, _ := valueobject.NewCanton("def")
	assert.False(t, a.Equals(b))
}
