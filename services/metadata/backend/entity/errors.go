package address

import "errors"

// address has been marked as deleted
var ErrAddressHasBeenDeleted = errors.New("Address has been marked as deleted")

// invalid update values
var ErrNoPropertiesChanged = errors.New("No new value for any property provided")

// address not found
var ErrAddressNotFound = errors.New("No address found with the provided uuid")
