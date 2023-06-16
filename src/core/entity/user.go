package entity

type User struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	SaveFiles []string `json:"saveFiles"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
}
