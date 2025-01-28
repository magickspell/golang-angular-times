package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type Schedule struct {
	Day   string `bson:"day"`
	Start string `bson:"start"`
	End   string `bson:"end"`
}

var (
	mongoClient         *mongo.Client
	MONGO_ROOT_USER     = "root"
	MONGO_ROOT_PASSWORD = "rootpass"
)

func InitDB() {
	// todo нужно проверить что данных нет в БД или удалить все и вставить данные заного
	// todo нужно подрубить ENV
	clientOptions := options.Client().ApplyURI("mongodb://root:rootpass@mongo:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	mongoClient = client

	// Initialize Users
	users := []User{
		{Email: "user1@some.com", Password: "user1@some.com"},
		{Email: "user2@some.com", Password: "user2@some.com"},
	}
	collection := mongoClient.Database("app").Collection("users")
	for _, user := range users {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Println(err)
		}
	}

	// Initialize Schedule
	schedule := []Schedule{
		{Day: "Monday", Start: "00:00", End: "24:00"},
		{Day: "Tuesday", Start: "00:00", End: "24:00"},
		{Day: "Wednesday", Start: "00:00", End: "24:00"},
		{Day: "Thursday", Start: "00:00", End: "24:00"},
		{Day: "Friday", Start: "00:00", End: "24:00"},
		{Day: "Saturday", Start: "00:00", End: "24:00"},
		{Day: "Sunday", Start: "00:00", End: "24:00"},
	}
	scheduleCollection := mongoClient.Database("app").Collection("schedule")
	for _, sch := range schedule {
		_, err := scheduleCollection.InsertOne(context.TODO(), sch)
		if err != nil {
			log.Println(err)
		}
	}
}
