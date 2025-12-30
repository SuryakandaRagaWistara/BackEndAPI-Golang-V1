package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"project/database"
	"project/models"
	"strconv"
	"strings"
	// "fmt"

)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	data, _,err := database.Client.
		From("Task").
		Select("*","",false).
		Execute()

	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		log.Println(err)
		http.Error(w, "Failed to parse tasks", http.StatusInternalServerError)
		return
	}

	if len(tasks) == 0 {
		http.Error(w, "tasks not found", http.StatusNotFound)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// 1. Ambil ID dari URL
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Validasi ID harus angka
	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
		return
	}

	var task models.Task

	// 2. Query database
	data, _, err := database.Client.
		From("Task").
		Select("*", "", false).
		Eq("id", id). 
		Single().
		Execute()

	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		log.Println(err)
		return
	}

	// 3. Parsing data
	if err := json.Unmarshal(data, &task); err != nil {
		http.Error(w, "Error parsing data", http.StatusInternalServerError)
		return
	}

	// 4. Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}


func GetUserWithTasks(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    userID := vars["id"]
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	var user models.User

	data, _, err := database.Client.
		From("User").
		// Menggunakan alias: tasks:task(*) agar JSON yang kembali punya key "tasks"
		Select("*, tasks:Task(*)", "", false). 
		Eq("id", userID).
		Single().
		Execute()
	
	if err != nil {
		http.Error(w, "User not found or databse error", http.StatusNotFound)
		log.Println(err)
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, "Failed to parsedata", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]

	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "ID user tidak valid", http.StatusBadRequest)
		return
	}

	// CEK USER (FIX UTAMA)
	_, _, err = database.Client.
		From("User").
		Select("id", "", false).
		Eq("id", idParam).
		Single().
		Execute()

	if err != nil {
		http.Error(w, "User tidak ditemukan", http.StatusNotFound)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Format data tidak valid", http.StatusBadRequest)
		return
	}

	task.UserID = userID

	if task.Status == "" {
		task.Status = models.StatusStarted
	}

	data, _, err := database.Client.
		From("Task").
		Insert(task, false, "", "", "").
		Execute()

	if err != nil {
		http.Error(w, "Gagal menyimpan task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	if taskID == "" {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	var input struct {
		Status models.TaskStatus `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Format data tidak valid", http.StatusBadRequest)
		return
	}
	// NORMALISASI 
	input.Status = models.TaskStatus(
		strings.ToLower(string(input.Status)),
	)

	// validasi status di API layer
	switch input.Status {
	case models.StatusStarted,
		models.StatusInProgress,
		models.StatusCompleted,
		models.StatusCanceled:
	default:
		http.Error(w, "Status tidak valid", http.StatusBadRequest)
		return
	}

	// cek dulu apakah task ada
	check, _, err := database.Client.
		From("Task").
		Select("id", "", false).
		Eq("id", taskID).
		Execute()

	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if len(check) == 0 {
		http.Error(w, "Task tidak ditemukan", http.StatusNotFound)
		return
	}

	// baru update
	data, _, err := database.Client.
		From("Task").
		Update(map[string]any{
			"status": input.Status,
		}, "", "").
		Eq("id", taskID).
		Execute()

	if err != nil {
		http.Error(w, "Gagal memperbarui status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}


func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	if taskID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if _, err := strconv.Atoi(taskID); err != nil {
		http.Error(w, "id must be number", http.StatusBadRequest)
		return
	}

	data, _, err := database.Client.
		From("Task").
		Delete("", "").
		Eq("id", taskID).
		Execute()

	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if string(data) == "[]" {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"task deleted successfully"}`))
}


