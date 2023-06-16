package service

import (
	"dexshare/src/core/entity"
)

type UserService struct{}

func (us *UserService) Create() (entity.User, error) {
	return entity.User{}, nil
}

func (us *UserService) Read(id string) (entity.User, error) {
	return entity.User{}, nil
}

func (us *UserService) Update(id string) (entity.User, error) {
	return entity.User{}, nil
}

func (us *UserService) Delete(id string) (entity.User, error) {
	return entity.User{}, nil
}

func (us *UserService) Authenticate() (entity.User, error) {
	return entity.User{}, nil
}
