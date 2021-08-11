package valueobject

import (
	"fmt"
	"strings"
)

type Locality struct {
	Value string
}

// NewLocality creates a new valid Locality object
func NewLocality(Value string) (Locality, error) {
	if len(Value) > 25 || strings.TrimSpace(Value) == "" {
		return Locality{}, fmt.Errorf("locality can't be enpty or longer than 25 characters")
	}

	return Locality{Value: Value}, nil
}

// implements fmt.Stringer interface
func (v Locality) String() string {
	return v.Value
}

// MarshalText serializes object
func (v Locality) MarshalText() ([]byte, error) {
	return []byte(v.Value), nil
}

// UnmarshalText deserializes object and returns an error if it's invalid
func (v *Locality) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewLocality(string(b))
	return err
}

// Equals checks if two value objects are the same
func (v Locality) Equals(value Value) bool {
	otherValueObject, ok := value.(Locality)
	return ok && v.Value == otherValueObject.Value
}
