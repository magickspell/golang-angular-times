package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	db "backend/db"
)

type CheckRequest struct {
	Day    string `json:"day"`
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

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
	// todo сделать проверки
	var schedules []Schedule
	err := json.NewDecoder(r.Body).Decode(&schedules)
	fmt.Println("[schedule]", schedules)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := db.MongoClient.Database("app").Collection("schedule")
	for _, schedule := range schedules {
		_, err = collection.UpdateOne(context.TODO(), bson.M{"day": schedule.Day}, bson.M{"$set": bson.M{"start": schedule.Start, "end": schedule.End}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := db.MongoClient.Database("app").Collection("schedule")
	var schedule Schedule
	err = collection.FindOne(context.TODO(), bson.M{"day": reqBody.Day}).Decode(&schedule)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Schedule not found for the specified day", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	requestTimeStr := checkTimeFormat(reqBody.Hour + ":" + reqBody.Minute)
	requestTime, err := time.Parse("15:04", requestTimeStr)
	if err != nil {
		http.Error(w, "Invalid time format", http.StatusBadRequest)
		return
	}

	schedulerStartTime := checkTimeFormat(schedule.Start)
	startTime, err := time.Parse("15:04", schedulerStartTime)
	if err != nil {
		http.Error(w, "Invalid start time format in database", http.StatusInternalServerError)
		return
	}

	schedulerEndTime := checkTimeFormat(schedule.End)
	endTime, err := time.Parse("15:04", schedulerEndTime)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid end time format in database", http.StatusInternalServerError)
		return
	}

	if !requestTime.Before(startTime) && (!requestTime.After(endTime) || requestTime.Equal(endTime)) {
		json.NewEncoder(w).Encode(map[string]bool{"allowed": true})
	} else {
		json.NewEncoder(w).Encode(map[string]bool{"allowed": false})
	}
}
