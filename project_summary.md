# .air.toml

```toml
root = "."
tmp_dir = "build/tmp"

[build]
cmd = "make build"
bin = "./build/books-api-go"
full_bin = ""
include_ext = ["go", "tpl", "tmpl", "html", "env"]
exclude_dir = ["build", "tmp", "vendor", "docs/swagger"]
include_dir = []
exclude_file = []
delay = 500
poll = true
stop_on_error = true
log = "air.log"

[log]
time = false

[color]
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
clean_on_exit = true
```

# .gitignore

```
# Binary files
build/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with go test -c
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# IDE specific files
.idea/
.vscode/
*.swp
*.swo

# Environment files
.env
.env.local

# OS specific files
.DS_Store
Thumbs.db

# Swagger generated files
docs/swagger/*
!docs/swagger/.gitkeep

# Logs
*.log

# Air temporary files
build/tmp
air.log
```

# cmd/api/main.go

```go
package main

import (
	"log"

	"github.com/AhmadMuj/books-api-go/internal/config"
	"github.com/AhmadMuj/books-api-go/internal/handlers"
	"github.com/AhmadMuj/books-api-go/internal/repository"
	"github.com/AhmadMuj/books-api-go/internal/service"
	"github.com/gin-gonic/gin"
)

// @title Books API Go
// @version 1.0
// @description A RESTful API for managing books with Kafka event streaming and Redis caching
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize database
	db, err := repository.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize repository
	bookRepo := repository.NewBookRepository(db.DB)

	// Initialize service
	bookService := service.NewBookService(bookRepo)

	// Initialize handler
	bookHandler := handlers.NewBookHandler(bookService)

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	handlers.SetupRoutes(r, bookHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

```

# docs/REQUIREMENTS.md

```md
# **Task: REST API with Go (Gin) \- Books Management System**

## **Objective:**

Build and host a **REST API** using **Go (Gin framework)** to manage a **"Book" entity**, incorporating **Swagger documentation**, **Kafka for event streaming**, and **Redis for caching**. The API should be hosted so that **Swagger UI** is accessible via a URL. Provide the GitHub repository link along with the Swagger file.

---

## **API Requirements**

* **Gin Framework** for routing.  
* **Swagger Documentation** using `swaggo/swag`.  
* **Kafka Integration** to publish book-related events (e.g., book creation, update, deletion).  
* **Redis Caching** to optimize book retrieval performance.  
* **Database**: Use **PostgreSQL** (or an in-memory database like SQLite for simplicity).  
* **Hosting**: Deploy on **Heroku, AWS, or any cloud provider**.  
* **Best Practices**: Use validation, logging, and proper error handling.

---

## **Book Model**

| Field | Type | Description |
| ----- | ----- | ----- |
| ID | integer | Auto-incremented identifier |
| Title | string | Book title |
| Author | string | Author's name |
| Year | integer | Year of publication |

---

## **API Endpoints**

1. **GET /books** \- Retrieve a list of all books (cached using Redis).  
2. **GET /books/{id}** \- Retrieve a specific book by its ID (cached using Redis).  
3. **POST /books** \- Create a new book and publish an event to Kafka.  
4. **PUT /books/{id}** \- Update an existing book’s details and publish an update event to Kafka.  
5. **DELETE /books/{id}** \- Delete a book by its ID and publish a deletion event to Kafka.

---

## **Additional Enhancements**

### **Kafka Integration**

* Every **POST, PUT, DELETE** request should publish an event to a Kafka topic (`book_events`).  
* Consumers can later be added to handle these events (e.g., logging, analytics, notifications).

### **Redis Caching**

* Cache book data to reduce database queries.  
* **GET /books** and **GET /books/{id}** should first check Redis before querying the database.  
* Expire cache entries when a **POST, PUT, DELETE** operation is performed.

### **Pagination**

* Implement **pagination** for the `GET /books` endpoint using **limit and offset**.

### **Validation**

* Ensure the **title is not empty**, **author is not empty**, and **year is a valid number**.

### **Proper HTTP Status Codes**

* `200` for successful retrieval.  
* `201` for successful creation.  
* `400` for validation errors.  
* `404` if a book is not found.  
* `500` for internal server errors.

---

## **Setup and Deployment Steps**

### **1\. Project Initialization**

* Initialize a Go project.  
* Install **Gin, Swaggo, Kafka Go Client, Redis Client, GORM** for PostgreSQL.

### **2\. Implement API Endpoints**

* Set up controllers, services, and repository layers.  
* Use **GORM** for database operations.

### **3\. Add Kafka Producer & Consumer**

* **Producer**: Publish book events when changes occur.  
* **Consumer (Optional)**: Log received events.

### **4\. Integrate Redis Caching**

* Store frequently accessed books.  
* Implement cache invalidation on updates.

### **5\. Generate Swagger Documentation**

* Use `swag init` to generate Swagger docs.  
* Serve Swagger UI at `/swagger`.

### **6\. Deployment**

* Deploy API using **Heroku, AWS, or another cloud provider**.  
* Ensure the Swagger UI and API are publicly accessible.

### **7\. GitHub Repository**

* Push code to GitHub.  
* Include **README** with setup instructions.

---

## **Expected Deliverables**

1. **GitHub Repository Link**:

   * Provide a link to the GitHub repo containing the full project.  
2. **Swagger Documentation URL**:

   * Provide a public URL where Swagger UI is hosted.  
3. **Public API Endpoints**:

   * Ensure endpoints are accessible for testing.  
4. **Postman Collection / cURL Commands**:

   * Share API testing scripts.

---

## **Tools & Technologies**

* **Go** (Gin Framework)  
* **PostgreSQL** (Primary Database)  
* **Redis** (Caching)  
* **Kafka** (Event Streaming)  
* **Swagger** (API Documentation)  
* **Docker** (Optional for local development)
```

