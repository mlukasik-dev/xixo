package authr

import (
	"sync"

	"github.com/google/uuid"
)

// Permission .
type Permission struct {
	RoleID uuid.UUID
	Method string
}

// Store .
type Store struct {
	permissions  map[Permission]bool
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

// NewStore .
func NewStore() *Store {
	return &Store{
		permissions: make(map[Permission]bool),
	}
}

// CheckPermission implements Checker
func (s *Store) CheckPermission(p Permission) (bool, error) {
	s.RLock()
	val := s.permissions[p]
	s.RUnlock()
	return val, nil
}

// GrantPermission .
func (s *Store) GrantPermission(p Permission) {
	s.Lock()
	s.permissions[p] = true
	s.Unlock()
}

// GrantPermissions .
func (s *Store) GrantPermissions(ps []Permission) {
	s.Lock()
	for _, p := range ps {
		s.permissions[p] = true
	}
	s.Unlock()
}

// DenyPermission .
func (s *Store) DenyPermission(p Permission) {
	s.Lock()
	delete(s.permissions, p)
	s.Unlock()
}
