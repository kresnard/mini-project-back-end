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


## REST API

### Get All Employees
```
localhost:9090/employees
```

### Result Get All Employee
```
{
    "limit": 10,
    "page": 0,
    "total_rows": 0,
    "rows": [
        {
            "id": 1,
            "name": "Adrian",
            "gender": "Male",
            "birthday": "1980-04-04T00:00:00Z",
            "department": "IT",
            "status": 0
        },
        {
            "id": 2,
            "name": "Bima",
            "gender": "Male",
            "birthday": "1987-05-06T00:00:00Z",
            "department": "HR",
            "status": 1
        },
        {
            "id": 3,
            "name": "Cantika",
            "gender": "Female",
            "birthday": "1989-03-04T00:00:00Z",
            "department": "Marketing",
            "status": 1
        },
        {
            "id": 4,
            "name": "Dadan",
            "gender": "Male",
            "birthday": "1990-11-01T00:00:00Z",
            "department": "Product",
            "status": 1
        },
        {
            "id": 7,
            "name": "sdalasdah",
            "gender": "aale",
            "birthday": "1920-01-02T00:00:00Z",
            "department": "aksks",
            "status": 1
        },
        {
            "id": 8,
            "name": "Mardun",
            "gender": "Male",
            "birthday": "1920-01-02T00:00:00Z",
            "department": "Design",
            "status": 0
        }
    ]
}
```

### Get Employee by ID
```
localhost:9090/employees/1
```

### Result Get Employee by ID
```
{
    "id": 1,
    "name": "Adrian",
    "gender": "Male",
    "birthday": "1980-04-04T00:00:00Z",
    "department": "IT",
    "status": 0
}
```

### Create Employee (POST)
```
localhost:9090/employees
```

### Create Employee Body
```
{
    "name":"Mardun",
    "gender":"Male",
    "birthday":"1920-01-02",
    "department":"Design",
    "status":0
}
```

### Create Employee Result
```
{
    "id": 9,
    "name": "Mardun",
    "gender": "Male",
    "birthday": "1920-01-02",
    "department": "Design",
    "status": 0
}
```

### Delete Employee by ID
```
localhost:9090/employees/1
```

### Result Delete Employee by ID
```
{
    "id": 0,
    "name": "",
    "gender": "",
    "birthday": "",
    "department": "",
    "status": 0
}
```

### Update Employee (POST)
```
localhost:9090/employees/7
```

### Update Employee Body
```
{
    "name":"Madan",
    "gender":"Male",
    "birthday":"1920-01-02",
    "department":"HR",
    "status":0
}
```

### Update Employee Result
```
{
    "id": 0,
    "name": "Madan",
    "gender": "Male",
    "birthday": "1920-01-02",
    "department": "HR",
    "status": 0
}
```

### Registrasi User (POST)
```
localhost:9090/users
```

### Registrasi User Body
```
{
    "username":"kresna123",
    "password":"sandikresna123",
    "email":"kresna123@gmail.com",
    "role": "user"
}
```

### Registrasi User Result
```
{
    "user_id": 19,
    "username": "kresna123",
    "email": "kresna123@gmail.com",
    "password": "$2a$04$O//3McJFT7lGnwEvjxMRSeO8Fqff.C4bbWUs8/NjxDcruFV4L9aUC",
    "role": "user",
    "created_on": "2022-08-09T22:05:06.2359575+07:00"
}
```

### Login User (POST)
```
localhost:9090/login
```

### Login User Body
```
{
    "email":"kresna123@gmail.com",
    "password":"sandikresna123"
}
```

### Login User Result
```
{
    "user_id": 18,
    "email": "kresna123@gmail.com",
    "role": "user",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTgiLCJleHAiOjE2NjAxNDM1NzZ9.5Oye8WcAOo20LK5VkR9n_7f6KX6y-UY-7JfnY45jQbs"
}
```





