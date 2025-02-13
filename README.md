# Books API Go

A RESTful API built with Go (Gin) for managing books, featuring Swagger documentation, Kafka event streaming, and Redis caching. This project was originally requested as part of a job application coding task, which I've kept as a showcase of my Go development skills and understanding of modern backend technologies.

## Project Requirements

For detailed project requirements and specifications that were provided as part of the original task, see [REQUIREMENTS.md](docs/REQUIREMENTS.md)

## Technologies Used

- Go (Gin Framework)
- PostgreSQL for persistent storage
- Redis for caching
- Kafka for event streaming
- Swagger for API documentation
- Docker & Docker Compose for containerization
- GORM as ORM
- Air for live reload during development

## Features

- RESTful API endpoints for CRUD operations on books
- Event streaming via Kafka for book events (create/update/delete)
- Redis caching to optimize read performance
- Database persistence with PostgreSQL
- API documentation with Swagger UI
- Docker support for both development and production
- Request ID tracking and logging
- CORS support
- Error handling and validation
- Resource limits for production containers

## Quick Start

### Prerequisites

- Docker and Docker Compose installed
- Go 1.22 or later (for local development)
- Make (optional, but recommended)

### Development Setup

1. Clone the repository:

```bash
git clone https://github.com/username/books-api-go
cd books-api-go
```

2. Start the development environment:

```bash
make docker-dev
```

This will start:

- API server on port 8080
- PostgreSQL on port 5432
- Redis on port 6379
- Kafka broker on port 9092

### Production Setup

1. Set required environment variables:

```bash
export DB_PASSWORD=your_db_password
export REDIS_PASSWORD=your_redis_password
```

2. Start the production environment:

```bash
make docker-prod
```

## API Endpoints

- `GET /api/v1/books` - List all books (paginated)
- `GET /api/v1/books/{id}` - Get a specific book
- `POST /api/v1/books` - Create a new book
- `PUT /api/v1/books/{id}` - Update a book
- `DELETE /api/v1/books/{id}` - Delete a book

Swagger documentation is available at `/swagger`

## Development

### Available Make Commands

- `make build` - Build the application
- `make run` - Run the built binary
- `make dev` - Run with live reload using Air
- `make test` - Run tests
- `make swagger` - Generate Swagger documentation
- `make docker-dev` - Start development environment
- `make docker-prod` - Start production environment
- `make docker-down` - Stop all containers

## License

This project is licensed under the MIT License.
