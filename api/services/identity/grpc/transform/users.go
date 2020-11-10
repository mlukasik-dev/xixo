package transform

import (
	"go.xixo.com/api/services/identity/domain/users"
	"go.xixo.com/protobuf/identitypb"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserToPb .
func UserToPb(u *users.User) *identitypb.User {
	pb := &identitypb.User{
		Name:        u.Name(),
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		RoleNames:   u.RoleNames(),
		PhoneNumber: u.PhoneNumber.String,
		CreateTime:  timestamppb.New(u.CreatedAt),
		UpdateTime:  timestamppb.New(u.UpdatedAt),
	}
	return pb
}

// UsersToPb .
func UsersToPb(slice []users.User) []*identitypb.User {
	var marshaled []*identitypb.User
	for _, user := range slice {
		marshaled = append(marshaled, UserToPb(&user))
	}
	return marshaled
}

// PbToUser .
func PbToUser(pb *identitypb.User) *users.User {
	return nil
}

// PbToUserUpdateMask .
func PbToUserUpdateMask(pb *field_mask.FieldMask) (*users.UpdateMask, error) {
	if !pb.IsValid(&identitypb.User{}) {
		return nil, ErrInvalidUpdateMask
	}
	mask := &users.UpdateMask{}
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
		case "phone_number":
			mask.PhoneNumber = true
			break
		case "role_names":
			mask.Roles = true
			break
		}
	}
	return mask, nil
}
