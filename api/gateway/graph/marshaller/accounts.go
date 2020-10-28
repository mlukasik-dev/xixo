package marshaller

import (
	"go.xixo.com/api/gateway/graph/model"
	"go.xixo.com/protobuf/accountpb"
)

// PbToAccount transforms protobuf user's type
// to user's model struct
func PbToAccount(pb *accountpb.Account) (*model.Account, error) {
	acc := &model.Account{
		DisplayName: pb.DisplayName,
		CreatedAt:   pb.CreateTime.AsTime(),
		UpdatedAt:   pb.UpdateTime.AsTime(),
	}
	return acc, nil
}
