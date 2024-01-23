## Golang Fiber CRUD

### Overview

This project serves as a boilerplate for a CRUD (Create, Read, Update, Delete) application using [Fiber](https://github.com/gofiber/fiber) framework in Golang.

### Project Structure

```lua
gofiber-boilerplate
|-- src
|   |-- api
|       |-- configs
|           |-- config
|           |-- mongo
|           |-- ...
|       |-- controllers
|           |-- ...
|       |-- helpers
|           |-- error
|           |-- ...
|       |-- keys
|           |-- private.pem, public.pem (mocking)
|       |-- middleware
|           |-- ...
|       |-- models
|           |-- ...
|       |-- routes
|           |-- v0, v1, v2
|   |-- .env
|   |-- go.mod
|   |-- go.sum
|   |-- main.go
|-- .gitignore
|-- ...dockerfile
|-- readme.md
|-- test.http
```

### Installation

Install the Fiber framework:
```bash
go mod init golang-fiber-crud
```

Initialize the Go module:
```bash
go get github.com/gofiber/fiber/v2

# Load environment variables
go get github.com/joho/godotenv
```

Initialize the Mongodb module:
```bash
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```

Install JWT package for authentication
```bash
go get github.com/form3tech-oss/jwt-go
```

### Run the Application

Run the application using the following command:
```bash
go run .
# or
go run main.go
```
