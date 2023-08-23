package domain

import (
	"database/sql"
	"errors"
	"todoTest/models"
)

var DB *sql.DB


func CreateTodo(db *sql.DB, todo *models.Todo) error {
	sqlStatement := "INSERT INTO todos (title, completed) VALUES ($1, $2)"
	_, err := db.Exec(sqlStatement, todo.Title, todo.Completed)
	return err
}

func UpdateTodo(db *sql.DB, todo *models.Todo) error {
	sqlStatement := "UPDATE todos SET title = $1, completed = $2 WHERE id = $3"
	result, err := db.Exec(sqlStatement, todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("No Changes Made")
	}

	return nil
}

func DeleteTodo(db *sql.DB, id uint32) error {
	sqlStatement := "DELETE FROM todos WHERE id=$1"
	_, err := db.Exec(sqlStatement, id)
	return err
}

func GetTodos(db *sql.DB) ([]models.Todo, error) {
	// fetching rows
	sqlStatement := "SELECT id, title, completed FROM todos"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// binding rows with struct array
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, err
}
