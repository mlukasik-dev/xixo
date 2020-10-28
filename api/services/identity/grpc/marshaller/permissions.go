package marshaller

import (
	"go.xixo.com/protobuf/identitypb"
)

// PermissionsToPb .
func PermissionsToPb(permissions []string) []*identitypb.Permission {
	var marshaled []*identitypb.Permission
	for _, method := range permissions {
		marshaled = append(marshaled, &identitypb.Permission{Method: method})
	}
	return marshaled
}

// PbToPermissions .
func PbToPermissions(slice []*identitypb.Permission) []string {
	var unmarshaled []string
	for _, p := range slice {
		unmarshaled = append(unmarshaled, p.Method)
	}
	return unmarshaled
}
