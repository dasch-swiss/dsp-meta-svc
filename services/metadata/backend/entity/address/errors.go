package address

import "errors"

var ErrAddressHasBeenDeleted = errors.New("address has been marked as deleted")
var ErrNoPropertiesChanged = errors.New("no new value for any property provided")
var ErrAddressNotFound = errors.New("no address found with the provided uuid")
var ErrServerNotResponding = errors.New("server is not responding")
var ErrNoAddressDataReturned = errors.New("no address data returned")
