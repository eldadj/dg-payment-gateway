package models_test

import (
	"github.com/eldadj/dgpg/dto/merchant/authenticate"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (ts *TestSuite) TestAuthenticate() {
	ts.CreateTestMerchants()
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantValue interface{}
	}{
		{
			name:      "username not set",
			wantErr:   true,
			wantValue: errors.ErrFieldValueRequired("username"),
			args:      args{username: "", password: "test"},
		},
		{
			name:      "password not set",
			wantErr:   true,
			wantValue: errors.ErrFieldValueRequired("password"),
			args:      args{username: "test", password: ""},
		},
		{
			name:      "invalid username",
			wantErr:   true,
			wantValue: errors.ErrMerchantAuthenticationFailed,
			args:      args{username: "merchant", password: "test"},
		},
		{
			name:      "invalid password",
			wantErr:   true,
			wantValue: errors.ErrMerchantAuthenticationFailed,
			args:      args{username: "m1", password: "test"},
		},
		{
			name:    "merchant authenticated",
			wantErr: false,
			args:    args{username: "tm1", password: "tm1"},
		},
	}

	for _, tt := range tests {
		t := ts.T()
		t.Run(tt.name, func(t *testing.T) {
			arg := tt.args
			req := authenticate.Request{
				Username: arg.username,
				Password: arg.password,
			}
			jwtToken, err := models.Authenticate(req)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, tt.wantValue.(error).Error())
			} else {
				assert.NotEmpty(t, jwtToken)
				assert.Nil(t, err)
			}
		})
	}
}
