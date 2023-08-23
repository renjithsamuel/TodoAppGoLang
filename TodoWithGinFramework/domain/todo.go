package domain

import (
	"database/sql"
	"errors"
	"todoGin/models"
)

var DB *sql.DB

func GetTodos() ([]models.Todo, error) {
	var todos []models.Todo
	sqlStatement := "SELECT id, title, completed FROM todos"
	rows, err := DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func GetTodoByID(id int) (*models.Todo, error) {
	var todo models.Todo
	sqlStatement := "SELECT id, title, completed FROM todos WHERE id = $1"
	err := DB.QueryRow(sqlStatement, id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func CreateTodo(todo *models.Todo) error {
	sqlStatement := "INSERT INTO todos (title, completed) VALUES ($1, $2)"
	_, err := DB.Exec(sqlStatement, todo.Title, todo.Completed)
	return err
}

func UpdateTodo(todo *models.Todo) error {
	sqlStatement := "UPDATE todos SET title = $1, completed = $2 WHERE id = $3"
	result, err := DB.Exec(sqlStatement, todo.Title, todo.Completed, todo.ID)
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

func DeleteTodoByID(id int) error {
	sqlStatement := "DELETE FROM todos WHERE id = $1"
	_, err := DB.Exec(sqlStatement, id)
	return err
}
