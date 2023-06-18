package port

import "dexshare/src/core/entity"

type LoginPort interface {
	Login(email string, pass string) (string, error)
	Authenticate(uid string, token string) bool
}

type UserSessionRepositoryPort interface {
	Save(user entity.UserSessionEntity) (string, error)
	Find(userId string) (entity.UserSessionEntity, error)
}
