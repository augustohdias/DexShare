package entity

type UserEntity struct {
	ID        string   `json:"id,omitempty" validate:"omitempty"`
	Name      string   `json:"name" validate:"required"`
	SaveFiles []string `json:"saveFiles,omitempty" validate:"omitempty"`
	Followers []string `json:"followers,omitempty" validate:"omitempty"`
	Following []string `json:"following,omitempty" validate:"omitempty"`
	Email     string   `json:"email" validate:"required"`
	Password  string   `json:"password" validate:"required"`
}
