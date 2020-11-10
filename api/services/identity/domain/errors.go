package domain

import (
	"errors"
)

var (
	// ErrNotFound is returned when requested resource cannot be found
	ErrNotFound = errors.New("not found")

	// ErrInvalidName is returned when provided resource name
	// does not match pattern
	ErrInvalidName = errors.New("invalid resource name")

	// ErrPageSizeOurOfBoundaries is returned when provided pageSize
	// is smaller than 1 or when it exceedes domain specific maxPageSize constant.
	ErrPageSizeOurOfBoundaries = errors.New("specified page size is out of boundaries")

	// ErrInvalidPageToken is returned when page token decode
	// failed to decode provided pageToken.
	//
	// It shouldn't be returned when pageToken is empty:
	// - empty pageToken denotes start of listing of the resource.
	ErrInvalidPageToken = errors.New("invalid page token provided, failed to decode")
)
