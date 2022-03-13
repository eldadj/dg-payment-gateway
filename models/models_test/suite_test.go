package models_test

import (
	"testing"

	"github.com/eldadj/dgpg/models"
	_ "github.com/eldadj/dgpg/routers"
	"github.com/eldadj/dgpg/shared_suite"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
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
		tx.Exec(`delete from merchant where user_name like 'tm*'`)
		tx.Exec(`delete from refund`)
		tx.Exec(`delete from capture`)
		return nil
	})
	models.CloseDB()
}

func (ts *TestSuite) AfterTest(suiteName, testName string) {
	ts.DeleteTestMerchants()
}
