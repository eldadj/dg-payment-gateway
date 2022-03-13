package payment_test

import (
	"github.com/eldadj/dgpg/internal/merchant"
	"github.com/eldadj/dgpg/models"
	_ "github.com/eldadj/dgpg/routers"
	"github.com/eldadj/dgpg/shared_suite"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type TestSuite struct {
	shared_suite.TestSuite
}

func Test(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (ts *TestSuite) SetupSuite() {
	models.InitDB()
}

func (ts *TestSuite) TearDownSuite() {
	ts.DeleteTestMerchants()
	models.CloseDB()
}

func (ts *TestSuite) ValidateMerchantUpdateRequestContext(r *http.Request) (*http.Request, error) {
	ctx := r.Context()
	err := merchant.Validate(&ctx, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtaWQiOjJ9.2UtNBvBcZJlwatvbiuFkFwWS7ZliHcIs7_ZxMTFt9sE")
	//update request with context since we are creating context here
	return r.WithContext(ctx), err
}