# docs/swagger/docs.go

```go
// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/books": {
            "get": {
                "description": "Get a paginated list of books",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "List all books",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Page size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ListBooksResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new book with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Create a new book",
                "parameters": [
                    {
                        "description": "Book details",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBookRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.BookResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "description": "Get a book's details by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Get a book by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BookResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a book's details by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Update a book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Book details",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BookResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a book by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Delete a book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BookResponse": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateBookRequest": {
            "type": "object",
            "required": [
                "author",
                "title",
                "year"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "integer",
                    "minimum": 1500
                }
            }
        },
        "dto.ListBooksResponse": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.BookResponse"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total_items": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateBookRequest": {
            "type": "object",
            "required": [
                "author",
                "title",
                "year"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "integer",
                    "minimum": 1500
                }
            }
        },
        "errors.AppError": {
            "type": "object",
            "properties": {
                "err": {},
                "message": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/errors.ErrorType"
                }
            }
        },
        "errors.ErrorType": {
            "type": "string",
            "enum": [
                "NOT_FOUND",
                "ALREADY_EXISTS",
                "VALIDATION_ERROR",
                "DATABASE_ERROR",
                "INTERNAL_ERROR"
            ],
            "x-enum-varnames": [
                "NotFound",
                "AlreadyExists",
                "ValidationErr",
                "DatabaseErr",
                "InternalErr"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Books API Go",
	Description:      "A RESTful API for managing books with Kafka event streaming and Redis caching",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

```

# docs/swagger/swagger.json

```json
{
    "swagger": "2.0",
    "info": {
        "description": "A RESTful API for managing books with Kafka event streaming and Redis caching",
        "title": "Books API Go",
        "contact": {},
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/api/v1",
    "paths": {
        "/books": {
            "get": {
                "description": "Get a paginated list of books",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "List all books",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Page size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ListBooksResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new book with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Create a new book",
                "parameters": [
                    {
                        "description": "Book details",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBookRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.BookResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "description": "Get a book's details by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Get a book by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BookResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a book's details by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Update a book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Book details",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BookResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a book by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Delete a book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BookResponse": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateBookRequest": {
            "type": "object",
            "required": [
                "author",
                "title",
                "year"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "integer",
                    "minimum": 1500
                }
            }
        },
        "dto.ListBooksResponse": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.BookResponse"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total_items": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateBookRequest": {
            "type": "object",
            "required": [
                "author",
                "title",
                "year"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "year": {
                    "type": "integer",
                    "minimum": 1500
                }
            }
        },
        "errors.AppError": {
            "type": "object",
            "properties": {
                "err": {},
                "message": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/errors.ErrorType"
                }
            }
        },
        "errors.ErrorType": {
            "type": "string",
            "enum": [
                "NOT_FOUND",
                "ALREADY_EXISTS",
                "VALIDATION_ERROR",
                "DATABASE_ERROR",
                "INTERNAL_ERROR"
            ],
            "x-enum-varnames": [
                "NotFound",
                "AlreadyExists",
                "ValidationErr",
                "DatabaseErr",
                "InternalErr"
            ]
        }
    }
}
```

