package api_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/yakiza/BankZ/api"
	"testing"
)

type AccountHandlerTestSuite struct {
	suite.Suite

	handler api.AccountHandler
}

func (suite *AccountHandlerTestSuite) SetupSuite() {

}

func TestNewAccountHandler(t *testing.T) {

}
