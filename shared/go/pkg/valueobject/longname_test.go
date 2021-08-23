package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLongName(t *testing.T) {
	a, _ := valueobject.NewLongName("abc")
	assert.Equal(t, a.String(), "abc")
}

func TestNewInvalidLongName(t *testing.T) {
	_, err := valueobject.NewLongName("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewLongName(" ")
	assert.NotNil(t, err2)
}

func TestLongName_Equals(t *testing.T) {
	a, _ := valueobject.NewLongName("abc")
	b, _ := valueobject.NewLongName("abc")
	assert.True(t, a.Equals(b))
}

func TestLongName_NotEquals(t *testing.T) {
	a, _ := valueobject.NewLongName("abc")
	b, _ := valueobject.NewLongName("def")
	assert.False(t, a.Equals(b))
}
