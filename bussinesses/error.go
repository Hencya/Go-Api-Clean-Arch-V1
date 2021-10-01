package bussinesses

import "errors"

var (
	ErrInternalServer = errors.New("Something gone wrong, contact administrator")
	ErrNotFound       = errors.New("Data not found")
	ErrIdNotFound     = errors.New("Id Not Found")
	ErrDuplicateData  = errors.New("Duplicate data")
)