# docs/swagger/swagger.yaml

```yaml
basePath: /api/v1
definitions:
  dto.BookResponse:
    properties:
      author:
        type: string
      created_at:
        type: string
      id:
        type: integer
      title:
        type: string
      updated_at:
        type: string
      year:
        type: integer
    type: object
  dto.CreateBookRequest:
    properties:
      author:
        type: string
      title:
        type: string
      year:
        minimum: 1500
        type: integer
    required:
    - author
    - title
    - year
    type: object
  dto.ListBooksResponse:
    properties:
      books:
        items:
          $ref: '#/definitions/dto.BookResponse'
        type: array
      page:
        type: integer
      page_size:
        type: integer
      total_items:
        type: integer
      total_pages:
        type: integer
    type: object
  dto.UpdateBookRequest:
    properties:
      author:
        type: string
      title:
        type: string
      year:
        minimum: 1500
        type: integer
    required:
    - author
    - title
    - year
    type: object
  errors.AppError:
    properties:
      err: {}
      message:
        type: string
      type:
        $ref: '#/definitions/errors.ErrorType'
    type: object
  errors.ErrorType:
    enum:
    - NOT_FOUND
    - ALREADY_EXISTS
    - VALIDATION_ERROR
    - DATABASE_ERROR
    - INTERNAL_ERROR
    type: string
    x-enum-varnames:
    - NotFound
    - AlreadyExists
    - ValidationErr
    - DatabaseErr
    - InternalErr
host: localhost:8080
info:
  contact: {}
  description: A RESTful API for managing books with Kafka event streaming and Redis
    caching
  title: Books API Go
  version: "1.0"
paths:
  /books:
    get:
      description: Get a paginated list of books
      parameters:
      - default: 1
        description: Page number
        in: query
        name: page
        type: integer
      - default: 10
        description: Page size
        in: query
        name: size
        type: integer
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            $ref: '#/definitions/dto.ListBooksResponse'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/errors.AppError'
      summary: List all books
      tags:
      - books
    post:
      consumes:
      - application/json
      description: Create a new book with the provided details
      parameters:
      - description: Book details
        in: body
        name: book
        required: true
        schema:
          $ref: '#/definitions/dto.CreateBookRequest'
      produces:
      - application/json
      responses:
        "201":
          description: Created
          schema:
            $ref: '#/definitions/dto.BookResponse'
        "400":
          description: Bad Request
          schema:
            $ref: '#/definitions/errors.AppError'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/errors.AppError'
      summary: Create a new book
      tags:
      - books
  /books/{id}:
    delete:
      description: Delete a book by its ID
      parameters:
      - description: Book ID
        in: path
        name: id
        required: true
        type: integer
      produces:
      - application/json
      responses:
        "204":
          description: No Content
        "404":
          description: Not Found
          schema:
            $ref: '#/definitions/errors.AppError'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/errors.AppError'
      summary: Delete a book
      tags:
      - books
    get:
      description: Get a book's details by its ID
      parameters:
      - description: Book ID
        in: path
        name: id
        required: true
        type: integer
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            $ref: '#/definitions/dto.BookResponse'
        "404":
          description: Not Found
          schema:
            $ref: '#/definitions/errors.AppError'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/errors.AppError'
      summary: Get a book by ID
      tags:
      - books
    put:
      consumes:
      - application/json
      description: Update a book's details by its ID
      parameters:
      - description: Book ID
        in: path
        name: id
        required: true
        type: integer
      - description: Book details
        in: body
        name: book
        required: true
        schema:
          $ref: '#/definitions/dto.UpdateBookRequest'
      produces:
      - application/json
      responses:
        "200":
          description: OK
          schema:
            $ref: '#/definitions/dto.BookResponse'
        "400":
          description: Bad Request
          schema:
            $ref: '#/definitions/errors.AppError'
        "404":
          description: Not Found
          schema:
            $ref: '#/definitions/errors.AppError'
        "500":
          description: Internal Server Error
          schema:
            $ref: '#/definitions/errors.AppError'
      summary: Update a book
      tags:
      - books
swagger: "2.0"

```

# go.mod

