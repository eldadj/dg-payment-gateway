package payment_test

import (
	"context"
	payment2 "github.com/eldadj/dgpg/controllers/payment"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/internal/merchant"
	"github.com/stretchr/testify/assert"
	"testing"
)

func (s *TestSuite) TestSetMerchantId() {
	s.T().Run("no_merchant_set_in_context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		req := payment.Request{}
		//assert.Equal(t, int6(0), req.MerchantId)
		err := payment2.SetMerchantId(ctx, &req)
		assert.EqualError(t, errors.ErrMerchantLoad, err.Error())
		cancel()
	})

	s.T().Run("merchant_set_in_context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		//validate token so we have a merchantId stored
		//merchant_id should be 1 and never expires
		err := merchant.Validate(&ctx, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOjF9.xhy07275jrBO0sGmIDAe4TwgVNrLgd146PSae3os3MI")
		assert.Nil(t, err)
		req := payment.Request{}
		err = payment2.SetMerchantId(ctx, &req)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), req.MerchantId)

		cancel()
		//ctx, _ := context.WithCancel(context.Background())
		////validate token so we have a merchantId stored
		//err := merchant.Validate(&ctx, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDcwMTUzNTYsIm1pZCI6MX0.pSkzl_WCU0VRF07sJwtlgeeHfatwDZNtqCJqIgQqi0Q")
		//assert.Nil(t, err)
		//
		//req := authorize.Request{
		//	CreditCard: payment.CreditCard{
		//		OwnerName: "test owner",
		//		Number:    "T234 E234 S234 T234 1234",
		//		ExpMonth:  10,
		//		ExpYear:   22,
		//		CVV:       "abc",
		//	},
		//	AmountCurrency: payment.AmountCurrency{
		//		Amount: 200, Currency: "USD",
		//	},
		//}
		//resp, err := authorize2.DoAuthorize(ctx, req)
		//assert.Nil(t, err)
		//assert.NotEmpty(t, resp)
	})
}
