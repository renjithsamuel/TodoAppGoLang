package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"todoTest/domain"
	"todoTest/handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// DB initialization
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	connectionStr := os.Getenv("connectionString")
	var err error
	domain.DB, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	// get port from OS
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	defer domain.DB.Close()
	// routes
	http.HandleFunc("/todos/get", ValidateRoute(handlers.GetHandler(domain.DB), http.MethodGet))
	http.HandleFunc("/todos/create", ValidateRoute(handlers.CreateHandler(domain.DB), http.MethodPost))
	http.HandleFunc("/todos/update", ValidateRoute(handlers.UpdateHandler(domain.DB), http.MethodPut))
	http.HandleFunc("/todos/delete", ValidateRoute(handlers.DeleteHandler(domain.DB), http.MethodDelete))

	fmt.Printf("Server running in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ValidateRoute(next http.HandlerFunc, allowedMethod string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}
