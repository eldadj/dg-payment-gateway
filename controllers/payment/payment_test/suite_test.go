package payment_test

import (
	"github.com/eldadj/dgpg/models"
	_ "github.com/eldadj/dgpg/routers"
	"github.com/eldadj/dgpg/shared_suite"
	"github.com/stretchr/testify/suite"
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
	ts.DeleteTestCaptures()
	ts.DeleteTestAuthorizes()
	ts.DeleteTestMerchants()
	ts.DeleteTestCreditCards()
	models.CloseDB()
}
