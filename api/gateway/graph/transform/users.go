package transform

import (
	"github.com/google/uuid"
	"go.xixo.com/api/gateway/graph/model"
	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/pkg/str"
	"go.xixo.com/api/services/identity/domain/users"
	"go.xixo.com/protobuf/identitypb"

	"google.golang.org/genproto/protobuf/field_mask"
)

// PbToUser transforms protobuf user's type
// to user's model struct
func PbToUser(pb *identitypb.User) (*model.User, error) {
	name, err := users.ParseResourceName(pb.Name)
	if err != nil {
		return nil, err
	}
	roleIDs, err := roleNamesToRoleIDs(pb.RoleNames)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:          name.UserID.String(),
		FirstName:   pb.FirstName,
		LastName:    pb.LastName,
		Email:       pb.Email,
		PhoneNumber: &pb.PhoneNumber,
		Roles:       roleIDs,
		CreatedAt:   pb.CreateTime.AsTime(),
		UpdatedAt:   pb.UpdateTime.AsTime(),
	}, nil
}

// CreateUserInputToPB .
func CreateUserInputToPB(accountID string, input *model.CreateUserInput) *identitypb.CreateUserRequest {
	user := &identitypb.User{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		PhoneNumber: str.Dereference(input.PhoneNumber),
	}
	return &identitypb.CreateUserRequest{
		Parent: "accounts/" + accountID,
		User:   user,
	}
}

// UpdateUserInputToPB .
func UpdateUserInputToPB(accountID, userID string, i *model.UpdateUserInput) (*identitypb.UpdateUserRequest, error) {
	accountUUID, err := uuid.Parse(accountID)
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	user := &identitypb.User{
		Name:        users.Name{AccountID: accountUUID, UserID: userUUID}.String(),
		FirstName:   str.Dereference(i.FirstName),
		LastName:    str.Dereference(i.LastName),
		Email:       str.Dereference(i.Email),
		PhoneNumber: str.Dereference(i.PhoneNumber),
	}
	var mask field_mask.FieldMask
	var m identitypb.User
	if err := maskAppend(&mask, &m, i.FirstName, "first_name"); err != nil {
		return nil, err
	}
	if err := maskAppend(&mask, &m, i.LastName, "last_name"); err != nil {
		return nil, err
	}
	if err := maskAppend(&mask, &m, i.Email, "email"); err != nil {
		return nil, err
	}
	if err := maskAppend(&mask, &m, i.PhoneNumber, "phone_number"); err != nil {
		return nil, err
	}
	return &identitypb.UpdateUserRequest{
		User:       user,
		UpdateMask: &mask,
	}, nil
}

// PbToUserEdges .
func PbToUserEdges(users []*identitypb.User) (edges []*model.UserEdge, err error) {
	for _, user := range users {
		r, err := PbToUser(user)
		uuid, err := uuid.Parse(r.ID)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		edges = append(edges, &model.UserEdge{
			Node: r,
			Cursor: cursor.Encode(&cursor.Cursor{
				Timestamp: user.CreateTime.AsTime(),
				UUID:      uuid,
			}),
		})
	}
	return edges, nil
}
