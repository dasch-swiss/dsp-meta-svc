package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAdditional(t *testing.T) {
	v, _ := valueobject.NewAdditional("more address info")
	assert.Equal(t, v.String(), "more address info")
}

func TestInvalidNewAdditional(t *testing.T) {
	_, err := valueobject.NewAdditional("")
	assert.NotNil(t, err)

	_, err2 := valueobject.NewAdditional(" ")
	assert.NotNil(t, err2)

	_, err3 := valueobject.NewAdditional("This is more than 50 characters of additional address info")
	assert.NotNil(t, err3)
}
