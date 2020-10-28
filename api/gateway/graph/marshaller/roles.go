package marshaller

import (
	"go.xixo.com/api/gateway/graph/model"
	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/pkg/str"
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/protobuf/identitypb"

	"google.golang.org/genproto/protobuf/field_mask"
)

// PbToRole transforms protobuf user's type
// to user's model struct
func PbToRole(pb *identitypb.Role) (*model.Role, error) {
	name, err := roles.ParseResourceName(pb.Name)
	if err != nil {
		return nil, err
	}
	return &model.Role{
		ID:          name.RoleID,
		DisplayName: pb.DisplayName,
		Description: &pb.Description,
		Permissions: PbToPermissions(pb.Permissions),
		CreatedAt:   pb.CreateTime.AsTime(),
		UpdatedAt:   pb.UpdateTime.AsTime(),
	}, nil
}

// CreateRoleInputToPB .
func CreateRoleInputToPB(i *model.CreateRoleInput) *identitypb.CreateRoleRequest {
	var permissions []*identitypb.Permission
	for _, p := range i.Permissions {
		permissions = append(permissions, &identitypb.Permission{Method: p.Method})
	}
	role := &identitypb.Role{
		AdminOnly:   false, // users can't create admin_only roles
		DisplayName: i.DisplayName,
		Permissions: permissions,
		Description: str.Dereference(i.Description),
	}
	return &identitypb.CreateRoleRequest{Role: role}
}

// UpdateRoleInputToPB .
func UpdateRoleInputToPB(id string, i *model.UpdateRoleInput) (*identitypb.UpdateRoleRequest, error) {
	var permissions []*identitypb.Permission
	for _, p := range i.Permissions {
		permissions = append(permissions, &identitypb.Permission{Method: p.Method})
	}
	role := &identitypb.Role{
		DisplayName: str.Dereference(i.DisplayName),
		Permissions: permissions,
		Description: str.Dereference(i.Description),
	}
	var mask field_mask.FieldMask
	var m identitypb.Role
	if err := maskAppend(&mask, &m, i.DisplayName, "display_name"); err != nil {
		return nil, err
	}
	if err := maskAppend(&mask, &m, i.Description, "description"); err != nil {
		return nil, err
	}
	return &identitypb.UpdateRoleRequest{
		UpdateMask: &mask,
		Role:       role,
	}, nil
}

func roleNamesToRoleIDs(roleNames []string) ([]string, error) {
	var roleIDs []string
	for _, roleName := range roleNames {
		name, err := roles.ParseResourceName(roleName)
		if err != nil {
			return nil, err
		}
		roleIDs = append(roleIDs, name.RoleID)
	}
	return roleIDs, nil
}

func roleIDsToRoleNames(roleIDs []string) []string {
	var roleNames []string
	for _, roleID := range roleIDs {
		roleNames = append(roleNames, roles.Name{RoleID: roleID}.String())
	}
	return roleNames
}

// PbToRoleEdges .
func PbToRoleEdges(roles []*identitypb.Role) (edges []*model.RoleEdge, err error) {
	for _, role := range roles {
		r, err := PbToRole(role)
		if err != nil {
			return nil, err
		}
		edges = append(edges, &model.RoleEdge{
			Node: r,
			Cursor: cursor.Encode(&cursor.Cursor{
				Timestamp: role.CreateTime.AsTime(),
				UUID:      r.ID,
			}),
		})
	}
	return edges, nil
}
