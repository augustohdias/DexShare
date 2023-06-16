package entity

type Pokemon struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	OwnerID           string `json:"ownerId"`
	NationalDexNumber int    `json:"nationalDexNumber"`
	Level             int    `json:"level"`
}