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


## Database

### Database Name
```
employees
```

### Table Users
```
DROP TABLE IF EXISTS "users";
create table "users" ( 
  "user_id" serial primary key not null,
  "username" varchar(20) NOT NULL,
  "email" varchar(255) NOT NULL,
  "password" varchar(20) NOT NULL,
  "role" varchar(20) DEFAULT user,
  "created_on" timestamp DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE "users" ALTER COLUMN "password" TYPE varchar(255);

insert into "users" values
	('1','admin','admin@gmail.com','sandiadmin','admin', '2022-08-03 15:00:00'),
	('2','1','user@gmail.com','sandiuser','user', '2022-08-03 15:00:00');
```

### Table Employees
```
DROP TABLE IF EXISTS "employees";
create table "employees" (
	"employee_id" serial primary key not null,
	"name" varchar(50) NOT NULL,
	"gender" varchar(10) NOT NULL,
	"birthday" date NOT NULL,
	"department" varchar(50) NOT null,
	"status" smallint NOT NULL DEFAULT '1'
);

insert into "employees" values
	(1,'Adrian','Male','1980-04-04','IT', 0),
	(2,'Bima','Male','1987-05-06','HR', 1),
	(3,'Cantika','Female','1989-03-04','Marketing', 1),
	(4,'Dadan','Male','1990-11-01','Product', 1),
	(5,'Ema','female','1992-08-09','Operational', 0);
```


