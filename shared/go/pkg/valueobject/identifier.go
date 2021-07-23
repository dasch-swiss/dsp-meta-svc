package valueobject

import "github.com/gofrs/uuid"

type Identifier struct {
	value uuid.UUID
}

//NewIdentifier creates a new identifier value object
func NewIdentifier() (Identifier, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return Identifier{}, err
	}
	return Identifier{value: id}, nil
}

//IdentifierFromBytes creates an identifier value object from the provided byte array
func IdentifierFromBytes(input []byte) (Identifier, error) {
	// create empty Identifier
	id := Identifier{}

	// assign the value of the Identifier
	err := id.UnmarshalText(input)
	if err != nil {
		return Identifier{}, err
	}

	return id, nil
}

//AsUUID returns the UUID of the identifier.
func (v Identifier) UUID() uuid.UUID {
	return v.value
}

// String implements the fmt.Stringer interface.
func (v Identifier) String() string {
	return v.value.String()
}

// MarshalText used to serialize the object
func (v Identifier) MarshalText() ([]byte, error) {
	return v.value.MarshalText()
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *Identifier) UnmarshalText(b []byte) error {
	var u uuid.UUID
	err := u.UnmarshalText(b)
	if err != nil {
		return err
	}
	*v = Identifier{value: u}
	return nil
}

//Equals tests for equality with another value object
func (v Identifier) Equals(value Value) bool {
	otherIdentifier, ok := value.(Identifier)
	return ok && v.value == otherIdentifier.value
}
