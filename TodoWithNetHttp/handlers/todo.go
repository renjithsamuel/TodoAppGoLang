package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"todoTest/domain"
	"todoTest/models"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := domain.GetTodos(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(todos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo models.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "Error : Invalid Data provided!", http.StatusBadRequest)
			return
		}

		// Using the validator to validate the Todo
		if err := validate.Struct(todo); err != nil {
			var ErrorString string
			for _, err := range err.(validator.ValidationErrors) {
				ErrorString = ErrorString + fmt.Sprintf("Field: %s -> Error: %s , ", err.Field(), err.Tag())
			}
			http.Error(w, "Validation Errors : " + ErrorString , http.StatusBadRequest)
			return
		}

		err = domain.CreateTodo(db, &todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var todo models.Todo
		todo.ID = uint32(id)
		err = json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "Error : Invalid Data provided!", http.StatusBadRequest)
			return
		}

		// Using the validator to validate the Todo
		if err := validate.Struct(todo); err != nil {
			var ErrorString string
			for _, err := range err.(validator.ValidationErrors) {
				ErrorString = ErrorString + fmt.Sprintf("Field: %s -> Error: %s , ", err.Field(), err.Tag())
			}
			http.Error(w, "Validation Errors : " + ErrorString , http.StatusBadRequest)
			return
		}

		err = domain.UpdateTodo(db, &todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func DeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = domain.DeleteTodo(db, uint32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(`{status : success}`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
