package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PostalCode(t *testing.T) {
	v, _ := valueobject.NewPostalCode("12345")
	assert.Equal(t, v.String(), "12345")
}

func Test_InvalidNewPostalCode(t *testing.T) {
	_, err := valueobject.NewPostalCode("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewPostalCode(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewPostalCode("123456")
	assert.NotNil(t, err3)
}

func Test_PostalCode_Equals(t *testing.T) {
	a, _ := valueobject.NewPostalCode("abc")
	b, _ := valueobject.NewPostalCode("abc")
	assert.True(t, a.Equals(b))
}

func Test_PostalCode_NotEquals(t *testing.T) {
	a, _ := valueobject.NewPostalCode("abc")
	b, _ := valueobject.NewPostalCode("def")
	assert.False(t, a.Equals(b))
}
