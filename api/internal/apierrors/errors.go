package apierrors

import "errors"

// ErrNotFound indicates an expected resource was not found.
var ErrNotFound = errors.New("not found")
