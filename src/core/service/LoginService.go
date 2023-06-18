package service

import (
	"crypto/rand"
	"dexshare/src/core/entity"
	"dexshare/src/infrastructure/repository"
	"dexshare/src/port"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginService struct {
	UserSessionRepository port.UserSessionRepositoryPort
	UserRepository        port.UserRepositoryPort
}

func DefaultLoginService() LoginService {
	return LoginService{
		UserSessionRepository: &repository.UserSessionRepository{},
		UserRepository:        &repository.UserRepository{},
	}
}

func (l *LoginService) Authenticate(email string, password string) (string, error) {
	user, repositoryErr := l.UserRepository.FindByEmail(email)
	if repositoryErr != nil {
		log.Println(repositoryErr)
		return "", repositoryErr
	}

	if user.Password != password {
		log.Println("Invalid password for user " + email)
		return "", errors.New("invalid password")
	}

	keySize := 32
	key := make([]byte, keySize)
	_, keyErr := rand.Read(key)
	if keyErr != nil {
		log.Println("Error creating secret key:", keyErr)
		return "", errors.New("couldnt create key")
	}
	secret := hex.EncodeToString(key)
	unixTime := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   unixTime,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, signErr := token.SignedString([]byte(secret))
	if signErr != nil {
		log.Println("Error signing secret key:", signErr)
		return "", errors.New("couldnt sign key")
	}
	l.UserSessionRepository.Save(entity.UserSessionEntity{
		UserID:         user.ID,
		Email:          user.Email,
		Key:            secret,
		ExpirationDate: unixTime,
	})

	return tokenString, nil
}
