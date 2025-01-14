package models

import (
	"errors"
	"fmt"
)

// ErrNotFound indicates the resource being queried for does not exist.
var ErrNotFound = errors.New("not found")

var ErrAssetNotFound = fmt.Errorf("asset not found: %w", ErrNotFound)
var ErrAssetSerialNotUnique = errors.New("serial not unique within model")

var ErrModelHasAssets = errors.New("model contains assets")
var ErrModelNotFound = fmt.Errorf("model not found: %w", ErrNotFound)
var ErrModelNotUnique = errors.New("model not unique within vendor")

var ErrVendorHasModels = errors.New("vendor with models cannot be deleted")
var ErrVendorNotFound = fmt.Errorf("vendor not found: %w", ErrNotFound)
