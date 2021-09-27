package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewCountry(t *testing.T) {
	v, _ := valueobject.NewCountry("Kongo")
	assert.Equal(t, v.String(), "Kongo")
}

func Test_InvalidNewCountry(t *testing.T) {
	_, err := valueobject.NewCountry("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewCountry(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewCountry("LongerThan25CharactersCountry")
	assert.NotNil(t, err3)
}

func Test_Country_Equals(t *testing.T) {
	a, _ := valueobject.NewCountry("abc")
	b, _ := valueobject.NewCountry("abc")
	assert.True(t, a.Equals(b))
}

func Test_Country_NotEquals(t *testing.T) {
	a, _ := valueobject.NewCountry("abc")
	b, _ := valueobject.NewCountry("def")
	assert.False(t, a.Equals(b))
}
