package entity

type UserEntity struct {
	ID        string   `bson:"id" validate:"omitempty"`
	Name      string   `bson:"name" validate:"required"`
	Email     string   `bson:"email" validate:"required"`
	Password  string   `bson:"password" validate:"required"`
	Pokemons  []string `bson:"pokemons" validate:"omitempty"`
	Followers []string `bson:"followers,omitempty" validate:"omitempty"`
	Following []string `bson:"following,omitempty" validate:"omitempty"`
}
