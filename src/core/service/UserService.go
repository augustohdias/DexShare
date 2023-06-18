package service

import (
	"dexshare/src/core/entity"
	"dexshare/src/core/parser"
	"dexshare/src/infrastructure/repository"
	"dexshare/src/port"
	"encoding/base64"
	"log"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository port.UserRepositoryPort
	PokemonRepository port.PokemonRepositoryPort
}

func DefaultUserService() UserService {
	return UserService{UserRepository: &repository.UserRepository{}}
}

func (us *UserService) Create(user entity.UserEntity) (string, error) {
	user.Followers, user.Following = []string{}, []string{}
	user.ID = uuid.NewString()
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
	user, err := us.UserRepository.Find(id)
	if  err != nil {
		return entity.UserEntity{}, err
	}
	return user, nil
}

func (us *UserService) UploadSaveFile(userID string, data string) (entity.UserEntity, error) {
	user, err := us.UserRepository.Find(userID)
	log.SetPrefix("[userID=" + userID + "] ")
	if err != nil {
		log.Println("Couldn't find user.")
		return entity.UserEntity{}, err
	}
	decodedSaveFile, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Println("Couldn't decode file.")
		return entity.UserEntity{}, err
	}
	saveData := parser.LoadSaveFile(decodedSaveFile)
	var pokemonIDs []string
	for _, pokemon := range saveData.Team.Pokemons {
		pokemon.ID = uuid.NewString()
		_, err := us.PokemonRepository.Save(pokemon)
		if err != nil {
			log.Println("Failed to update pokemons.")
			return entity.UserEntity{}, err
		}
		pokemonIDs = append(pokemonIDs, pokemon.ID)
	}
	user.Pokemons = pokemonIDs
	_, err = us.UserRepository.Save(user)
	if err != nil {
		log.Println("Couldn't update user.")
		return entity.UserEntity{}, err
	}
	return entity.UserEntity{}, nil
}

func (us *UserService) Delete(id string) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}
