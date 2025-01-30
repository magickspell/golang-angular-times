package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	db "backend/db"
)

type CheckRequest struct {
	Day    string `json:"day"`
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

type Token struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	collection := db.MongoClient.Database("app").Collection("schedule")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "[GetSchedule][Error][cursor]", http.StatusInternalServerError)
		return
	}
	var schedules []Schedule
	if err = cursor.All(context.TODO(), &schedules); err != nil {
		http.Error(w, "[GetSchedule][Error][schedules]", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(schedules)
}

func getEmailFromToken(tokenString string, secretKey []byte) (*string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		return &claims.Email, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func isValidToken(token string) bool {
	email, err := getEmailFromToken(strings.TrimPrefix(token, "Bearer "), JwtSecret)
	if err != nil {
		fmt.Println("[isValidToken][invalid token]")
		return false
	}

	collection := db.MongoClient.Database("app").Collection("users")
	filter := bson.M{"email": *email}
	var result User
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("[isValidToken][no user]")
			return false
		}
		log.Println(err)
		return false
	}

	return true
}

func UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		http.Error(w, "[UpdateSchedule][Error][no-token]", http.StatusUnauthorized)
		return
	}

	if !isValidToken(authToken) {
		http.Error(w, "[UpdateSchedule][Error][invalid-token]", http.StatusUnauthorized)
		return
	}

	var schedules []Schedule
	err := json.NewDecoder(r.Body).Decode(&schedules)
	if err != nil {
		http.Error(w, "[UpdateSchedule][Error][schedules]", http.StatusBadRequest)
		return
	}

	collection := db.MongoClient.Database("app").Collection("schedule")
	for _, schedule := range schedules {
		if !strings.Contains(schedule.Start, ":") || !strings.Contains(schedule.End, ":") {
			http.Error(w, "[UpdateSchedule][Error][schedule][wrong-format]", http.StatusBadRequest)
			return
		}
		_, err = collection.UpdateOne(context.TODO(), bson.M{"day": schedule.Day}, bson.M{"$set": bson.M{"start": schedule.Start, "end": schedule.End}})
		if err != nil {
			http.Error(w, "[UpdateSchedule][Error][schedule]", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

func checkTimeFormat(time string) string {
	if time == "24:00" {
		return "23:59"
	} else {
		return time
	}
}

func CheckSchedule(w http.ResponseWriter, r *http.Request) {
	var reqBody CheckRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "[CheckSchedule][Error][reqBody]", http.StatusBadRequest)
		return
	}

	collection := db.MongoClient.Database("app").Collection("schedule")
	var schedule Schedule
	err = collection.FindOne(context.TODO(), bson.M{"day": reqBody.Day}).Decode(&schedule)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "[CheckSchedule][Error][schedule]", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	requestTimeStr := checkTimeFormat(reqBody.Hour + ":" + reqBody.Minute)
	requestTime, err := time.Parse("15:04", requestTimeStr)
	if err != nil {
		http.Error(w, "[CheckSchedule][Error][requestTime]", http.StatusBadRequest)
		return
	}

	schedulerStartTime := checkTimeFormat(schedule.Start)
	startTime, err := time.Parse("15:04", schedulerStartTime)
	if err != nil {
		http.Error(w, "[CheckSchedule][Error][startTime]", http.StatusInternalServerError)
		return
	}

	schedulerEndTime := checkTimeFormat(schedule.End)
	endTime, err := time.Parse("15:04", schedulerEndTime)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "[CheckSchedule][Error][endTime]", http.StatusInternalServerError)
		return
	}

	if !requestTime.Before(startTime) && (!requestTime.After(endTime) || requestTime.Equal(endTime)) {
		json.NewEncoder(w).Encode(map[string]bool{"allowed": true})
	} else {
		json.NewEncoder(w).Encode(map[string]bool{"allowed": false})
	}
}
