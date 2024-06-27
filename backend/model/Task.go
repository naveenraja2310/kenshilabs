package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	TaskID        primitive.ObjectID `bson:"_id" json:"TaskID"`
	TaskName      string             `bson:"TaskName" json:"TaskName"`
	TaskAssinedTo string             `bson:"TaskAssinedTo" json:"TaskAssinedTo"`
	TaskAssinedBy string             `bson:"TaskAssinedBy" json:"TaskAssinedBy"`
	CreatedAt     time.Time          `bson:"CreatedAt" json:"CreatedAt"`
	UpdatedAt     time.Time          `bson:"UpdatedAt" json:"UpdatedAt"`
}
