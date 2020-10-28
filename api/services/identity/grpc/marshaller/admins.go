package marshaller

import (
	"go.xixo.com/api/services/identity/domain/admins"
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/protobuf/identitypb"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AdminToPb .
func AdminToPb(a *admins.Admin) *identitypb.Admin {
	pb := &identitypb.Admin{
		Name:       a.Name(),
		FirstName:  a.FirstName,
		LastName:   a.LastName,
		Email:      a.Email,
		CreateTime: timestamppb.New(a.CreatedAt),
		UpdateTime: timestamppb.New(a.UpdatedAt),
	}

	var roleNames []string
	for _, id := range a.RoleIDs {
		roleNames = append(roleNames, roles.Name{RoleID: id}.String())
	}
	pb.RoleNames = roleNames

	return pb
}

// AdminsToPb .
func AdminsToPb(slice []*admins.Admin) []*identitypb.Admin {
	var marshaled []*identitypb.Admin
	for _, user := range slice {
		marshaled = append(marshaled, AdminToPb(user))
	}
	return marshaled
}

// PbToCreateAdminInput .
func PbToCreateAdminInput(pb *identitypb.Admin) *admins.CreateAdminInput {
	admin := &admins.CreateAdminInput{
		FirstName: pb.FirstName,
		LastName:  pb.LastName,
		Email:     pb.Email,
	}
	var roleIDs []string
	for _, n := range pb.RoleNames {
		// TODO: handle error
		name, _ := roles.ParseResourceName(n)
		roleIDs = append(roleIDs, name.RoleID)
	}
	admin.RoleIDs = roleIDs
	return admin
}

// PbToUpdateAdminInput .
func PbToUpdateAdminInput(pb *identitypb.Admin) *admins.UpdateAdminInput {
	admin := &admins.UpdateAdminInput{
		FirstName: pb.FirstName,
		LastName:  pb.LastName,
		Email:     pb.Email,
	}
	var roleIDs []string
	for _, n := range pb.RoleNames {
		// TODO: handle error
		name, _ := roles.ParseResourceName(n)
		roleIDs = append(roleIDs, name.RoleID)
	}
	admin.RoleIDs = roleIDs
	return admin
}

// PbToAdminUpdateMask .
func PbToAdminUpdateMask(pb *field_mask.FieldMask) (*admins.UpdateMask, error) {
	if !pb.IsValid(&identitypb.Admin{}) {
		return nil, ErrInvalidUpdateMask
	}
	mask := &admins.UpdateMask{}
	for _, path := range pb.Paths {
		switch path {
		case "first_name":
			mask.FirstName = true
			break
		case "last_name":
			mask.LastName = true
			break
		case "email":
			mask.Email = true
			break
		case "role_names":
			mask.RoleIDs = true
			break
		}
	}
	return mask, nil
}
