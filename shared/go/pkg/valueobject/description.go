package valueobject

import (
	"fmt"
	"strings"
)

type Description struct {
	value string
}

// NewDescription creates a new valid description object.
func NewDescription(value string) (Description, error) {
	if len(value) > 300 || strings.TrimSpace(value) == "" {
		return Description{}, fmt.Errorf("invalid description, must be within 300 characters and non-empty")
	}

	return Description{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v Description) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v Description) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *Description) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewDescription(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v Description) Equals(value Value) bool {
	otherValueObject, ok := value.(Description)
	return ok && v.value == otherValueObject.value
}
