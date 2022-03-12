// Package shared_suite shared testing suite to make setup/teardown easy

package shared_suite

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/eldadj/dgpg/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"path/filepath"
	"runtime"
	"testing"
)

type TestSuite struct {
	suite.Suite
}

// allow our conf file to be read during testing
func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	web.TestBeegoInit(apppath)
}

func Test(t *testing.T) {
	//beego.TestBeegoInit("/Users/eldadonojetah/Documents/git/sholagracetech/accesspower/powervending/integrations/buypower/go_buypower")
	//suite.Run(t, new(TestSuite))
	suite.Run(t, &TestSuite{})
}

func (ts *TestSuite) CreateTestCreditCards() {
	ts.DeleteTestCreditCards()
	models.ExecDBFunc(func(tx *gorm.DB) error {
		//merchantId: 1000000, creditCardId: 1000000, currency: "USD", amount: 100
		tx.Exec(`
insert into credit_card(credit_card_id, owner_name,card_no,exp_month,exp_year,cvv,currency_code,current_amount) values
(1000000,'Test Card 1', 'T123 E123 S123 T123 1234', 12, 22, 'CVV', 'USD', 1000),
(1000001,'Test Card 2', 'T234 E234 S234 T234 2345', 10, 22, 'CV2', 'EUR', 1500)`)
		return nil
	})
}

func (ts *TestSuite) DeleteTestCreditCards() {
	models.ExecDBFunc(func(tx *gorm.DB) error {
		tx.Exec(`delete from credit_card where credit_card_id >= 1000000 or card_no like 'T%'`)
		return nil
	})
}

func (ts *TestSuite) CreateTestMerchants() {
	ts.DeleteTestMerchants()
	models.ExecDBFunc(func(tx *gorm.DB) error {
		//merchantId: 1000000, creditCardId: 1000000, currency: "USD", amount: 100
		tx.Exec(`
insert into merchant(merchant_id, fullname, user_name, pwd_hash) values
(1000000,'Test Merchant 1', 'tm1', '$2a$12$lLBP7ylkt0pTp5EkeQQJwur5rtqB82LVvEvJC.qMr904EKFG/YjGy'),
(1000001,'Test Merchant 1', 'tm1', '$2a$12$SLbYUEu5h5cs7bPrW71lSebul0K/2/JB/0jDkVv1NS/NIWhukZTVy')`)
		return nil
	})
}

func (ts *TestSuite) DeleteTestMerchants() {
	models.ExecDBFunc(func(tx *gorm.DB) error {
		tx.Exec(`delete from merchant where merchant_id >= 1000000`)
		return nil
	})
}

func (ts *TestSuite) CreateTestAuthorizes() {
	ts.DeleteTestAuthorizes()
	models.ExecDBFunc(func(tx *gorm.DB) error {
		//merchantId: 1000000, creditCardId: 1000000, currency: "USD", amount: 100
		tx.Exec(`
insert into authorize(authorize_id, merchant_id, credit_card_id, currency,amount,authorize_code, status) values 
(1000000,1000000,1000000,'USD', 50, '10000001', 'n'),
(1000001,1000001,1000000,'USD', 150, '10000002', 'n'),
(1000002,1000000,1000000,'USD', 252, '10000003', 'n'),
(1000010,1000001,1000001,'EUR', 27, '10000011', 'n'),
(1000011,1000001,1000001,'EUR', 85, '10000012', 'n'),
(1000012,1000001,1000001,'EUR', 63, '10000013', 'n'),
(1000013,1000001,1000001,'EUR', 110, '10000014', 'n'),
(2000100,1000001,1000001,'EUR', 63, '11000013', 'v'),
(2000101,1000001,1000001,'EUR', 63, '11000014', 'r'),
(2000102,1000001,1000001,'EUR', 25, '11000015', 'p'),
(2000103,1000001,1000001,'EUR', 63, '11000016', 'c')

`)
		return nil
	})
}

func (ts *TestSuite) DeleteTestAuthorizes() {
	models.ExecDBFunc(func(tx *gorm.DB) error {
		tx.Exec(`delete from authorize where authorize_id >= 1000000 `)
		return nil
	})
}

func (ts *TestSuite) CreateTestCaptures() {
	ts.DeleteTestCaptures()
	models.ExecDBFunc(func(tx *gorm.DB) error {
		tx.Exec(`
insert into capture(capture_id, amount, authorize_id) values
(1000000,20, 1000000),
(1000001,10, 1000010),
(1000002,7, 1000010)
`)
		return nil
	})
}

func (ts *TestSuite) CreateTestRefundCaptures() {

	models.ExecDBFunc(func(tx *gorm.DB) error {
		err := tx.Exec(`
insert into authorize(authorize_id, merchant_id, credit_card_id, currency,amount,authorize_code, status) values
(3000000,1000000,1000000,'USD', 50, '30000001', 'p')`).Error
		//println(err)
		err = tx.Exec(`insert into capture(capture_id, amount, authorize_id) values
(3000000,20, 3000000),
(3000001,10, 3000000),
(3000002,20, 3000000)`).Error
		//println(err)
		//reduce card amount
		err = tx.Exec(`update credit_card set current_amount = current_amount - 50 where credit_card_id = 1000000`).Error
		//println(err)
		return err
	})

}

func (ts *TestSuite) DeleteTestCaptures() {
	models.ExecDBFunc(func(tx *gorm.DB) error {
		tx.Exec(`delete from capture where capture_id >= 1000000 or authorize_id >= 1000000`)
		return nil
	})
}

func (ts *TestSuite) ResetCreditCardAmount() {
	models.ExecDBFunc(func(tx *gorm.DB) error {
		tx.Exec(`update credit_card set current_amount = 1000 where credit_card_id = 1000000`)
		return nil
	})
}

func (ts *TestSuite) ResetCreditCardAmountForRefund() {
	models.ExecDBFunc(func(tx *gorm.DB) error {
		tx.Exec(`update credit_card set current_amount = 950 where credit_card_id = 1000000`)
		return nil
	})
}

func (ts *TestSuite) AuthorizeCodeAlreadyVoided() string {
	return "11000013"
}

func (ts *TestSuite) AuthorizeCodeCanBeVoided() string {
	return "10000001"
}

func (ts *TestSuite) AuthorizeCodeCannotBeRefunded() string {
	return "11000013"
}

func (ts *TestSuite) CaptureCanCaptureAuthorizeCode() string {
	return "10000001"
}

func (ts *TestSuite) CaptureCannotCaptureAuthorizeCode() string {
	return "11000016"
}

func (ts *TestSuite) CaptureAuthorizedAmountExceedsAuthorizeCode() string {
	return "1000010"
}

func (ts *TestSuite) RefundAmountInvalidAuthorizeCode() string {
	return "11000015"
}

//n = new just created,
//v = voided,
//r = refunded,
//p = capturing or when at least when capture has taken place
//c = fully captured
