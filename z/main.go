// Author: Eleftherios Angelos Tsourdiou
// Simple CLI Task Manager in Go

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID     int
	Title  string
	IsDone bool
}

var tasks []Task
var nextID int = 1

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: add | list | done | delete")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		addTask()
	case "list":
		listTasks()
	case "done":
		markTaskDone()
	case "delete":
		deleteTask()
	default:
		fmt.Println("Unknown command. Use: add | list | done | delete")
	}
}

func addTask() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a task title.")
		return
	}

	title := os.Args[2]
	newTask := Task{
		ID:     nextID,
		Title:  title,
		IsDone: false,
	}
	tasks = append(tasks, newTask)
	nextID++
	fmt.Println("Added task:", title)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	for _, task := range tasks {
		status := "Pending"
		if task.IsDone {
			status = "Done"
		}
		fmt.Printf("ID: %d | Title: %s | Status: %s\n", task.ID, task.Title, status)
	}
}

func markTaskDone() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the task ID to mark as done.")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].IsDone = true
			fmt.Println("Marked task as done:", task.Title)
			return
		}
	}
	fmt.Println("Task not found.")
}

func deleteTask() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide the task ID to delete.")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Deleted task:", task.Title)
			return
		}
	}
	fmt.Println("Task not found.")
}
