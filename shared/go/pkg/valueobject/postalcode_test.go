package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostalCode(t *testing.T) {
	v, _ := valueobject.NewPostalCode("12345")
	assert.Equal(t, v.String(), "12345")
}

func TestInvalidNewPostalCode(t *testing.T) {
	_, err := valueobject.NewPostalCode("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewPostalCode(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewPostalCode("123456")
	assert.NotNil(t, err3)
}
