# Gobook
Gobook is a golang boilerplate with web framework Gin & Gorm ORM.

Features:
- [x] MVC Pattern
- [x] Login, Register & Logout System
- [x] Middleware System with JWT and UniqueKey
- [x] Paginations Request
- [x] Example Books CRUD
- [x] Support with MySQL
- [x] Validator System

## How To Add New Library
Open file in **config/packages.go** and insert your library here
```go
//Insert your package library here
func GetPackages() []string {
	return []string{
		"github.com/gin-gonic/gin",
		"github.com/ulule/deepcopier",
		"github.com/go-sql-driver/mysql",
		"github.com/jinzhu/gorm",
		"golang.org/x/crypto/bcrypt",
		"github.com/dgrijalva/jwt-go"}
}
```

Use this command to install library requred in project
```go
	go run main.go install
```

## How To Add New Model
1. Create Model in **app/models**
2. Open file **app/database/migrations.go**
3. Insert your model here
```go
//Insert your model here
func GetModel() [][]interface{} {
	return [][]interface{}{
		{&models.Books{}, "Books"},
		{&models.User{}, "User"}}
}
```
4. Migrate your database using command
```go
	go run main.go migrate
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
	go run main.go install
```

Show Help
```go
	go run main.go help
```

## Postman Collection
Check this collection if you want to see the example requests
[Gobook Collection](https://documenter.getpostman.com/view/3808786/SzYevFGV "Gobook Collection")