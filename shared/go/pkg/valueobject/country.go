package valueobject

import (
	"fmt"
	"strings"
)

type Country struct {
	Value string
}

// creates a new valid Country object
func NewCountry(Value string) (Country, error) {
	if len(Value) > 25 || strings.TrimSpace(Value) == "" {
		return Country{}, fmt.Errorf("Error: country can't be enpty or longer than 25 characters")
	}

	return Country{Value: Value}, nil
}

// implements fmt.Stringer interface
func (v Country) String() string {
	return v.Value
}

// serializes object
func (v Country) MarshalText() ([]byte, error) {
	return []byte(v.Value), nil
}

// deserializes object and returns an error if it's invalid
func (v *Country) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewCountry(string(b))
	return err
}
