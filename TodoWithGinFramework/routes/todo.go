package routes

import (
	"fmt"
	"net/http"
	"todoGin/handlers"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// creating gin router
func SetupRouter() *gin.Engine {
	r = gin.Default()
	r.Use(corsMiddleware())
	r.GET("/todos", handlers.GetTodos)
	r.GET("/todos/:id", handlers.GetTodo)
	r.POST("/todos", handlers.CreateTodo)
	r.PUT("/todos/:id", handlers.UpdateTodo)
	r.DELETE("/todos/:id", handlers.DeleteTodo)
	return r
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("%s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

