package port

import "dexshare/src/core/entity"

type UserServicePort interface {
	Create() (entity.User, error)
	Read(id string) (entity.User, error)
	Update(id string) (entity.User, error)
	Delete(id string) (entity.User, error)
	Athenticate(credentials string) (entity.User, error)
}

type UserDatabasePort interface {
	Save(entity.User) (string, error)
}
