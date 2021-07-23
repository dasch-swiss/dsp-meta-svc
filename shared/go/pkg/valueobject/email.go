package valueobject

import (
	"errors"
	"regexp"
)

// EmailAddress errors
var (
	ErrInvalidEmail = errors.New("not a valid email address")
)

// Email represents a valid email address.
type Email struct {
	value string
}

// NewEmail creates a new email address.
func NewEmail(email string) (Email, error) {
	var e Email
	match, _ := regexp.MatchString(`([\w.]+)@([\w.]+)`, email)
	if !match {
		return e, ErrInvalidEmail
	}
	e.value = email

	return e, nil
}

// String returns string representation of the email address.
func (v Email) String() string {
	return v.value
}

// MarshalText used to serialize the object
func (v Email) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

// UnmarshalText used to deserialize the object and returns an error if it's invalid.
func (v *Email) UnmarshalText(b []byte) error {
	var err error
	*v, err = NewEmail(string(b))
	return err
}

//ZeroEmail represents the zero value for an email value object.
func ZeroEmail() Email {
	return Email{}
}

// Equals checks that two email addresses are the same.
func (v Email) Equals(value Value) bool {
	otherEmail, ok := value.(Email)
	return ok && v.value == otherEmail.value
}
