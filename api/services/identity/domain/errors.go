package domain

import (
	"errors"
)

// ErrNotFound is returned when requested resource cannot be found
var ErrNotFound = errors.New("not found")

// ErrInvalidName is returned when provided resource name
// does not match pattern
var ErrInvalidName = errors.New("invalid resource name")

// ErrPageSizeOurOfBoundaries is returned when provided pageSize
// is smaller than 1 or when it exceedes domain specific maxPageSize constant.
var ErrPageSizeOurOfBoundaries = errors.New("specified page size is out of boundaries")

// ErrInvalidPageToken is returned when page token decode
// failed to decode provided pageToken.
//
// It shouldn't be returned when pageToken is empty:
// - empty pageToken denotes start of listing of the resource.
var ErrInvalidPageToken = errors.New("invalid page token provided, failed to decode")
