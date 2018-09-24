package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/todo/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/", controllers.IndexGet)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/tasks", controllers.TasksGet)
		v1.POST("/tasks", controllers.TaskPost)
		v1.PATCH("/tasks/:id", controllers.TaskPatch)
		v1.DELETE("/tasks/:id", controllers.TaskDelete)
	}

	router.Run(":5000")
}
