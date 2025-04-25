# Task Manager API
This is a tiny task manager RESTful API built with GoLang using Gin Web Framework. It allows you to Add Tasks, Remove Tasks, Display a list with the tasks and their id, name, status and mark tasks as done. 

---

## Features
Add Task - Create a new task
Delete Task - Delete a task from the list
List Tasks - Display a list with the tasks, both done and pending
Delete a Task - Delete a task by giving the task id(You can find the id of each task on the List!)

Error Handling - The code supports basic error handling on each function
Code Modularity - The code is fully modular by:
1) Having seperate functions which are all then connected in the setupRoutes function.
2) Having 2 different files, main.go and tasks.go 

## Project Structure

    /screenshots<br>
    go.mod<br>
    go.sum<br>
    main.go<br>
    tasks.go<br>
    README.md<br>
---
## How to Run
1) Install Dependencies:<br>
Make sure you have Go installed and the Gin Framework. If you want to install Gin, simply run on the terminal:<br>
```
go get -u github.com/gin-gonic/gin
```
2) Navigate through terminal on the project directory and run:<br>
```bash
go run main.go tasks.go
```
3) Server will start at http://localhost:8080

##   API Endpoints

| Method | Endpoint          | Description             |
|--------|-------------------|-------------------------|
| POST   | `/tasks`          | Adds a new task         |
| GET    | `/tasks`          | List all tasks          |
| POST   | `/tasks/:id/done` | Mark a task as done     |
| DELETE | `/tasks/:id`      | Delete a task based on the ID|
