package valueobject

import (
	"fmt"
	"strings"
)

type Street struct {
	Value string
}

// NewStreet creates a new valid Street object
func NewStreet(Value string) (Street, error) {
	if len(Value) > 50 || strings.TrimSpace(Value) == "" {
		return Street{}, fmt.Errorf("street can't be enpty or longer than 50 characters")
	}

	return Street{Value: Value}, nil
}

// implements fmt.Stringer interface
func (v Street) String() string {
	return v.Value
}

// MarshalText serializes object
func (v Street) MarshalText() ([]byte, error) {
	return []byte(v.Value), nil
}

// UnmarshalText deserializes object and returns an error if it's invalid
func (v *Street) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewStreet(string(b))
	return err
}

// Equals checks if two value objects are the same
func (v Street) Equals(value Value) bool {
	otherValueObject, ok := value.(Street)
	return ok && v.Value == otherValueObject.Value
}
