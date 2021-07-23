package valueobject

import "fmt"

type GivenName struct {
	value string
}

// NewGivenName creates a new valid given name object.
func NewGivenName(value string) (GivenName, error) {
	if len(value) > 50 || value == "" {
		return GivenName{}, fmt.Errorf("invalid given name, must be within 50 characters and non-empty")
	}

	return GivenName{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v GivenName) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v GivenName) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *GivenName) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewGivenName(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v GivenName) Equals(value Value) bool {
	otherValueObject, ok := value.(GivenName)
	return ok && v.value == otherValueObject.value
}
