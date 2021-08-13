package valueobject

import (
	"fmt"
	"strconv"
)

type ShortCode struct {
	value string
}

// NewShortCode creates a new valid short code object.
func NewShortCode(value string) (ShortCode, error) {
	_, err := strconv.ParseUint(value, 16, 64)
	if err != nil {
		return ShortCode{}, fmt.Errorf("invalid short code, must be two hexadecimal digits and non-empty")
	}

	return ShortCode{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v ShortCode) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v ShortCode) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *ShortCode) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewShortCode(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v ShortCode) Equals(value Value) bool {
	otherValueObject, ok := value.(ShortCode)
	return ok && v.value == otherValueObject.value
}
