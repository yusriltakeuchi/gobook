# Gobook
Gobook is a golang boilerplate with web framework Gin & Gorm ORM.

Features:
- [x] MVC Pattern
- [x] Login, Register, Logout & Profile System
- [x] Middleware System with JWT and UniqueKey
- [x] Paginations Request
- [x] Example Books CRUD
- [x] Support with MySQL
- [x] Validator System
- [x] Relation Has One
- [x] Env Configurations
- [x] Helper Command

## How To Add New Library
```go
go run main.go install <packagename>
```

Example command 
```go
go run main.go install github.com/google/uuid
```

## How To Add New Model
1. Use this following command to make a new model
```go
go run main.go make model <name>
```
2. Migrate your database using command
```go
	go run main.go migrate
```

## How To Add New Controller
1. Use this following command to make a new controller
```go
go run main.go make controller <name>
```

## Database Setup
We use .env file to keep the configurations, just open the .env file and edit your database credentials here
```go
DB_USERNAME=root
DB_PASSWORD=123456
DB_NAME=gobook
```

## Commands
Start server using
```go
	go run main.go start
```

Migrate Database
```go
	go run main.go migrate
```

Install Required Package
```go
	go run main.go install <packagename>
```

Make New Controller
```go
	go run main.go make controller <name>
```

Make New Model
```go
	go run main.go make model <name>
```

Show Help
```go
	go run main.go help
```

## Postman Collection
Check this collection if you want to see the example requests
[Gobook Collection](https://documenter.getpostman.com/view/3808786/SzYevFGV "Gobook Collection")