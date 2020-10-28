package marshaller

import (
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/protobuf/identitypb"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// RoleToPb takes Role domain entity
// and transforms it to the gRPC Role message
func RoleToPb(role *roles.Role) *identitypb.Role {
	marshaled := &identitypb.Role{
		Name:        role.Name(),
		AdminOnly:   role.AdminOnly,
		DisplayName: role.DisplayName,
		Description: role.Description,
		CreateTime:  timestamppb.New(role.CreatedAt),
		UpdateTime:  timestamppb.New(role.CreatedAt),
	}
	var permissions []*identitypb.Permission
	for _, method := range role.Permissions {
		permissions = append(permissions, &identitypb.Permission{
			Method: method.String,
		})
	}
	marshaled.Permissions = permissions
	return marshaled
}

// RolesToPb takes slice of Role domain entities
// and transforms them to the slice of gRPC Role messages
func RolesToPb(slice []*roles.Role) []*identitypb.Role {
	var marshaled []*identitypb.Role
	for _, role := range slice {
		marshaled = append(marshaled, RoleToPb(role))
	}
	return marshaled
}

// PbToCreateRoleInput .
func PbToCreateRoleInput(pb *identitypb.Role) *roles.CreateRoleInput {
	r := &roles.CreateRoleInput{
		AdminOnly:   pb.AdminOnly,
		DisplayName: pb.DisplayName,
		Description: pb.Description,
	}
	var permissions []string
	for _, p := range pb.Permissions {
		permissions = append(permissions, p.Method)
	}
	r.Permissions = permissions
	return r
}

// PbToUpdateRoleInput .
func PbToUpdateRoleInput(pb *identitypb.Role) *roles.UpdateRoleInput {
	r := &roles.UpdateRoleInput{
		AdminOnly:   pb.AdminOnly,
		DisplayName: pb.DisplayName,
		Description: pb.Description,
	}
	var permissions []string
	for _, p := range pb.Permissions {
		permissions = append(permissions, p.Method)
	}
	r.Permissions = permissions
	return r
}

// PbToRoleUpdateMask .
func PbToRoleUpdateMask(pb *field_mask.FieldMask) (*roles.UpdateMask, error) {
	if !pb.IsValid(&identitypb.Role{}) {
		return nil, ErrInvalidUpdateMask
	}
	mask := &roles.UpdateMask{}
	for _, path := range pb.Paths {
		switch path {
		case "admin_only":
			mask.AdminOnly = true
			break
		case "display_name":
			mask.DisplayName = true
			break
		case "description":
			mask.Description = true
			break
		case "permissions":
			mask.Permissions = true
			break
		}
	}
	return mask, nil
}
