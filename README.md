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


3. Struktur Project
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

## 4. Cara Install dan Menjalankan

### Clone Repo

```
git clone https://github.com/SuryakandaRagaWistara/BackEndAPI-Golang-V1.git
cd project


# Install dependencies
go mod tidy

# Jalankan server
go run main.go```
