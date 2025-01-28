package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := MongoClient.Database("app").Collection("users")
	var result User
	err = collection.FindOne(context.TODO(), bson.M{"email": user.Email, "password": user.Password}).Decode(&result)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 240).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(tokenString))
}
