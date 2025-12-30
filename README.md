## 1. Teknologi yang Digunakan

- **Bahasa Pemrograman:** Go
- **Database:** Supabase (PostgreSQL)
- **Router / Framework HTTP:** Gorilla Mux
- **Tool / Library Tambahan:**
  - `github.com/gorilla/mux` → routing
  - `github.com/joho/godotenv` → load environment variables
  - `github.com/supabase-community/supabase-go` → client Supabase


## 2. Tujuan Project

Project ini dibuat untuk:

- Membuat API backend sederhana untuk manajemen user dan task.
- Menyediakan fitur CRUD (Create, Read, Update, Delete) untuk task.
- Menerapkan error handling standar (400, 404, 500).
- Menggunakan Go dan Supabase sebagai backend.

##3. Struktur Project
```
project/
├─ main.go
├─ .env
├─ handlers/
│  ├─ user.go
│  └─ task.go
├─ models/
│  ├─ user.go
│  └─ task.go
├─ database/
│  └─ db.go
├─ go.mod
└─ README.md
```
## 4. Cara Install dan Menjalankan

### Clone Repo

```
git clone https://github.com/SuryakandaRagaWistara/BackEndAPI-Golang-V1.git
cd BackEndAPI-Golang-V1


# Install dependencies
go mod tidy

# Jalankan server
go run main.go
```
## 5. Contoh Request dan Respone API 
```
BASE_URL="http://localhost:8080"
```

================== TASK ==================
```
echo "1. GET /tasks"
curl -s -X GET "$BASE_URL/tasks" | jq
echo -e "\n"

echo "2. GET /tasks/{id} (contoh id=1)"
curl -s -X GET "$BASE_URL/tasks/1" | jq
echo -e "\n"

echo "3. GET /users/{id}/task (contoh id=1)"
curl -s -X GET "$BASE_URL/users/1/task" | jq
echo -e "\n"

echo "4. POST /users/{id}/task (contoh id=1)"
curl -s -X POST "$BASE_URL/users/1/task" \
    -H "Content-Type: application/json" \
    -d '{"title":"Task Baru","description":"Deskripsi task","deadline":"2025-12-31"}' | jq
echo -e "\n"

echo "5. PATCH /tasks/{id} (contoh id=1)"
curl -s -X PATCH "$BASE_URL/tasks/1" \
    -H "Content-Type: application/json" \
    -d '{"status":"completed"}' | jq
echo -e "\n"

echo "6. DELETE /tasks/{id} (contoh id=1)"
curl -s -X DELETE "$BASE_URL/tasks/1" | jq
echo -e "\n"

echo "================== USER =================="

echo "7. GET /users"
curl -s -X GET "$BASE_URL/users" | jq
echo -e "\n"

echo "8. POST /users"
curl -s -X POST "$BASE_URL/users" \
    -H "Content-Type: application/json" \
    -d '{"name":"Budi","email":"budi@example.com"}' | jq
echo -e "\n"

echo "9. GET /users/{id} (contoh id=1)"
curl -s -X GET "$BASE_URL/users/1" | jq
echo -e "\n"

echo "10. DELETE /users/{id} (contoh id=1)"
curl -s -X DELETE "$BASE_URL/users/1" | jq
echo -e "\n"

echo "================== END =================="
```
