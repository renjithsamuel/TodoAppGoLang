package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"todoGin/domain"
	"todoGin/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetTodos(c *gin.Context) {
	todos, err := domain.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := domain.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	
	// Using the validator to validate the Todo
	if err := validate.Struct(todo); err != nil {
		var ErrorString string
		for _, err := range err.(validator.ValidationErrors) {
			ErrorString = ErrorString + fmt.Sprintf("Field: %s -> Error: %s , ", err.Field(), err.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"Validation Errors": ErrorString})
		return
	}

	err := domain.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Using the validator to validate the Todo
	if err := validate.Struct(todo); err != nil {
		var ErrorString string
		for _, err := range err.(validator.ValidationErrors) {
			ErrorString = ErrorString + fmt.Sprintf("Field: %s -> Error: %s , ", err.Field(), err.Tag())
		}
		c.JSON(http.StatusBadRequest, gin.H{"Validation Errors": ErrorString})
		return
	}

	todo.ID = uint32(id)
	err = domain.UpdateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = domain.DeleteTodoByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data deleted Successfully!"})
}
