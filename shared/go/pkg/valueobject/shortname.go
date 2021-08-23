package valueobject

import (
	"fmt"
	"strings"
)

type ShortName struct {
	value string
}

// NewShortName creates a new valid short name object.
func NewShortName(value string) (ShortName, error) {
	if len(value) > 15 || strings.TrimSpace(value) == "" {
		return ShortName{}, fmt.Errorf("invalid short name, must be within 15 characters and non-empty")
	}

	return ShortName{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v ShortName) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v ShortName) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *ShortName) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewShortName(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v ShortName) Equals(value Value) bool {
	otherValueObject, ok := value.(ShortName)
	return ok && v.value == otherValueObject.value
}
