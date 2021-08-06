package address

import "errors"

var ErrAddressHasBeenDeleted = errors.New("Address has been marked as deleted")
var ErrNoPropertiesChanged = errors.New("No new value for any property provided")
var ErrAddressNotFound = errors.New("No address found with the provided uuid")
var ErrServerNotResponding = errors.New("Server is not responding")
var ErrNoAddressDataReturned = errors.New("No address data returned")
