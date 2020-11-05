// Package str containes utility functions
// for dealing with string, string slices and `sql.NullString` type
package str

import (
	"github.com/google/uuid"
)

// SliceContains returnes true if `s` - first argument
// containes string `e` - second arguments, otherwise returnes false
func SliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// UUIDSliceToStrings cenverts slice of uuid.UUID to
// slice of strings
func UUIDSliceToStrings(uuidSlice []uuid.UUID) (slice []string) {
	for _, id := range uuidSlice {
		slice = append(slice, id.String())
	}
	return
}

// Dereference if `s` is nil returnes empty string,
// otherwise dereferences it
func Dereference(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
