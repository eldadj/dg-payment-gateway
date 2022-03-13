package models_test

import (
	"testing"

	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/request"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models/capture"
	"github.com/eldadj/dgpg/models/refund"
	"github.com/stretchr/testify/assert"
)

func (ts *TestSuite) TestDoRefund() {
	ts.CreateTestMerchants()
	validAuthorizeCode, _, err := ts.AuthoriseTestCreateUSDAuthorize()
	assert.Nil(ts.T(), err)

	type args struct {
		authorizeCode string
		amount        float64
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
			arg:       args{authorizeCode: "invalid", amount: 20},
			wantErr:   true,
			wantValue: errors.ErrAuthorizeCodeNotFound,
		},
		{
			name:      "amount cannot be refuned",
			arg:       args{authorizeCode: validAuthorizeCode, amount: 0},
			wantErr:   true,
			wantValue: errors.ErrRefundAmount,
		},
		{
			name:      "refund amount not found",
			arg:       args{authorizeCode: validAuthorizeCode, amount: 1000},
			wantErr:   true,
			wantValue: errors.ErrRefundAmount,
		},
	}

	for _, tt := range tests {
		t := ts.T()
		t.Run(tt.name, func(t *testing.T) {
			req := request.Request{
				AuthorizeCode: payment.AuthorizeCode{Code: tt.arg.authorizeCode},
				Amount:        tt.arg.amount,
			}
			resp, err := refund.DoRefund(req)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, tt.wantValue.(error).Error())
			} else {
				assert.NotEmpty(t, resp)
			}
		})
	}

	//test refund ok
	ts.T().Run("refund ok", func(t *testing.T) {
		authorizeCode, resp, err := ts.CaptureTestCreate200USDAuthorizeCapture10N50USD()
		assert.Nil(t, err)

		authorizeAmountBalance := resp.Amount
		req := request.Request{
			AuthorizeCode: payment.AuthorizeCode{Code: authorizeCode},
			Amount:        10,
		}
		//will refund
		resp, err = refund.DoRefund(req)
		assert.Nil(t, err)
		authorizeAmountBalance = authorizeAmountBalance + req.Amount
		assert.Equal(t, authorizeAmountBalance, resp.Amount)

		//cannot refund since not captured
		req.Amount = 30
		resp, err = refund.DoRefund(req)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.ErrRefundAmount.Error())

		//will refund
		req.Amount = 50
		resp, err = refund.DoRefund(req)
		assert.Nil(t, err)
		assert.Equal(t, authorizeAmountBalance+req.Amount, resp.Amount)

		//already done a refund, cannot capture
		_, err = capture.DoCapture(req)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.ErrAuthorizeCannotCapture.Error())
	})
}
