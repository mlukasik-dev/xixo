package transform

import (
	"go.xixo.com/api/gateway/graph/model"
	"go.xixo.com/protobuf/identitypb"
)

// PermissionToPb transforms user's model struct
// to protobuf associated type
func PermissionToPb(p *model.Permission) *identitypb.Permission {
	return &identitypb.Permission{
		Method: p.Method,
	}
}

// PbToPermission transforms protobuf user's type
// to user's model struct
func PbToPermission(pb *identitypb.Permission) *model.Permission {
	return &model.Permission{
		Method: pb.Method,
	}
}

// PermissionsToPb .
func PermissionsToPb(pm []*model.Permission) []*identitypb.Permission {
	var pb []*identitypb.Permission
	for _, p := range pm {
		pb = append(pb, PermissionToPb(p))
	}
	return pb
}

// PbToPermissions .
func PbToPermissions(pb []*identitypb.Permission) (pm []*model.Permission) {
	for _, p := range pb {
		pm = append(pm, PbToPermission(p))
	}
	return pm
}
