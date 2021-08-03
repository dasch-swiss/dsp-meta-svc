package valueobject

import (
	"fmt"
	"strings"
)

type Canton struct {
	Value string
}

// creates a new valid Canton object
func NewCanton(Value string) (Canton, error) {
	if len(Value) > 25 || strings.TrimSpace(Value) == "" {
		return Canton{}, fmt.Errorf("Error: canton can't be enpty or longer than 25 characters")
	}

	return Canton{Value: Value}, nil
}

// implements fmt.Stringer interface
func (v Canton) String() string {
	return v.Value
}

// serializes object
func (v Canton) MarshalText() ([]byte, error) {
	return []byte(v.Value), nil
}

// deserializes object and returns an error if it's invalid
func (v *Canton) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewCanton(string(b))
	return err
}

// checks if two value objects are the same
func (v Canton) Equals(value Value) bool {
	otherValueObject, ok := value.(Canton)
	return ok && v.Value == otherValueObject.Value
}
