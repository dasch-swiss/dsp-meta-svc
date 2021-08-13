package valueobject

import "fmt"

type Password struct {
	value string
}

// NewPassword creates a new valid password object.
func NewPassword(value string) (Password, error) {
	if len(value) >= 8 && len(value) <= 50 {
		return Password{}, fmt.Errorf("invalid password, must be within 8 and 50 characters")
	}

	return Password{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v Password) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v Password) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *Password) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewPassword(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v Password) Equals(value Value) bool {
	otherValueObject, ok := value.(Password)
	return ok && v.value == otherValueObject.value
}
