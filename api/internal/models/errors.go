package models

import (
	"errors"
)

// ErrNotFound indicates the resource being queried for does not exist.
var ErrNotFound = errors.New("not found")

// ErrVendorNotFound indicates that the referenced vendor does not exist.
var ErrVendorNotFound = errors.New("vendor not found")
