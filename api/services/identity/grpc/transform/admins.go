package transform

import (
	"go.xixo.com/api/services/identity/domain/admins"
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
		RoleNames:  a.RoleNames(),
		CreateTime: timestamppb.New(a.CreatedAt),
		UpdateTime: timestamppb.New(a.UpdatedAt),
	}
	return pb
}

// AdminsToPb .
func AdminsToPb(slice []admins.Admin) []*identitypb.Admin {
	var marshaled []*identitypb.Admin
	for _, user := range slice {
		marshaled = append(marshaled, AdminToPb(&user))
	}
	return marshaled
}

// PbToAdmin .
func PbToAdmin(pb *identitypb.Admin) *admins.Admin {
	return nil
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
			mask.Roles = true
			break
		}
	}
	return mask, nil
}
