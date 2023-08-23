package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"todoGin/domain"
	"todoGin/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// initializing DB
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	connectionStr := os.Getenv("connectionString")
	var err error
	domain.DB, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := domain.DB.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	fmt.Println("Connection to DB successful!")

	// initializing router
	r := routes.SetupRouter()
	if err = r.Run(":8080"); err != nil {
		log.Fatal("Server Stopped Running!:", err)
	}
}