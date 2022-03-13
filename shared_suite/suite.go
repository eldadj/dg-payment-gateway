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
	suite.Run(t, &TestSuite{})
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

//n = new just created,
//v = voided,
//r = refunded,
//p = capturing or when at least when capture has taken place
//c = fully captured
