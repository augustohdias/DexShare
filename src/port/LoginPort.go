package port

import "dexshare/src/core/entity"

type LoginPort interface {
	Authenticate(string, string) (string, error)
}

type UserSessionRepositoryPort interface {
	Save(entity.UserSessionEntity) (string, error)
	Find(string) (entity.UserSessionEntity, error)
}
