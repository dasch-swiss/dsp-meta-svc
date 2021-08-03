package valueobject

import (
	"fmt"
	"strings"
)

type Additional struct {
	Value string
}

// creates a new valid Additional object
func NewAdditional(Value string) (Additional, error) {
	if len(Value) > 50 || strings.TrimSpace(Value) == "" {
		return Additional{}, fmt.Errorf("Error: additional can't be enpty or longer than 50 characters")
	}

	return Additional{Value: Value}, nil
}

// implements fmt.Stringer interface
func (v Additional) String() string {
	return v.Value
}

// serializes object
func (v Additional) MarshalText() ([]byte, error) {
	return []byte(v.Value), nil
}

// deserializes object and returns an error if it's invalid
func (v *Additional) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewAdditional(string(b))
	return err
}

// checks if two value objects are the same
func (v Additional) Equals(value Value) bool {
	otherValueObject, ok := value.(Additional)
	return ok && v.Value == otherValueObject.Value
}
