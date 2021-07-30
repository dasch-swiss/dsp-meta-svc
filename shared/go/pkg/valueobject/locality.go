package valueobject

import (
	"fmt"
	"strings"
)

type Locality struct {
	Value string
}

// creates a new valid Locality object
func NewLocality(Value string) (Locality, error) {
	if len(Value) > 25 || strings.TrimSpace(Value) == "" {
		return Locality{}, fmt.Errorf("Error: locality can't be enpty or longer than 25 characters")
	}

	return Locality{Value: Value}, nil
}

// implements fmt.Stringer interface
func (v Locality) String() string {
	return v.Value
}

// serializes object
func (v Locality) MarshalText() ([]byte, error) {
	return []byte(v.Value), nil
}

// deserializes object and returns an error if it's invalid
func (v *Locality) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewLocality(string(b))
	return err
}
