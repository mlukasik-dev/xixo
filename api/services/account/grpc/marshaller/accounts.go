package marshaller

import (
	"go.xixo.com/api/services/account/domain/accounts"
	"go.xixo.com/protobuf/accountpb"

	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AccountToPb .
func AccountToPb(acc *accounts.Account) *accountpb.Account {
	return &accountpb.Account{
		Name:        acc.Name(),
		DisplayName: acc.DisplayName,
		CreateTime:  timestamppb.New(acc.CreatedAt),
		UpdateTime:  timestamppb.New(acc.CreatedAt),
	}
}

// AccountsToPb .
func AccountsToPb(slice []*accounts.Account) []*accountpb.Account {
	var marshaled []*accountpb.Account
	for _, account := range slice {
		marshaled = append(marshaled, AccountToPb(account))
	}
	return marshaled
}

// PbToCreateAccountInput transforms accountpb.Account to CreateAccountInput entity
// in case of invalid resource name return an error and nil
func PbToCreateAccountInput(pb *accountpb.Account) *accounts.CreateAccountInput {
	return &accounts.CreateAccountInput{
		DisplayName: pb.DisplayName,
	}
}

// PbToUpdateAccountInput transforms accountpb.Account to UpdateAccountInput entity
// in case of invalid resource name return an error and nil
func PbToUpdateAccountInput(pb *accountpb.Account) *accounts.UpdateAccountInput {
	return &accounts.UpdateAccountInput{
		DisplayName: pb.DisplayName,
	}
}

// PbToUpdateMask .
func PbToUpdateMask(pb *field_mask.FieldMask) (*accounts.UpdateMask, error) {
	if !pb.IsValid(&accountpb.Account{}) {
		return nil, ErrInvalidUpdateMask
	}
	mask := &accounts.UpdateMask{}
	for _, path := range pb.Paths {
		switch path {
		case "display_mame":
			mask.DisplayName = true
			break
		}
	}
	return mask, nil
}
