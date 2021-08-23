package valueobject

import "fmt"

type AggregateType struct {
	value string
}

// NewAggregateType creates a new valid aggregate type value object.
func NewAggregateType(value string) (AggregateType, error) {
	if len(value) > 50 || value == "" {
		return AggregateType{}, fmt.Errorf("invalid aggregate type, must be within 50 characters and non-empty")
	}
	return AggregateType{value: value}, nil
}

// String implements the fmt.Stringer interface.
func (v AggregateType) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v AggregateType) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *AggregateType) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewAggregateType(string(b))
	return err
}

// Equals checks that two value objects are the same.
func (v AggregateType) Equals(value Value) bool {
	otherValueObject, ok := value.(AggregateType)
	return ok && v.value == otherValueObject.value
}
