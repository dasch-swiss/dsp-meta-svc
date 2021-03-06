package valueobject

import (
	"fmt"
	"strings"
)

type PostalCode struct {
	Value string
}

// NewPostalCode creates a new valid PostalCode object
func NewPostalCode(Value string) (PostalCode, error) {
	if len(Value) > 5 || strings.TrimSpace(Value) == "" {
		return PostalCode{}, fmt.Errorf("postalCode can't be enpty or longer than 5 characters")
	}

	return PostalCode{Value: Value}, nil
}

// implements fmt.Stringer interface
func (v PostalCode) String() string {
	return v.Value
}

// MarshalText serializes object
func (v PostalCode) MarshalText() ([]byte, error) {
	return []byte(v.Value), nil
}

// UnmarshalText deserializes object and returns an error if it's invalid
func (v *PostalCode) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewPostalCode(string(b))
	return err
}

// Equals checks if two value objects are the same
func (v PostalCode) Equals(value Value) bool {
	otherValueObject, ok := value.(PostalCode)
	return ok && v.Value == otherValueObject.Value
}
