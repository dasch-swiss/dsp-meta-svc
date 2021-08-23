package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewStreet(t *testing.T) {
	v, _ := valueobject.NewStreet("new street")
	assert.Equal(t, v.String(), "new street")
}

func Test_InvalidNewStreet(t *testing.T) {
	_, err := valueobject.NewStreet("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewStreet(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewStreet("this is longer than fifty characters street name!!!")
	assert.NotNil(t, err3)
}

func Test_Street_Equals(t *testing.T) {
	a, _ := valueobject.NewStreet("abc")
	b, _ := valueobject.NewStreet("abc")
	assert.True(t, a.Equals(b))
}

func Test_Street_NotEquals(t *testing.T) {
	a, _ := valueobject.NewStreet("abc")
	b, _ := valueobject.NewStreet("def")
	assert.False(t, a.Equals(b))
}
