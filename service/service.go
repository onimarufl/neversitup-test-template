package service

import (
	"encoding/json"
	"fmt"

	"github.com/onimarufl/neversitup-test-template/models"
	"github.com/onimarufl/neversitup-test-template/repository"
	"github.com/tidwall/gjson"
)

type Servicer interface {
	GetUserService(req models.Request) (*models.Response, error)
}

type service struct {
	repository repository.Repositoryer
}

func NewService(repository repository.Repositoryer) *service {
	return &service{
		repository: repository,
	}
}

func (s service) GetUserService(req models.Request) (*models.Response, error) {
	resp := models.Response{}
	byteValue, err := s.repository.GetUser()
	if err != nil {
		return nil, err
	}

	value := gjson.Get(string(byteValue), fmt.Sprintf("user.#(userId==%d)", req.UserID))
	if err = json.Unmarshal([]byte(value.String()), &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
