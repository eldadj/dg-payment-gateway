// Package shared_suite shared testing suite to make setup/teardown easy

package shared_suite

import (
	"context"
	"github.com/beego/beego/v2/server/web"
	"github.com/eldadj/dgpg/dto/payment"
	"github.com/eldadj/dgpg/dto/payment/authorize"
	"github.com/eldadj/dgpg/dto/payment/request"
	"github.com/eldadj/dgpg/dto/payment/response"
	"github.com/eldadj/dgpg/internal/errors"
	"github.com/eldadj/dgpg/internal/merchant"
	"github.com/eldadj/dgpg/models"
	authorize2 "github.com/eldadj/dgpg/models/authorize"
	capture2 "github.com/eldadj/dgpg/models/capture"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"math/rand"
	"path/filepath"
	"runtime"
	"testing"
	"time"
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
(1000000,'Test Merchant 1', 'tm1', '$2a$12$l6pwgwg0vCknO5heebp/Ze/5FRnC7JD/8Tp.j/tY.sbHJ2nu6Lf3m'),
(1000001,'Test Merchant 2', 'tm2', '$2a$12$GjTxK/2.t.mlHLcP6OjAVunxM/U1LI7NWrwItMSVwUF0EukCvMzbe')`)
		return nil
	})
}

func (ts *TestSuite) DeleteTestMerchants() {
	models.ExecDBFunc(func(tx *gorm.DB) error {
		tx.Exec(`delete from merchant where merchant_id >= 1000000`)
		return nil
	})
}

func (ts *TestSuite) AuthoriseTestCreateUSDAuthorize() (string, float64, error) {
	ctx, cancel := context.WithCancel(context.Background())
	//validate token so we have a merchantId stored
	merchant.Validate(&ctx, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOjJ9.2UtNBvBcZJlwatvbiuFkFwWS7ZliHcIs7_ZxMTFt9sE")
	rand.Seed(time.Now().UnixNano())
	//min := 100
	//max := 1000
	authorizeAmount := float64(200) //rand.Intn((max-min)+min) * 1.0)
	req := authorize.Request{
		CreditCard: payment.CreditCard{
			OwnerName: "eldad onojetah",
			Number:    "4035 5010 0000 0008",
			ExpMonth:  10,
			ExpYear:   22,
			CVV:       "TTT",
		},
		AmountCurrency: payment.AmountCurrency{
			Amount: authorizeAmount, Currency: "USD",
		},
	}
	resp, err := authorize2.DoAuthorize(ctx, req)
	cancel()

	return resp.Code, authorizeAmount, err
}

func (ts *TestSuite) CaptureTestCreate200USDAuthorizeCapture10N50USD() (string, response.Response, error) {
	authorizeCode, authorizeAmount, err := ts.AuthoriseTestCreateUSDAuthorize()
	//capture 10USD
	req := request.Request{
		AuthorizeCode: payment.AuthorizeCode{Code: authorizeCode},
		Amount:        10,
	}
	resp, err := capture2.DoCapture(req)

	//capture  50USD
	req = request.Request{
		AuthorizeCode: payment.AuthorizeCode{Code: authorizeCode},
		Amount:        50,
	}
	resp, err = capture2.DoCapture(req)
	if resp.Amount != authorizeAmount-50-10 {
		return "", resp, errors.ErrCapture
	}
	return authorizeCode, resp, err
}

func (ts *TestSuite) VoidTestCreate200USDAuthorizeCapture20USD() (string, error) {
	authorizeCode, _, err := ts.AuthoriseTestCreateUSDAuthorize()
	//capture 10USD
	req := request.Request{
		AuthorizeCode: payment.AuthorizeCode{Code: authorizeCode},
		Amount:        20,
	}
	_, err = capture2.DoCapture(req)
	return authorizeCode, err
}

//n = new just created,
//v = voided,
//r = refunded,
//p = capturing or when at least when capture has taken place
//c = fully captured
