package valueobject

import (
	"fmt"
	"strings"
)

type LongName struct {
	value string
}

// NewLongName creates a new valid long name object.
func NewLongName(value string) (LongName, error) {
	if len(value) > 50 || strings.TrimSpace(value) == "" {
		return LongName{}, fmt.Errorf("invalid long name, must be within 50 characters and non-empty")
	}

	return LongName{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v LongName) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v LongName) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *LongName) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewLongName(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v LongName) Equals(value Value) bool {
	otherValueObject, ok := value.(LongName)
	return ok && v.value == otherValueObject.value
}
