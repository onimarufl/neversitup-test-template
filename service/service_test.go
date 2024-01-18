package service

import (
	"errors"
	"testing"

	mockservice "github.com/onimarufl/neversitup-test-template/mock_service"
	"github.com/onimarufl/neversitup-test-template/models"
	"github.com/onimarufl/neversitup-test-template/repository"
	"github.com/stretchr/testify/suite"
)

type ResultTestSuite struct {
	suite.Suite
	service    service
	repository repository.Repositoryer
}

func (suite *ResultTestSuite) SetupTest() {
	suite.repository = new(mockservice.RepositoryerMock)

	suite.service = service{
		repository: suite.repository,
	}

}

func TestResultTestSuite(t *testing.T) {
	suite.Run(t, new(ResultTestSuite))
}

func (suite *ResultTestSuite) TestNewService() {
	expected := &service{
		repository: suite.repository,
	}

	actual := NewService(
		suite.repository,
	)

	suite.Equal(expected, actual, "should be equal")
}

func (suite *ResultTestSuite) TestUnitServiceSuccess() {

	req := models.Request{
		UserID: 2,
	}

	dataMock := "{\"user\":[{\"userId\":1,\"firstname\":\"TestFirstname 1\",\"Lastname\":\"TestLastname 1\",\"userRole\":\"test\"},{\"userId\":2,\"firstname\":\"TestFirstname 2\",\"Lastname\":\"TestLastname 2\",\"userRole\":\"test\"},{\"userId\":3,\"firstname\":\"TestFirstname 3\",\"Lastname\":\"TestLastname 3\",\"userRole\":\"test\"}]}"
	suite.repository.(*mockservice.RepositoryerMock).On("GetUser").Return([]byte(dataMock), nil)

	resp, err := suite.service.GetUserService(req)
	suite.NoError(err)
	suite.Equal(resp.Firstname, "TestFirstname 2")

}

func (suite *ResultTestSuite) TestUnitServiceError() {

	req := models.Request{
		UserID: 2,
	}

	dataMock := "{\"user\":[{\"userId\":1,\"firstname\":\"TestFirstname 1\",\"Lastname\":\"TestLastname 1\",\"userRole\":\"test\"},{\"userId\":2,\"firstname\":\"TestFirstname 2\",\"Lastname\":\"TestLastname 2\",\"userRole\":\"test\"},{\"userId\":3,\"firstname\":\"TestFirstname 3\",\"Lastname\":\"TestLastname 3\",\"userRole\":\"test\"}]}"
	suite.repository.(*mockservice.RepositoryerMock).On("GetUser").Return([]byte(dataMock), errors.New("something..."))

	_, err := suite.service.GetUserService(req)
	suite.Error(err)

}
