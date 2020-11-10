package transform

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

// PbToAccount .
func PbToAccount(acc *accountpb.Account) *accounts.Account {
	return &accounts.Account{}
}

// AccountsToPb .
func AccountsToPb(slice []accounts.Account) []*accountpb.Account {
	var marshaled []*accountpb.Account
	for _, account := range slice {
		marshaled = append(marshaled, AccountToPb(&account))
	}
	return marshaled
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
