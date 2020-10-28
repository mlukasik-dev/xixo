package postgres

import (
	"context"
	"testing"
)

func TestCheckPermission(t *testing.T) {
	db := MustConnect()
	r := repo{db}

	hasPermission, err := r.CheckPermission(context.Background(), "Super Admin", "/identitypb/Users/GetUser")
	if err != nil {
		t.FailNow()
	}

	t.Logf("result: %v\n", hasPermission)
}
