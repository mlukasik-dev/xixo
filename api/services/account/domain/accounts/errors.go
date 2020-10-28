package accounts

import "errors"

// ErrInvalidResourceName .
var ErrInvalidResourceName = errors.New("invalid resource name")

// ErrNotFound .
var ErrNotFound = errors.New("account not found")

// ErrPageSizeOurOfBoundaries .
var ErrPageSizeOurOfBoundaries = errors.New("specified page size if our of boundaries")

// ErrInvalidPageToken .
var ErrInvalidPageToken = errors.New("failed to decode provided page token")
