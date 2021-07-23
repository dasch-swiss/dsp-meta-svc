package valueobject

import "fmt"

type Username struct {
	value string
}

// NewUsername creates a new valid username object.
func NewUsername(value string) (Username, error) {
	if len(value) > 50 || value == "" {
		return Username{}, fmt.Errorf("invalid username, must be within 50 characters and non-empty")
	}

	return Username{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v Username) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v Username) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *Username) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewUsername(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v Username) Equals(value Value) bool {
	otherValueObject, ok := value.(Username)
	return ok && v.value == otherValueObject.value
}
