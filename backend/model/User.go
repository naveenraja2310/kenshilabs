package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"ID"`
	PhoneNumber int                `bson:"PhoneNumber" json:"PhoneNumber"`
	Password    string             `bson:"Password" json:"Password"`
	CreatedAt   time.Time          `bson:"CreatedAt" json:"CreatedAt"`
}
