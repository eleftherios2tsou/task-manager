package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

var tasks []Task
var idCounter = 1

func addTask(taskName string) {
	//The function to add a task, it receives the task name as a parameter, gives an id,
	//  and set status false as default and increase the idCounter for the next task
	t := Task{
		ID:     idCounter,
		Name:   taskName,
		Status: false,
	}
	tasks = append(tasks, t)
	idCounter++

}
func delTask(id int) bool {
	//The function to delete a task, it uses the given id, it loops through each task
	// and every new tasks it does not match the given id, it appends into the newSlice

	newSlice := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID != id {
			newSlice = append(newSlice, t)

		} else {
			found = true
		}
	}
	tasks = newSlice
	return found
}
func listTasks() []Task {
	return tasks

}
func doneTask(id int) bool {
	// The function to mark a task as Done. It iterates through tasks to find
	// the id and when it find it it sets the status as true -> Done
	found := false
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks[i].Status = true
			found = true
			break
		}

	}
	return found

}
func setupRoutes(router *gin.Engine) {
	//This function defines all the HTTP routes for the API
	//It connect each route with the handler functions accordingly
	// - POST /tasks: Add a new task based on a JSON input
	// - GET /tasks: Returns a list of all tasks with their status (pending or done)
	// - POST /tasks/:id/done: Marks a task based on the id as done
	// - DELETE /tasks/:id: Delete a task based on the id.

	// Handler for adding new tasks
	router.POST("/tasks", func(c *gin.Context) {
		var data struct {
			Name string `json:"name"`
		}
		if err := c.BindJSON(&data); err != nil || data.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
			return

		}
		addTask(data.Name)
		c.JSON(http.StatusCreated, gin.H{"massage": "Task Added", "id": idCounter - 1})
	})
	//Handler for listing tasks
	router.GET("/tasks", func(c *gin.Context) {
		list := listTasks()
		var response []gin.H
		for _, t := range list {
			status := "pending"
			if t.Status {
				status = "Done"
			}
			response = append(response, gin.H{
				"id":     t.ID,
				"name":   t.Name,
				"status": status,
			})
		}
		c.IndentedJSON(http.StatusOK, response)
	})
	//Handler for marking a task as done
	router.POST("/tasks/:id/done", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task id"})
			return

		}
		found := doneTask(id)
		if found {
			c.JSON(http.StatusOK, gin.H{"message": "Task is done"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task is not found"})

		}
	})
	//Handler for deleting a task
	router.DELETE("/tasks/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You provided wrong task ID"})
			return
		}
		found := delTask(id)
		if found {
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		}

	})

}
