package http

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	collection := MongoClient.Database("app").Collection("schedule")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var schedules []Schedule
	if err = cursor.All(context.TODO(), &schedules); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(schedules)
}

func UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule Schedule
	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := MongoClient.Database("app").Collection("schedule")
	_, err = collection.UpdateOne(context.TODO(), bson.M{"day": schedule.Day}, bson.M{"$set": bson.M{"start": schedule.Start, "end": schedule.End}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
