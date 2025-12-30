package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

// Tanda *Berfungsi sebagai pointer
// - pointer adalah variable yang menyimpan 
//   alamat memori dari variable lain

func Init(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")

	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		log.Fatal("Failed connect", err)
	}

	Client = client
}