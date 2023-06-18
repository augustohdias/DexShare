package port

import "dexshare/src/core/entity"

type UserServicePort interface {
	Create(entity.UserEntity) (string, error)
	Read(string) (entity.UserEntity, error)
	Update(string) (entity.UserEntity, error)
	Delete(string) (entity.UserEntity, error)
}

type UserRepositoryPort interface {
	Save(entity.UserEntity) (string, error)
	Find(string) (entity.UserEntity, error)
	FindByEmail(string) (entity.UserEntity, error)
}
