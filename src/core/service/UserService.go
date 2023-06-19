package service

import (
	"dexshare/src/core/entity"
	"dexshare/src/core/parser"
	"dexshare/src/port"
	"encoding/base64"
	"log"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepository    port.UserRepositoryPort
	PokemonRepository port.PokemonRepositoryPort
}

func DefaultUserService(userRepository port.UserRepositoryPort, pokemonRepository port.PokemonRepositoryPort) UserService {
	return UserService{
		UserRepository:    userRepository,
		PokemonRepository: pokemonRepository,
	}
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
	if err != nil {
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

	teamPokemons := saveData.Team.Pokemons
	pcPokemons := saveData.PC.Pokemons
	allPokemons := make([]entity.PokemonEntity, len(teamPokemons) + len(pcPokemons))
	copy(allPokemons, teamPokemons)
	copy(allPokemons[len(teamPokemons):], pcPokemons)

	var pokemonIDs []string
	for _, pokemon := range allPokemons {
		pokemon.ID = uuid.NewString()
		pokemon.OwnerID = userID
		_, err := us.PokemonRepository.Save(pokemon)
		if err != nil {
			log.Println("Failed to update pokemons.")
			return entity.UserEntity{}, err
		}
		pokemonIDs = append(pokemonIDs, pokemon.ID)
	}
	for _, pokemon := range user.Pokemons {
		us.PokemonRepository.Delete(pokemon)
	}
	user.Pokemons = pokemonIDs
	updatedUser, err := us.UserRepository.UpdatePokemons(user)
	if err != nil {
		log.Println("Couldn't update user.")
		return entity.UserEntity{}, err
	}
	return updatedUser, nil
}

func (us *UserService) Delete(id string) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}
