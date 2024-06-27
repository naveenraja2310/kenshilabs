package database

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"kenshilabs_demo/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitConnection() {

	cfg, confiferr := config.LoadConfig()
	if confiferr != nil {
		fmt.Println("Error loading configuration:", confiferr)
		return
	}

	connectionstr := fmt.Sprintf("mongodb://%s:%s", cfg.Database.Host, strconv.Itoa(cfg.Database.Port))
	clientOptions := options.Client().ApplyURI(connectionstr)

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func GetClientDb() *mongo.Client {
	return client
}

func GetDataBase() string {
	cfg, confiferr := config.LoadConfig()
	if confiferr != nil {
		fmt.Println("Error loading configuration:", confiferr)
	}

	return cfg.Database.Name
}