```mod
module github.com/AhmadMuj/books-api-go

go 1.22.3

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/joho/godotenv v1.5.1
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/swag v1.16.4
	gorm.io/driver/postgres v1.5.11
	gorm.io/gorm v1.25.12
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/bytedance/sonic v1.12.8 // indirect
	github.com/bytedance/sonic/loader v0.2.3 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.0.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.24.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.2 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.14.0 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	golang.org/x/tools v0.30.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

```

# internal/config/config.go

```go
package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Kafka    KafkaConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

func LoadConfig(envFile string) (*Config, error) {
	if envFile == "" {
		envFile = ".env"
	}

	if err := godotenv.Load(envFile); err != nil {
		return nil, fmt.Errorf("error loading env file: %w", err)
	}

	config := &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "books_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		Kafka: KafkaConfig{
			Brokers: strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ","),
			Topic:   getEnv("KAFKA_TOPIC", "book_events"),
		},
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

```

# internal/dto/book.go

```go
package dto

import (
	"fmt"
	"time"

	"github.com/AhmadMuj/books-api-go/internal/errors"
	"github.com/AhmadMuj/books-api-go/internal/models"
)

type CreateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required,min=1500,ltefield=CurrentYear"`
}

type UpdateBookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year" binding:"required,min=1500,ltefield=CurrentYear"`
}

type BookResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListBooksResponse struct {
	Books      []BookResponse `json:"books"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalItems int64          `json:"total_items"`
	TotalPages int            `json:"total_pages"`
}

// Conversion helpers
func ToBookResponse(book *models.Book) *BookResponse {
	return &BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Year:      book.Year,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}

func ToBookResponseList(books []models.Book) []BookResponse {
	responses := make([]BookResponse, len(books))
	for i, book := range books {
		responses[i] = *ToBookResponse(&book)
	}
	return responses
}

// Validation helpers
func (r *CreateBookRequest) Validate() error {
	currentYear := time.Now().Year()
	if r.Year < 1500 || r.Year > currentYear {
		return errors.NewValidationError(fmt.Sprintf(
			"book year must be between 1500 and %d",
			currentYear,
		))
	}
	return nil
}

func (r *UpdateBookRequest) Validate() error {
	currentYear := time.Now().Year()
	if r.Year < 1500 || r.Year > currentYear {
		return errors.NewValidationError(fmt.Sprintf(
			"book year must be between 1500 and %d",
			currentYear,
		))
	}
	return nil
}

```

# internal/errors/errors.go

```go
package errors

import "fmt"

type ErrorType string

const (
	NotFound      ErrorType = "NOT_FOUND"
	AlreadyExists ErrorType = "ALREADY_EXISTS"
	ValidationErr ErrorType = "VALIDATION_ERROR"
	DatabaseErr   ErrorType = "DATABASE_ERROR"
	InternalErr   ErrorType = "INTERNAL_ERROR"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// Constructor functions for common errors
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    NotFound,
		Message: message,
	}
}

func NewAlreadyExistsError(message string) *AppError {
	return &AppError{
		Type:    AlreadyExists,
		Message: message,
	}
}

func NewDatabaseError(err error) *AppError {
	return &AppError{
		Type:    DatabaseErr,
		Message: "database operation failed",
		Err:     err,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Type:    ValidationErr,
		Message: message,
	}
}

func NewInternalError(err error) *AppError {
	return &AppError{
		Type:    InternalErr,
		Message: "internal server error",
		Err:     err,
	}
}

```

# internal/handlers/book_handler.go

```go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/AhmadMuj/books-api-go/internal/dto"
	"github.com/AhmadMuj/books-api-go/internal/errors"
	"github.com/AhmadMuj/books-api-go/internal/models"
	"github.com/AhmadMuj/books-api-go/internal/service"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

// @Summary Create a new book
// @Description Create a new book with the provided details
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.CreateBookRequest true "Book details"
// @Success 201 {object} dto.BookResponse
// @Failure 400 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	book := &models.Book{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	if err := h.bookService.CreateBook(c.Request.Context(), book); err != nil {
		statusCode := http.StatusInternalServerError
		if _, ok := err.(*errors.AppError); ok {
			statusCode = http.StatusBadRequest
		}
		c.JSON(statusCode, err)
		return
	}

	c.JSON(http.StatusCreated, dto.ToBookResponse(book))
}

// @Summary Get a book by ID
// @Description Get a book's details by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.BookResponse
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError("invalid book ID"))
		return
	}

	book, err := h.bookService.GetBook(c.Request.Context(), uint(id))
	if err != nil {
		statusCode := http.StatusInternalServerError
		if appErr, ok := err.(*errors.AppError); ok {
			if appErr.Type == errors.NotFound {
				statusCode = http.StatusNotFound
			}
		}
		c.JSON(statusCode, err)
		return
	}

	c.JSON(http.StatusOK, dto.ToBookResponse(book))
}

// @Summary List all books
// @Description Get a paginated list of books
// @Tags books
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Success 200 {object} dto.ListBooksResponse
// @Failure 500 {object} errors.AppError
// @Router /books [get]
func (h *BookHandler) ListBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	books, total, err := h.bookService.ListBooks(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	totalPages := (int(total) + pageSize - 1) / pageSize

	response := dto.ListBooksResponse{
		Books:      dto.ToBookResponseList(books),
		Page:       page,
		PageSize:   pageSize,
		TotalItems: total,
		TotalPages: totalPages,
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Update a book
// @Description Update a book's details by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body dto.UpdateBookRequest true "Book details"
// @Success 200 {object} dto.BookResponse
// @Failure 400 {object} errors.AppError
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError("invalid book ID"))
		return
	}

	var req dto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	book := &models.Book{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}

	if err := h.bookService.UpdateBook(c.Request.Context(), uint(id), book); err != nil {
		statusCode := http.StatusInternalServerError
		if appErr, ok := err.(*errors.AppError); ok {
			switch appErr.Type {
			case errors.NotFound:
				statusCode = http.StatusNotFound
			case errors.ValidationErr:
				statusCode = http.StatusBadRequest
			}
		}
		c.JSON(statusCode, err)
		return
	}

	c.JSON(http.StatusOK, dto.ToBookResponse(book))
}

// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 204 "No Content"
// @Failure 404 {object} errors.AppError
// @Failure 500 {object} errors.AppError
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewValidationError("invalid book ID"))
		return
	}

	if err := h.bookService.DeleteBook(c.Request.Context(), uint(id)); err != nil {
		statusCode := http.StatusInternalServerError
		if appErr, ok := err.(*errors.AppError); ok {
			if appErr.Type == errors.NotFound {
				statusCode = http.StatusNotFound
			}
		}
		c.JSON(statusCode, err)
		return
	}

	c.Status(http.StatusNoContent)
}

```

# internal/handlers/routes.go

```go
package handlers

import (
	_ "github.com/AhmadMuj/books-api-go/docs/swagger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, bookHandler *BookHandler) {
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.POST("", bookHandler.CreateBook)
			books.GET("", bookHandler.ListBooks)
			books.GET("/:id", bookHandler.GetBook)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
		}
	}
}

```

# internal/models/book.go

```go
package models

import "time"

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" binding:"required" gorm:"not null"`
	Author    string    `json:"author" binding:"required" gorm:"not null"`
	Year      int       `json:"year" binding:"required" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

```

# internal/repository/book_repository_pg.go

```go
package repository

import (
	"context"

	"github.com/AhmadMuj/books-api-go/internal/errors"
	"github.com/AhmadMuj/books-api-go/internal/models"
	"gorm.io/gorm"
)

type BookRepositoryPG struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryPG{
		db: db,
	}
}

func (r *BookRepositoryPG) Create(ctx context.Context, book *models.Book) error {
	// Check if book with same title and author exists
	var exists bool
	err := r.db.WithContext(ctx).
		Model(&models.Book{}).
		Select("count(*) > 0").
		Where("title = ? AND author = ?", book.Title, book.Author).
		Find(&exists).
		Error
	if err != nil {
		return errors.NewDatabaseError(err)
	}
	if exists {
		return errors.NewAlreadyExistsError("book with same title and author already exists")
	}

	result := r.db.WithContext(ctx).Create(book)
	if result.Error != nil {
		return errors.NewDatabaseError(result.Error)
	}
	return nil
}

func (r *BookRepositoryPG) GetByID(ctx context.Context, id uint) (*models.Book, error) {
	var book models.Book
	result := r.db.WithContext(ctx).First(&book, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFoundError("book not found")
		}
		return nil, errors.NewDatabaseError(result.Error)
	}
	return &book, nil
}

func (r *BookRepositoryPG) List(ctx context.Context, limit, offset int) ([]models.Book, int64, error) {
	var books []models.Book
	var total int64

	// Get total count
	if err := r.db.WithContext(ctx).Model(&models.Book{}).Count(&total).Error; err != nil {
		return nil, 0, errors.NewDatabaseError(err)
	}

	result := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&books)

	if result.Error != nil {
		return nil, 0, errors.NewDatabaseError(result.Error)
	}
	return books, total, nil
}

func (r *BookRepositoryPG) Update(ctx context.Context, book *models.Book) error {
	result := r.db.WithContext(ctx).Save(book)
	if result.Error != nil {
		return errors.NewDatabaseError(result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.NewNotFoundError("book not found")
	}
	return nil
}

func (r *BookRepositoryPG) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Book{}, id)
	if result.Error != nil {
		return errors.NewDatabaseError(result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.NewNotFoundError("book not found")
	}
	return nil
}

```

# internal/repository/book_repository.go

```go
package repository

import (
	"context"

	"github.com/AhmadMuj/books-api-go/internal/models"
)

type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
	GetByID(ctx context.Context, id uint) (*models.Book, error)
	List(ctx context.Context, limit, offset int) ([]models.Book, int64, error)
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id uint) error
}

```

# internal/repository/database.go

```go
package repository

import (
	"fmt"

	"github.com/AhmadMuj/books-api-go/internal/config"
	"github.com/AhmadMuj/books-api-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&models.Book{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &Database{DB: db}, nil
}

```

# internal/service/book_service_impl.go

```go
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/AhmadMuj/books-api-go/internal/errors"
	"github.com/AhmadMuj/books-api-go/internal/models"
)

func (s *bookService) CreateBook(ctx context.Context, book *models.Book) error {
	if err := validateBook(book); err != nil {
		return err
	}
	return s.repo.Create(ctx, book)
}

func (s *bookService) GetBook(ctx context.Context, id uint) (*models.Book, error) {
	if id == 0 {
		return nil, errors.NewValidationError("invalid book ID")
	}
	return s.repo.GetByID(ctx, id)
}

func (s *bookService) ListBooks(ctx context.Context, page, pageSize int) ([]models.Book, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return s.repo.List(ctx, pageSize, offset)
}

func (s *bookService) UpdateBook(ctx context.Context, id uint, book *models.Book) error {
	if id == 0 {
		return errors.NewValidationError("invalid book ID")
	}

	if err := validateBook(book); err != nil {
		return err
	}

	existingBook, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingBook == nil {
		return errors.NewNotFoundError("book not found")
	}

	book.ID = id
	return s.repo.Update(ctx, book)
}

func (s *bookService) DeleteBook(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.NewValidationError("invalid book ID")
	}
	return s.repo.Delete(ctx, id)
}

func validateBook(book *models.Book) error {
	if book == nil {
		return errors.NewValidationError("book cannot be nil")
	}
	if book.Title == "" {
		return errors.NewValidationError("book title is required")
	}
	if book.Author == "" {
		return errors.NewValidationError("book author is required")
	}

	currentYear := time.Now().Year()
	if book.Year < 1500 || book.Year > currentYear {
		return errors.NewValidationError(fmt.Sprintf(
			"book year must be between 1500 and %d",
			currentYear,
		))
	}
	return nil
}

```

# internal/service/book_service.go

```go
package service

import (
	"context"

	"github.com/AhmadMuj/books-api-go/internal/models"
	"github.com/AhmadMuj/books-api-go/internal/repository"
)

type BookService interface {
	CreateBook(ctx context.Context, book *models.Book) error
	GetBook(ctx context.Context, id uint) (*models.Book, error)
	ListBooks(ctx context.Context, page, pageSize int) ([]models.Book, int64, error)
	UpdateBook(ctx context.Context, id uint, book *models.Book) error
	DeleteBook(ctx context.Context, id uint) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}

```

# Makefile

```
.PHONY: build run test clean swagger dev

BINARY_NAME=books-api-go
BUILD_DIR=build

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/api/main.go

run:
	./$(BUILD_DIR)/$(BINARY_NAME)

dev:
	air -c .air.toml --build.poll=true

test:
	go test -v ./...

clean:
	rm -rf $(BUILD_DIR)

swagger:
	swag init -g ./cmd/api/main.go -o ./docs/swagger
	
.DEFAULT_GOAL := build

```

# README.md

```md
# Books API Go

A RESTful API built with Go (Gin) for managing books, featuring Swagger documentation, Kafka event streaming, and Redis caching.

## Project Requirements

For detailed project requirements and specifications, see [REQUIREMENTS.md](docs/REQUIREMENTS.md)
```

