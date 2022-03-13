package models_test

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/void"
	"github.com/eldadj/dgpg/internal/errors"
	void2 "github.com/eldadj/dgpg/models/void"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (ts *TestSuite) TestVoid() {
	ts.CreateTestCreditCards()
	ts.CreateTestMerchants()
	ts.CreateTestAuthorizes()
	type args struct {
		authorizeCode string
		merchantId    int64
	}
	tests := []struct {
		name string
		args
		wantErr   bool
		wantValue interface{}
	}{
		{
			name:      "invalid authorize code",
			args:      args{authorizeCode: "invalid", merchantId: ts.ValidMerchantID()},
			wantErr:   true,
			wantValue: errors.ErrAuthorizeCodeNotFound,
		},
		{
			name:      "cannot void authorize code",
			args:      args{authorizeCode: ts.AuthorizeCodeAlreadyVoided(), merchantId: 1000001},
			wantErr:   true,
			wantValue: errors.ErrAuthorizeCannotVoid,
		},
		{
			name:      "authorize code voided",
			args:      args{authorizeCode: ts.AuthorizeCodeCanBeVoided(), merchantId: 1000000},
			wantErr:   false,
			wantValue: ts.AuthorizeCodeCanBeVoided(),
		},
	}

	for _, tt := range tests {
		t := ts.T()
		t.Run(tt.name, func(t *testing.T) {
			req := void.Request{
				AuthorizeCode: payment.AuthorizeCode{Code: tt.args.authorizeCode},
				Request:       payment.Request{MerchantId: tt.args.merchantId},
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

	ts.DeleteTestAuthorizes()
	ts.DeleteTestMerchants()
	ts.DeleteTestCreditCards()
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
