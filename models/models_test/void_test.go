package models_test

import (
	"testing"

	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/void"
	"github.com/eldadj/dgpg/internal/errors"
	void2 "github.com/eldadj/dgpg/models/void"
	"github.com/stretchr/testify/assert"
)

func (ts *TestSuite) TestVoid() {
	ts.CreateTestMerchants()
	authorizeCodeWithCapture, err := ts.VoidTestCreate200USDAuthorizeCapture20USD()
	assert.Nil(ts.T(), err)
	authorizeCodeCanVoid, _, err := ts.AuthoriseTestCreateUSDAuthorize()

	tests := []struct {
		name          string
		authorizeCode string
		wantErr       bool
		wantValue     interface{}
	}{
		{
			name:          "invalid authorize code",
			authorizeCode: "invalid",
			wantErr:       true,
			wantValue:     errors.ErrAuthorizeCodeNotFound,
		},
		{
			name:          "cannot void authorize code",
			authorizeCode: authorizeCodeWithCapture,
			wantErr:       true,
			wantValue:     errors.ErrAuthorizeCannotVoid,
		},
		{
			name:          "authorize code voided",
			authorizeCode: authorizeCodeCanVoid,
			wantErr:       false,
			wantValue:     authorizeCodeCanVoid,
		},
	}

	for _, tt := range tests {
		t := ts.T()
		t.Run(tt.name, func(t *testing.T) {
			req := void.Request{
				AuthorizeCode: payment.AuthorizeCode{Code: tt.authorizeCode},
				Request:       payment.Request{MerchantId: 2},
			}
			//ctx, cancel := context.WithCancel(context.Background())
			resp, err := void2.Void(req)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, tt.wantValue.(error).Error())
			} else {
				assert.NotEmpty(t, resp)
				//assert.Equal(t, tt.wantValue.(string), resp)
			}
			//cancel()
		})
	}
}

/*
func (ts *TestSuite) TestVoidInvalidAuthorizeCode() {
	ts.T().Run("void_invalid_authorize_code", func(t *testing.T) {
		req := void.Request{
			AuthorizeCode: payment.AuthorizeCode{Code: "123"},
		}
		ctx, cancel := context.WithCancel(context.Background())
		_, err := void2.Void(ctx, req)
		assert.NotNil(t, err)
		cancel()
	})
}
func (ts *TestSuite) TestVoidCannotVoid() {
	ts.T().Run("void_cannot_void", func(t *testing.T) {
		req := void.Request{
			AuthorizeCode: payment.AuthorizeCode{Code: ts.AuthorizeCodeAlreadyVoided()},
		}
		ctx, cancel := context.WithCancel(context.Background())
		_, err := void2.Void(ctx, req)
		assert.NotNil(t, err)
		cancel()
	})
}
*/
