# Mini project Back End Golang
Repository for my mini project about GO REST API with Gin, GORM, JWT & PostgreSQL

## Mini project contain 2 domains
* Employees
* Users

## Component of this mini project
* CRUD
* Register
* Login
* JWT

## Design Pattern
* Hexagonal Architechture

## File structure
```
mini-project-back-end   
│
└───app
│   │app.go   
│   │employeeHandler.go 
│   │userHandler.go 
|
└───auth
|   │authService.go
|   
└───domain
|   |employee.go
|   |employeeRepositoryDB.go
|   |user.go
|   |userRepositoryDB.go
|
└───dto
|   |employeePagination.go
|
└───errs
|   |errors.go
|
└───helper
|   |pagination.go
|
└───input
|   |employeeInput.go
|   |userInput.go
| 
└───logger
|   |logger.go
|
└───service
|   |serviceEmployee.go
|   |serviceUser.go
|
└───setupDB
|   |setupDB.go
|
|env
|go.mod
|go.sum
|main.go
|README.md

```
