package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStreet(t *testing.T) {
	v, _ := valueobject.NewStreet("new street")
	assert.Equal(t, v.String(), "new street")
}

func TestInvalidNewStreet(t *testing.T) {
	_, err := valueobject.NewStreet("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewStreet(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewStreet("this is longer than fifty characters street name!!!")
	assert.NotNil(t, err3)
}
