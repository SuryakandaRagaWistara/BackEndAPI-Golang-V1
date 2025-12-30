package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"project/database"
	"project/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	data, _,err := database.Client.
		From("User").
		Select("*","",false).
		Execute()
		
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := json.Unmarshal(data, &users); err != nil {
		http.Error(w, "Failed to parse users", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	
	if len(users) == 0 {
		http.Error(w, "users not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	data, _, err := database.Client.
		From("User").
		Insert(user, false, "", "representation", "").
		Execute()

	if err != nil {
		log.Println("DB ERROR:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	var createdUsers []models.User
	err = json.Unmarshal(data, &createdUsers)
	if err != nil {
		log.Println("Unmarshal error:", err)
		http.Error(w, "Failed to process response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if len(createdUsers) > 0 {
		json.NewEncoder(w).Encode(createdUsers[0])
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
    // 1. Ambil ID dari URL parameter
    idStr := mux.Vars(r)["id"]

    // 2. Validasi ID
    if _, err := strconv.Atoi(idStr); err != nil {
        http.Error(w, "Invalid user id", http.StatusBadRequest)
        return
    }

    var users []models.User
    data, _, err := database.Client.
        From("User").
        Select("*", "", false).
        Eq("id", idStr). 
        Execute()

    if err != nil {
        log.Println("Database error:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // 3. Unmarshal data
    if err := json.Unmarshal(data, &users); err != nil {
        http.Error(w, "Failed to parse user data", http.StatusInternalServerError)
        return
    }

    // 4. Cek apakah user ditemukan
    if len(users) == 0 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // 5. Kirim response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users[0])
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    // 1. Ambil ID dari URL
    idStr := mux.Vars(r)["id"]

    // 2. Validasi ID
    if _, err := strconv.Atoi(idStr); err != nil {
        http.Error(w, "Invalid user id", http.StatusBadRequest)
        return
    }

    resp, _, err := database.Client.
        From("User").
        Delete("", ""). 
        Eq("id", idStr).
        Execute()

    if err != nil {
        log.Println("Database error:", err)
        http.Error(w, "Failed to delete user", http.StatusInternalServerError)
        return
    }

    // 3. Cek apakah ada row yang dihapus
    var result []interface{}
    if err := json.Unmarshal(resp, &result); err != nil {
        http.Error(w, "Failed to parse delete response", http.StatusInternalServerError)
        return
    }

    if len(result) == 0 {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

	// Kirim pesan berhasil
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "User deleted successfully",
    })
}
