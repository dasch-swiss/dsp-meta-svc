package valueobject

import (
	"errors"
	url2 "net/url"
)

// EmailAddress errors
var (
	ErrInvalidURL = errors.New("not a valid URL")
)

// URL represents a valid URL address.
type URL struct {
	value string
}

// NewURL creates a new URL.
func NewURL(value string) (URL, error) {
	url := ZeroURL()

	u, err := url2.Parse(value)
	if err != nil {
		return url, ErrInvalidURL
	}

	if u.Host == "" {
		return url, ErrInvalidEmail
	}

	url.value = value

	return url, nil
}

// String returns string representation of the URL.
func (u URL) String() string {
	return u.value
}

// Equals checks that two URL addresses are the same.
func (u URL) Equals(value Value) bool {
	otherURL, ok := value.(Email)
	return ok && u.value == otherURL.value
}

//ZeroURL represents the zero value for an URL value object.
func ZeroURL() URL {
	return URL{}
}
