package models_test

import (
	"github.com/eldadj/dgpg/models"
	_ "github.com/eldadj/dgpg/routers"
	"github.com/eldadj/dgpg/shared_suite"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"strings"
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
	ts.DeleteTestCaptures()
	ts.DeleteTestAuthorizes()
	ts.DeleteTestMerchants()
	ts.DeleteTestCreditCards()
	models.CloseDB()
}

func (ts *TestSuite) BeforeTest(suiteName, testName string) {
	if strings.EqualFold(testName, "TestDoCapture") || strings.EqualFold(testName, "TestDoRefund") {

		models.ExecDBFunc(func(tx *gorm.DB) error {
			err := tx.Exec(`
insert into authorize(authorize_id, merchant_id, credit_card_id, currency,amount,authorize_code, status) values
(3000000,1000000,1000000,'USD', 50, '30000001', 'p')`).Error
			//println(err)
			err = tx.Exec(`insert into capture(capture_id, amount, authorize_id) values
(3000000,20, 3000000 /* '10000001' */),
(3000001,10, 3000000 /* '10000001' */),
(3000002,20, 3000000 /* '10000001' */)`).Error
			//println(err)
			//reduce card amount
			err = tx.Exec(`update credit_card set current_amount = current_amount - 50 where credit_card_id = 1000000`).Error
			//println(err)
			return err
		})
	}
}

func (ts *TestSuite) AfterTest(suiteName, testName string) {
	ts.DeleteTestCaptures()
	ts.DeleteTestAuthorizes()
	ts.DeleteTestMerchants()
	ts.DeleteTestCreditCards()
	ts.ResetCreditCardAmount()
}

func (ts *TestSuite) ValidMerchantID() int64 {
	return 1000000
}
