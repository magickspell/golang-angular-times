package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Migrations struct {
	IsMigrated bool
}

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
	MongoClient         *mongo.Client
	MONGO_ROOT_USER     = "root"
	MONGO_ROOT_PASSWORD = "rootpass"
)

func checkMigrations() bool {
	collection := MongoClient.Database("app").Collection("migrations")

	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("[checkMigrations][Error][counting documents]", err)
		return true
	}
	if count > 0 {
		return true
	}

	migration := Migrations{IsMigrated: true}
	_, err = collection.InsertOne(context.TODO(), migration)
	if err != nil {
		fmt.Println("[checkMigrations][Error][saving migration]", err)
		return true
	}

	fmt.Println("[checkMigrations][migration saved]")
	return false
}

func InitDB() {
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
	MongoClient = client

	// Check migrations
	check := checkMigrations()
	if check {
		return
	}

	// Initialize Users
	users := []User{
		{Email: "user1@some.com", Password: "user1@some.com"},
		{Email: "user2@some.com", Password: "user2@some.com"},
	}
	collection := MongoClient.Database("app").Collection("users")
	for _, user := range users {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Println(err)
		}
	}

	// Initialize Schedule
	schedule := []Schedule{
		{Day: "Понедеьник", Start: "00:00", End: "24:00"},
		{Day: "Вторник", Start: "00:00", End: "24:00"},
		{Day: "Среда", Start: "00:00", End: "24:00"},
		{Day: "Четверг", Start: "00:00", End: "24:00"},
		{Day: "Пятница", Start: "00:00", End: "24:00"},
		{Day: "Суббота", Start: "00:00", End: "24:00"},
		{Day: "Воскресенье", Start: "00:00", End: "24:00"},
	}
	scheduleCollection := MongoClient.Database("app").Collection("schedule")
	for _, sch := range schedule {
		_, err := scheduleCollection.InsertOne(context.TODO(), sch)
		if err != nil {
			log.Println(err)
		}
	}
}
