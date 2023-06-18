package entity

type UserSessionEntity struct {
	UserID         string `bson:"userId"`
	Email          string `bson:"email"`
	Key            string `bson:"signatureKey"`
	ExpirationDate int64  `bson:"expirationDate"`
}
