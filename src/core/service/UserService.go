package service

import (
	"dexshare/src/core/entity"
	"dexshare/src/domain/port"
	"dexshare/src/infrastructure/repository"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository port.UserRepositoryPort
}

func DefaultUserService() UserService {
	return UserService{UserRepository: &repository.UserRepository{}}
}

func (us *UserService) Create(user entity.UserEntity) (string, error) {
	user.Followers, user.Following, user.SaveFiles = []string{}, []string{}, []string{}
	user.ID = uuid.New().String()
	id, err := us.UserRepository.Save(user)
	if err != nil {
		return "", err
	}
	createdUser, err := us.UserRepository.Find(id)
	if err != nil {
		return "", err
	}
	return createdUser.ID, nil
}

func (us *UserService) Read(id string) (entity.UserEntity, error) {
	if user, err := us.UserRepository.Find(id); err != nil {
		return entity.UserEntity{}, err
	} else {
		return user, nil
	}
}

func (us *UserService) Update(id string) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}

func (us *UserService) Delete(id string) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}

func (us *UserService) Authenticate(credentials string) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}
