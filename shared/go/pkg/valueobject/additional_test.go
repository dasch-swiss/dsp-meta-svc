package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewAdditional(t *testing.T) {
	v, _ := valueobject.NewAdditional("more address info")
	assert.Equal(t, v.String(), "more address info")
}

func Test_Invalid_NewAdditional(t *testing.T) {
	_, err := valueobject.NewAdditional("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewAdditional(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewAdditional("This is more than 50 characters of additional address info")
	assert.NotNil(t, err3)
}

func Test_Additional_Equals(t *testing.T) {
	a, _ := valueobject.NewAdditional("abc")
	b, _ := valueobject.NewAdditional("abc")
	assert.True(t, a.Equals(b))
}

func Test_Additional_NotEquals(t *testing.T) {
	a, _ := valueobject.NewAdditional("abc")
	b, _ := valueobject.NewAdditional("def")
	assert.False(t, a.Equals(b))
}
