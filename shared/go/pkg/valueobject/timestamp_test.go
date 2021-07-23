package valueobject_test

import (
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTimestamp(t *testing.T) {
	timestamp := valueobject.NewTimestamp()
	assert.NotEmpty(t, timestamp.String())
}

func TestNewTimestampFromTime(t *testing.T) {
	// TODO: add test
}

func TestTimestamp_Unix(t *testing.T) {
	expected := int64(1618337508)
	ts := valueobject.NewTimestampFromUnix(expected)
	assert.Equal(t, expected, ts.Unix())
}

func TestTimestamp_MarshalJSON(t *testing.T) {
	// var unixTimestamp int64 = 1618337508
	// expected := []byte(`{"bar": 1618337508}`)
	var unixTimestamp int64 = 1618337508
	expectedTimestamp := valueobject.NewTimestampFromUnix(unixTimestamp)
	marshaled, err := expectedTimestamp.MarshalJSON()
	assert.Nil(t, err)

	var unmarshalledTimestamp valueobject.Timestamp
	err = unmarshalledTimestamp.UnmarshalJSON(marshaled)
	assert.Nil(t, err)

	assert.Equal(t, expectedTimestamp.String(), unmarshalledTimestamp.String())
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	// TODO: add test
}
