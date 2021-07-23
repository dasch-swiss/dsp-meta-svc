package valueobject_test

import (
	"fmt"
	"github.com/dasch-swiss/dsp-meta-svc/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Create a different type of value object used in Equals() check
type NotEmailAddress struct {
	value string
}

func (n NotEmailAddress) String() string {
	return n.value
}

func (n NotEmailAddress) Equals(value valueobject.Value) bool {
	return false
}

func ExampleEmail_String() {
	e, _ := valueobject.NewEmail("joe@blogs.com")

	fmt.Println(e.String())
	// Output: joe@blogs.com
}

func TestAcceptCorrectEmail(t *testing.T) {
	a, _ := valueobject.NewEmail("joe@blogs.com")
	assert.Equal(t, a.String(), "joe@blogs.com")

	b, _ := valueobject.NewEmail("i@example.org")
	assert.Equal(t, b.String(), "i@example.org")

	c, _ := valueobject.NewEmail("i@example.swiss")
	assert.Equal(t, c.String(), "i@example.swiss")

	d, _ := valueobject.NewEmail("test+i@example.swiss")
	assert.Equal(t, d.String(), "test+i@example.swiss")

	e, _ := valueobject.NewEmail("foo.bar@example.org")
	assert.Equal(t, e.String(), "foo.bar@example.org")
}

func TestShouldntAcceptInvalidEmail(t *testing.T) {
	_, aErr := valueobject.NewEmail("invalid")
	assert.NotNil(t, aErr)

	_, bErr := valueobject.NewEmail("email_at_example.org")
	assert.NotNil(t, bErr)
}

func ExampleEmail_Equals() {
	a, _ := valueobject.NewEmail("joe@blogs.com")
	b, _ := valueobject.NewEmail("joe@blogs.com")

	fmt.Println(a.Equals(b))
	// Output: true
}

func TestShouldCompareTwoEmailAsNotEqual(t *testing.T) {
	a, _ := valueobject.NewEmail("joe@blogs.com")
	b, _ := valueobject.NewEmail("mandy@blogs.com")
	assert.NotEqual(t, a, b)
}

func TestShouldNotBeEqualIfNotEmail(t *testing.T) {
	var notEmailAddress NotEmailAddress
	numeral, _ := valueobject.NewEmail("joe@blogs.com")

	if numeral.Equals(notEmailAddress) == true {
		t.Fatal("Different value object types can not be equal")
	}
}
