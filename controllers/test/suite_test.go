package controllers_test

import (
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
