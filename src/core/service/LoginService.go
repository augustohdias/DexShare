package service

import (
	"crypto/rand"
	"dexshare/src/core/entity"
	"dexshare/src/port"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginService struct {
	UserSessionRepository port.UserSessionRepositoryPort
	UserRepository        port.UserRepositoryPort
}

func DefaultLoginService(userSessionRepository port.UserSessionRepositoryPort, userRepository port.UserRepositoryPort) LoginService {
	return LoginService{
		UserSessionRepository: userSessionRepository,
		UserRepository:        userRepository,
	}
}

func (l *LoginService) Authenticate(userID string, token string) bool {
	userSession, err := l.UserSessionRepository.Find(userID)
	if err != nil {
		return false
	}

	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(userSession.Key), nil
	})

	if err != nil {
		log.Println("Error parsing token: ", err)
		return false
	}

	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok && tokenObj.Valid {
		if claims["id"] == userID && int64(claims["exp"].(float64)) > time.Now().Unix() {
			return true
		}
	}

	return false
}

func (l *LoginService) Login(email string, password string) (string, error) {
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
		log.Println("Error signing secret key: ", signErr)
		return "", errors.New("couldnt sign key")
	}
	_, err := l.UserSessionRepository.Save(entity.UserSessionEntity{
		UserID:         user.ID,
		Email:          user.Email,
		Key:            secret,
		ExpirationDate: unixTime,
	})
	if err != nil {
		log.Println("Error saving session:", signErr)
		return "", err
	}
	return tokenString, nil
}
