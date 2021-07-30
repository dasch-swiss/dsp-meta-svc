package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCountry(t *testing.T) {
	v, _ := valueobject.NewCountry("Kongo")
	assert.Equal(t, v.String(), "Kongo")
}

func TestInvalidNewCountry(t *testing.T) {
	_, err := valueobject.NewCountry("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewCountry(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewCountry("LongerThan25CharactersCountry")
	assert.NotNil(t, err3)
}
