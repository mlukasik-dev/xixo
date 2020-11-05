package transform

import (
	"testing"

	"go.xixo.com/api/gateway/graph/model"
)

func TestUpdateRoleInputToPB(t *testing.T) {
	dn := ""
	_, err := UpdateRoleInputToPB("", &model.UpdateRoleInput{DisplayName: &dn})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
