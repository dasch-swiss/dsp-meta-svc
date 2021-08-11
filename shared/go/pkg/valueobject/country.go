package valueobject

import (
	"fmt"
	"strings"
)

type Country struct {
	Value string
}

// NewCountry creates a new valid Country object
func NewCountry(Value string) (Country, error) {
	if len(Value) > 25 || strings.TrimSpace(Value) == "" {
		return Country{}, fmt.Errorf("country can't be enpty or longer than 25 characters")
	}

	return Country{Value: Value}, nil
}

// implements fmt.Stringer interface
func (v Country) String() string {
	return v.Value
}

// MarshalText serializes object
func (v Country) MarshalText() ([]byte, error) {
	return []byte(v.Value), nil
}

// UnmarshalText deserializes object and returns an error if it's invalid
func (v *Country) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewCountry(string(b))
	return err
}

// Equals checks if two value objects are the same
func (v Country) Equals(value Value) bool {
	otherValueObject, ok := value.(Country)
	return ok && v.Value == otherValueObject.Value
}
