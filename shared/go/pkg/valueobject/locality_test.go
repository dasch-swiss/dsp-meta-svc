package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLocality(t *testing.T) {
	v, _ := valueobject.NewLocality("Szczebrzeszyn")
	assert.Equal(t, v.String(), "Szczebrzeszyn")
}

func TestInvalidNewLocality(t *testing.T) {
	_, err := valueobject.NewLocality("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewLocality(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewLocality("Longerthan25CharactersLocality")
	assert.NotNil(t, err3)
}

func Test_Locality_Equals(t *testing.T) {
	a, _ := valueobject.NewLocality("abc")
	b, _ := valueobject.NewLocality("abc")
	assert.True(t, a.Equals(b))
}

func Test_Locality_NotEquals(t *testing.T) {
	a, _ := valueobject.NewLocality("abc")
	b, _ := valueobject.NewLocality("def")
	assert.False(t, a.Equals(b))
}
