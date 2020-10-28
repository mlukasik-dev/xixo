package marshaller

import (
	"reflect"
	"testing"
	"time"

	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/protobuf/identitypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRoleToPb(t *testing.T) {
	time1, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	if err != nil {
		t.Fatalf("Failed to parse time: %v", err)
	}
	time2 := time.Now()

	type args struct {
		role *roles.Role
	}
	tests := []struct {
		name string
		args args
		want *identitypb.Role
	}{
		{
			"transforms domain entity to the gRPC message",
			args{
				&roles.Role{
					ID:        "1",
					CreatedAt: time1,
					UpdatedAt: time2,
				},
			},
			&identitypb.Role{
				Name:       "roles/1",
				CreateTime: timestamppb.New(time1),
				UpdateTime: timestamppb.New(time2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoleToPb(tt.args.role); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoleToPb() = %v, want %v", got, tt.want)
			}
		})
	}
}
