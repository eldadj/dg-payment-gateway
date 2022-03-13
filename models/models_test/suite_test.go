package models_test

import (
	"context"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/authorize"
	"github.com/eldadj/dgpg/internal/merchant"
	"github.com/eldadj/dgpg/models"
	authorize2 "github.com/eldadj/dgpg/models/authorize"
	_ "github.com/eldadj/dgpg/routers"
	"github.com/eldadj/dgpg/shared_suite"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type TestSuite struct {
	shared_suite.TestSuite
}

func Test(ts *testing.T) {
	suite.Run(ts, &TestSuite{})
}

//
func (ts *TestSuite) SetupSuite() {
	models.InitDB()
	// add test values
}

func (ts *TestSuite) TearDownSuite() {
	ts.DeleteTestMerchants()
	models.ExecDBFunc(func(tx *gorm.DB) error {
		tx.Exec(`delete from credit_card`)
		tx.Exec(`delete from authorize`)
		tx.Exec(`delete from merchant`)
		tx.Exec(`delete from refund`)
		tx.Exec(`delete from capture`)
		return nil
	})
	models.CloseDB()
}

func (ts *TestSuite) AfterTest(suiteName, testName string) {
	ts.DeleteTestMerchants()
}

func (ts *TestSuite) Create200USDAuthorize() (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	//validate token so we have a merchantId stored
	merchant.Validate(&ctx, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOjJ9.2UtNBvBcZJlwatvbiuFkFwWS7ZliHcIs7_ZxMTFt9sE")

	req := authorize.Request{
		CreditCard: payment.CreditCard{
			OwnerName: "eldad onojetah",
			Number:    "4035 5010 0000 0008",
			ExpMonth:  10,
			ExpYear:   22,
			CVV:       "TTT",
		},
		AmountCurrency: payment.AmountCurrency{
			Amount: 200, Currency: "USD",
		},
	}
	resp, err := authorize2.DoAuthorize(ctx, req)
	cancel()

	return resp.Code, err
}
