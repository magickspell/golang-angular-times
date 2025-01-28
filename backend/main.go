package main

import (
	"fmt"
	"net/http"

	db "backend/db"
	handlers "backend/http"
)

func main() {
	fmt.Println("Server started on port: 8080")
	db.InitDB()

	http.HandleFunc("/go", handlers.HelloGo)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/schedule", handlers.GetSchedule)
	http.HandleFunc("/update-schedule", handlers.UpdateSchedule)
	http.HandleFunc("/check", handlers.CheckAccess)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: 8080")
	}
}
