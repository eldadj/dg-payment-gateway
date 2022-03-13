package merchant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPasswordHashValid(t *testing.T) {
	//enable jwt validate to use key
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name      string
		args      args
		wantValue interface{}
	}{
		{
			name:      "password blank",
			args:      args{password: "", hash: "123"},
			wantValue: false,
		},
		{
			name:      "hash blank",
			args:      args{password: "123", hash: ""},
			wantValue: false,
		},
		{
			name:      "password is invalid",
			args:      args{password: "123", hash: "$2a$12$AXgSeJyIRN2ppJ.e7WDLUOHIVAWIcgNOLIybD9tPbHllIB4HJaLqW"},
			wantValue: false,
		},
		{
			name:      "hash is invalid",
			args:      args{password: "merchant1", hash: "$2a$12$AXgSeJyIRN2ppJ.e7WDLUOHIVAWIcgNOLIybD9tPbHllIB4HJa_invalid"},
			wantValue: false,
		},
		{
			name:      "password is valid",
			args:      args{password: "merchant1", hash: "$2a$12$AXgSeJyIRN2ppJ.e7WDLUOHIVAWIcgNOLIybD9tPbHllIB4HJaLqW"},
			wantValue: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := IsPasswordHashValid(tt.args.password, tt.args.hash)
			assert.Equal(t, tt.wantValue, valid)
		})
	}
}
