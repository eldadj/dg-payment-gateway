package controllers_test

import (
	"github.com/eldadj/dgpg/controllers"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (ts *TestSuite) TestValidateAuthorizeCode() {
	type args struct {
		authorizeCode string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   error //  assert.ErrorAssertionFunc
		wantValue assert.ValueAssertionFunc
	}{
		{
			name:    "no authorize code",
			args:    args{authorizeCode: ""},
			wantErr: controllers.ErrAuthorizeCodeNotSpecified,
		},
		{
			name:    "invalid authorize code",
			args:    args{authorizeCode: "not_found_code"},
			wantErr: errors.ErrAuthorizeCodeNotFound,
		},
	}

	for _, tt := range tests {
		t := ts.T()
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr != nil {
				assert.ErrorIs(t, tt.wantErr, controllers.ValidateAuthorizeCode(tt.args.authorizeCode))
			}
		})
	}
}

//
func (ts *TestSuite) SetupTest() {
	models.InitDB()
	// add test values
}

func (ts *TestSuite) TearDownTest() {
	models.CloseDB()
	//remove test values
}
