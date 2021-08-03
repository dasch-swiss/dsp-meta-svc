package address

import "errors"

// address has been marked as deleted
var ErrAddressHasBeenDeleted = errors.New("address has been marked as deleted")

// invalid update values
var ErrNoPropertiesChanged = errors.New("no new value for any property provided")
