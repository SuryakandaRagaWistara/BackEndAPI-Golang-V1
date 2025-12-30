1. Teknologi yang Digunakan

Bahasa Pemrograman: Go

Database: Supabase (PostgreSQL)

Router / Framework HTTP: Gorilla Mux

Tool / Library Tambahan:

github.com/gorilla/mux → routing

github.com/joho/godotenv → load environment variables

github.com/supabase-community/supabase-go → client Supabase


2. Tujuan Project

Menjelaskan tujuan dibuatnya project, misalnya:

Membuat API backend sederhana untuk manajemen user dan task, termasuk fitur CRUD (Create, Read, Update, Delete) dengan error handling standar (400, 404, 500) menggunakan Go dan Supabase.



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

# Clone repo
git clone <REPO_URL>
cd project

# Install dependencies
go mod tidy

# Jalankan server
go run main.go


