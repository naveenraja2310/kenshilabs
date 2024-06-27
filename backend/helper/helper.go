package helper

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckDataExistsByID(dataid string, collection *mongo.Collection) bool {
	formObjectId, formiderr := primitive.ObjectIDFromHex(dataid)
	if formiderr != nil {
		fmt.Println(formiderr)
	}
	filter := bson.M{"_id": formObjectId}
	var result map[string]interface{}
	collection.FindOne(context.TODO(), filter).Decode(&result)

	if result == nil || result["_id"].(primitive.ObjectID).IsZero() {
		return false
	} else {
		return true
	}
}

func CheckDataExistsByKeyValue(key string, value interface{}, collection *mongo.Collection) bool {
	filter := bson.M{key: value}
	var result map[string]interface{}
	collection.FindOne(context.TODO(), filter).Decode(&result)

	if result == nil || result["_id"].(primitive.ObjectID).IsZero() {
		return false
	} else {
		return true
	}
}

func Contains(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
