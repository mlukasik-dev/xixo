package roles

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestNewFilter(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    *Filter
		wantErr bool
	}{
		{
			"parses query correctly 1",
			args{query: "admin_only=true"},
			&Filter{AdminOnly: sql.NullBool{Bool: true, Valid: true}},
			false,
		},
		{
			"parses query correctly 2",
			args{query: "admin_only=false"},
			&Filter{AdminOnly: sql.NullBool{Bool: false, Valid: true}},
			false,
		},
		{
			"parses query correctly 3",
			args{query: "admin_only"},
			&Filter{AdminOnly: sql.NullBool{Bool: true, Valid: true}},
			false,
		},
		{
			"error 1",
			args{query: "admin_only=eljf"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFilter(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
