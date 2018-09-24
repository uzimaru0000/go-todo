package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/todo/models"
)

func TasksGet(c *gin.Context) {
	tasks := models.GetAllTasks()
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func TaskPost(c *gin.Context) {
	title := c.PostForm("title")
	now := time.Now()

	task := models.NewTask(title, now)
	err := task.Save()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post success"})
}

func TaskPatch(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task := models.GetTaskById(uint(id))
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": c.Param("id") + " is not found."})
		return
	}

	title := c.PostForm("title")
	err := task.Update(title)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

func TaskDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task := models.GetTaskById(uint(id))
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": c.Param("id") + " is not found."})
		return
	}

	err := task.Delete()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
