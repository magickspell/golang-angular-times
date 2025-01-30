package main

import (
	"fmt"
	"log"
	"net/http"

	db "backend/db"
	handlers "backend/http"

	"github.com/joho/godotenv"
)

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func main() {
	fmt.Println("Server started on port: 8080")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error laoding .env")
	}

	db.InitDB()

	http.HandleFunc("/go", cors(handlers.HelloGo))
	http.HandleFunc("/login", cors(handlers.Login))
	http.HandleFunc("/schedules", cors(handlers.GetSchedule))
	http.HandleFunc("/update-schedule", cors(handlers.UpdateSchedule))
	http.HandleFunc("/check", cors(handlers.CheckSchedule))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error on server")
	}
}
