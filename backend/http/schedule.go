package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	db "backend/db"
)

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GetSchedule][start]")
	collection := db.MongoClient.Database("app").Collection("schedule")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("[GetSchedule][Error][cursor]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var schedules []Schedule
	if err = cursor.All(context.TODO(), &schedules); err != nil {
		fmt.Println("[GetSchedule][Error][cursor]")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(schedules)
	fmt.Println(schedules)
}

func UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := db.MongoClient.Database("app").Collection("schedule")
	_, err = collection.UpdateOne(context.TODO(), bson.M{"day": schedule.Day}, bson.M{"$set": bson.M{"start": schedule.Start, "end": schedule.End}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CheckSchedule(w http.ResponseWriter, r *http.Request) {
	day := time.Now().Weekday().String()
	collection := db.MongoClient.Database("app").Collection("schedule")
	var schedule Schedule
	err := collection.FindOne(context.TODO(), bson.M{"day": day}).Decode(&schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	now := time.Now().Format("15:04")
	if now >= schedule.Start && now <= schedule.End {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
