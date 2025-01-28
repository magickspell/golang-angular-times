package http

import (
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckAccess(w http.ResponseWriter, r *http.Request) {
	day := time.Now().Weekday().String()
	collection := MongoClient.Database("app").Collection("schedule")
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
