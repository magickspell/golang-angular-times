package http

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongoClient *mongo.Client
	JwtSecret   = []byte("your-jwt-secret")
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
