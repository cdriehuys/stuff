package models

import (
	"errors"
)

// ErrNotFound indicates the resource being queried for does not exist.
var ErrNotFound = errors.New("not found")

var ErrVendorNotFound = errors.New("vendor not found")
var ErrModelNotUnique = errors.New("model not unique within vendor")

var ErrVendorHasModels = errors.New("vendor with models cannot be deleted")
