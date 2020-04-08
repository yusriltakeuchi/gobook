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

## Install Required Library
Use this command to install library requred in project
```go
	go run main.go install
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