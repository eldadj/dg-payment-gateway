package models_test

import (
	"testing"

	_ "github.com/eldadj/dgpg/routers"
	"github.com/eldadj/dgpg/shared_suite"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	shared_suite.TestSuite
}

func Test(ts *testing.T) {
	suite.Run(ts, &TestSuite{})
}
