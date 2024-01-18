package mockservice

import "github.com/stretchr/testify/mock"

type RepositoryerMock struct {
	mock.Mock
}

func (s *RepositoryerMock) GetUser() ([]byte, error) {
	args := s.Called()
	return args.Get(0).([]byte), args.Error(1)
}
