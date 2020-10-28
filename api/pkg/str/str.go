// Package str containes utility functions
// for dealing with string, string slices and `sql.NullString` type
package str

import "database/sql"

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

// NullStringsToStrings converts slice of `sql.NullString` to
// slice of strings, omits "invalid" strings
func NullStringsToStrings(nullSlice []sql.NullString) (slice []string) {
	for _, s := range nullSlice {
		if s.Valid {
			slice = append(slice, s.String)
		}
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
