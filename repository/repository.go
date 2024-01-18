package repository

import (
	"io/ioutil"
)

type Repositoryer interface {
	GetUser() ([]byte, error)
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r repository) GetUser() ([]byte, error) {
	byteValue, err := ioutil.ReadFile("./mock_data/mock_data.json")
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}
