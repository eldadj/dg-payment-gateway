package models_test

import (
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/request"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/models/capture"
	"github.com/eldadj/dgpg/models/refund"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (ts *TestSuite) TestDoRefund() {
	ts.CreateTestCaptures()
	ts.CreateTestAuthorizes()
	ts.CreateTestMerchants()
	ts.CreateTestCreditCards()
	ts.ResetCreditCardAmountForRefund()

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
			arg:       args{authorizeCode: "invalid", amount: 200},
			wantErr:   true,
			wantValue: errors.ErrAuthorizeCodeNotFound,
		},
		{
			name:      "cannot refund authorize code",
			arg:       args{authorizeCode: ts.AuthorizeCodeCannotBeRefunded(), amount: 0},
			wantErr:   true,
			wantValue: errors.ErrAuthorizeCannotRefund,
		},
		{
			name:      "refund amount not found",
			arg:       args{authorizeCode: ts.RefundAmountInvalidAuthorizeCode(), amount: 1000},
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
	ts.CreateTestMerchants()
	ts.CreateTestAuthorizes()
	ts.CreateTestCreditCards()
	ts.CreateTestRefundCaptures()
	ts.T().Run("refund ok", func(t *testing.T) {
		ts.ResetCreditCardAmountForRefund()
		creditCardAmount := 950.0
		authorizeCode := "30000001"
		//credit card initial amount = 1000
		req := request.Request{
			AuthorizeCode: payment.AuthorizeCode{Code: authorizeCode},
			Amount:        10,
		}
		resp, err := refund.DoRefund(req)
		assert.Nil(t, err)
		//card amount should increase by refund value
		assert.Equal(t, creditCardAmount+req.Amount, resp.Amount)

		creditCardAmount = resp.Amount
		req.Amount = 20
		resp, err = refund.DoRefund(req)
		assert.Nil(t, err)
		assert.Equal(t, creditCardAmount+req.Amount, resp.Amount)
		creditCardAmount = resp.Amount

		req.Amount = 30
		resp, err = refund.DoRefund(req)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.ErrRefundAmount.Error())

		_, err = capture.DoCapture(req)
		assert.NotNil(t, err)
		assert.EqualError(t, err, errors.ErrAuthorizeCannotCapture.Error())

		//creditCardAmount = resp.Amount
		req.Amount = 20
		resp, err = refund.DoRefund(req)
		assert.Nil(t, err)
		assert.Equal(t, creditCardAmount+req.Amount, resp.Amount)
	})
}
