package models_test

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/request"
	"github.com/eldadj/dgpg/internal/errors"
	capture2 "github.com/eldadj/dgpg/models/capture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (ts *TestSuite) TestDoCapture() {
	ts.CreateTestCreditCards()
	ts.CreateTestMerchants()
	ts.CreateTestAuthorizes()

	type args struct {
		authorizeCode string
		amount        float64
		merchantId    int64
	}
	tests := []struct {
		name      string
		arg       args
		wantErr   bool
		wantValue interface{}
		valueFunc assert.ValueAssertionFunc
	}{
		{
			name:      "invalid authorize code",
			arg:       args{authorizeCode: "invalid", amount: 200, merchantId: ts.ValidMerchantID()},
			wantErr:   true,
			wantValue: errors.ErrAuthorizeCodeNotFound,
		},
		{
			name:      "invalid amount",
			arg:       args{authorizeCode: ts.CaptureCanCaptureAuthorizeCode(), amount: 0, merchantId: ts.ValidMerchantID()},
			wantErr:   true,
			wantValue: errors.ErrCaptureAmount,
		},
		{
			name: "capture amount exceeds authorized amount",
			arg: args{
				authorizeCode: ts.CaptureCanCaptureAuthorizeCode(),
				amount:        1000,
				merchantId:    ts.ValidMerchantID(),
			},
			wantErr:   true,
			wantValue: errors.ErrCaptureAmountExceedsAuthorizeAmount,
		},
		{
			name:    "capture ok",
			arg:     args{authorizeCode: ts.CaptureCanCaptureAuthorizeCode(), amount: 10, merchantId: ts.ValidMerchantID()},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t := ts.T()
		t.Run(tt.name, func(t *testing.T) {
			req := request.Request{
				AuthorizeCode: payment.AuthorizeCode{Code: tt.arg.authorizeCode},
				Amount:        tt.arg.amount,
				Request:       payment.Request{MerchantId: tt.arg.merchantId},
			}
			resp, err := capture2.DoCapture(req)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, tt.wantValue.(error).Error())
			} else {
				assert.NotEmpty(t, resp)
			}
		})
	}

	//exceed amount
	ts.T().Run("capture amount exceeded", func(t *testing.T) {
		//ts.DeleteTestCaptures()
		ts.ResetCreditCardAmount()
		ts.CreateTestAuthorizes()
		//authAmount := 50
		creditCardAmount := 1000.0
		//credit card initial amount = 1000
		req := request.Request{
			AuthorizeCode: payment.AuthorizeCode{Code: ts.CaptureCanCaptureAuthorizeCode()},
			Amount:        10,
			Request:       payment.Request{MerchantId: ts.ValidMerchantID()},
		}
		resp, err := capture2.DoCapture(req)
		assert.Nil(t, err)
		assert.Equal(t, creditCardAmount-req.Amount, resp.Amount)

		creditCardAmount = resp.Amount
		req.Amount = 20
		resp, err = capture2.DoCapture(req)
		assert.Nil(t, err)
		assert.Equal(t, creditCardAmount-req.Amount, resp.Amount)

		req.Amount = 40
		resp, err = capture2.DoCapture(req)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.ErrCaptureAmountExceedsAuthorizeAmount.Error())
	})

}
