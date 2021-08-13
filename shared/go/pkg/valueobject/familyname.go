package valueobject

import "fmt"

type FamilyName struct {
	value string
}

// NewFamilyName creates a new valid given name object.
func NewFamilyName(value string) (FamilyName, error) {
	if len(value) > 50 || value == "" {
		return FamilyName{}, fmt.Errorf("invalid family name, must be within 50 characters and non-empty")
	}

	return FamilyName{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v FamilyName) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v FamilyName) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *FamilyName) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewFamilyName(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v FamilyName) Equals(value Value) bool {
	otherValueObject, ok := value.(FamilyName)
	return ok && v.value == otherValueObject.value
}
