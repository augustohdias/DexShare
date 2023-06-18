package entity

type UserEntity struct {
	ID        string   `bson:"id,omitempty" validate:"omitempty"`
	Name      string   `bson:"name" validate:"required"`
	SaveFiles []string `bson:"saveFiles,omitempty" validate:"omitempty"`
	Followers []string `bson:"followers,omitempty" validate:"omitempty"`
	Following []string `bson:"following,omitempty" validate:"omitempty"`
	Email     string   `bson:"email" validate:"required"`
	Password  string   `bson:"password" validate:"required"`
}
