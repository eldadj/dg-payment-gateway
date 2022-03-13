package models_test

import (
	"context"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/authorize"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/internal/merchant"
	"github.com/eldadj/dgpg/models"
	authorize2 "github.com/eldadj/dgpg/models/authorize"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (ts *TestSuite) TestAddAuthorize() {
	ts.CreateTestMerchants()
	ts.CreateTestCreditCards()
	type args struct {
		merchantId   int64
		creditCardId int64
		currency     string
		amount       float64
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantValue interface{}
	}{
		{
			name:      "invalid merchant_id",
			wantErr:   true,
			wantValue: errors.ErrAuthorizeInvalidFieldValue("merchant_id"),
			args:      args{merchantId: -1, creditCardId: 1000000, currency: "USD", amount: 10},
		},
		{
			name:      "invalid credit_card_id",
			wantErr:   true,
			wantValue: errors.ErrAuthorizeInvalidFieldValue("credit_card_id"),
			args:      args{merchantId: 1000000, creditCardId: -1, currency: "USD", amount: 10},
		},
		{
			name:      "invalid currency",
			wantErr:   true,
			wantValue: errors.ErrAuthorizeInvalidFieldValue("currency"),
			args:      args{merchantId: 1000000, creditCardId: 1000000, currency: "", amount: 10},
		},
		{
			name:      "invalid amount",
			wantErr:   true,
			wantValue: errors.ErrAuthorizeInvalidFieldValue("amount"),
			args:      args{merchantId: 1000000, creditCardId: 1000000, currency: "USD", amount: 0},
		},
		{
			name: "authorize created",
			//wantValue: models.ErrAuthorizeInvalidFieldValue("amount"),
			args: args{merchantId: 1000000, creditCardId: 1000000, currency: "USD", amount: 100},
		},
	}

	for _, tt := range tests {
		t := ts.T()
		t.Run(tt.name, func(t *testing.T) {
			arg := tt.args
			ac := payment.AmountCurrency{Currency: arg.currency, Amount: arg.amount}
			tx := models.NewTx()
			resp, err := authorize2.Add(tx, arg.merchantId, arg.creditCardId, ac)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, tt.wantValue.(error).Error())
			} else {
				assert.NotEmpty(t, resp)
			}
		})
	}
}

func (ts *TestSuite) TestAuthorize() {
	ts.DeleteTestCreditCards()
	ts.CreateTestMerchants()

	ts.T().Run("test_authorize", func(t *testing.T) {
		ctx, _ := context.WithCancel(context.Background())
		//validate token so we have a merchantId stored
		err := merchant.Validate(&ctx, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOjF9.xhy07275jrBO0sGmIDAe4TwgVNrLgd146PSae3os3MI")
		assert.Nil(t, err)

		req := authorize.Request{
			CreditCard: payment.CreditCard{
				OwnerName: "test owner",
				Number:    "T234 E234 S234 T234 1234",
				ExpMonth:  10,
				ExpYear:   22,
				CVV:       "abc",
			},
			AmountCurrency: payment.AmountCurrency{
				Amount: 200, Currency: "USD",
			},
		}
		resp, err := authorize2.DoAuthorize(ctx, req)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
}

func (ts *TestSuite) TestGet() {
	ts.DeleteTestCreditCards()
	ts.CreateTestMerchants()

	ts.T().Run("test_get", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		//validate token so we have a merchantId stored
		// merchant id is 1
		merchant1Token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOjF9.xhy07275jrBO0sGmIDAe4TwgVNrLgd146PSae3os3MI"
		err := merchant.Validate(&ctx, merchant1Token)
		merchant1Id, valid := merchant.FromContext(ctx)
		assert.True(t, valid)
		assert.Nil(t, err)

		req := authorize.Request{
			CreditCard: payment.CreditCard{
				OwnerName: "test owner",
				Number:    "T234 E234 S234 T234 1234",
				ExpMonth:  10,
				ExpYear:   22,
				CVV:       "abc",
			},
			AmountCurrency: payment.AmountCurrency{
				Amount: 200, Currency: "USD",
			},
		}
		resp, err := authorize2.DoAuthorize(ctx, req)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)

		authorizeCode := resp.Code

		a, err := authorize2.Get(authorizeCode, merchant1Id)
		assert.Nil(t, err)
		assert.NotEmpty(t, a)

		//validate a new merchant. MerchantId == 2
		merchant2Token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOjJ9.2UtNBvBcZJlwatvbiuFkFwWS7ZliHcIs7_ZxMTFt9sE"
		err = merchant.Validate(&ctx, merchant2Token)
		assert.Nil(t, err)
		//we get the new merchantId
		merchant2Id, valid := merchant.FromContext(ctx)
		assert.True(t, valid)

		//both ids shouldn't be equal
		assert.NotEqual(t, merchant1Id, merchant2Id)

		a, err = authorize2.Get(authorizeCode, merchant2Id)
		//it's for merchant 1
		assert.EqualError(t, err, errors.ErrInvalidMerchantAuthorizeCode.Error())

		cancel()
	})
}
