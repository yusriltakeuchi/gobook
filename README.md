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
- [x] ORM Features
- [x] Auto Generated Code
- [x] Support MySQL & PostgreSQL

## Installing
1. Clone the project inside src folder
```go
git clone https://github.com/yusriltakeuchi/gobook.git
```

2. Make sure your GOPATH and GOROOT already defined.
This is an example of my path:
```go
export GOROOT=/usr/local/go
export GOPATH=/Volumes/YURANI/Software\ Development/Golang
export PATH=$PATH:$GOROOT/bin:$GOPATH
```
GOROOT is the place where you install the golang software,
and GOPATH is the place your golang project stored

2. Install required library
```go
go run main.go install
```

3. Setup your .env Database Credentials
4. Migrate your model using command
```go
go run main.go migrate
```

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

Using MySQL
```go
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=password
DB_NAME=gobook
```

Using PostgreSQL
```go
DB_CONNECTION=postgres
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=password
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