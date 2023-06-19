package port

import "dexshare/src/core/entity"

type UserServicePort interface {
	Create(entity.UserEntity) (string, error)
	Read(string) (entity.UserEntity, error)
	UploadSaveFile(string, string) (entity.UserEntity, error)
	Delete(string) (entity.UserEntity, error)
}

type UserRepositoryPort interface {
	Save(user entity.UserEntity) (string, error)
	Find(id string) (entity.UserEntity, error)
	UpdatePokemons(user entity.UserEntity) (entity.UserEntity, error)
	FindByEmail(email string) (entity.UserEntity, error)
}
