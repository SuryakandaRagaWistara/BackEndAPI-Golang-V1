package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"project/database"
	"project/handlers"	

)

func main() {
	database.Init()

	fmt.Println("Database Terhubung....")

	r := mux.NewRouter()

	// Task
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")
	r.HandleFunc("/users/{id}/task", handlers.GetUserWithTasks).Methods("GET")
	r.HandleFunc("/users/{id}/task", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTaskStatus).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	// User
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	
	// Configurasi
	port := ":8080"
	fmt.Printf("Server berjalan di http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
