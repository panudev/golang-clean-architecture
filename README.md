# Go Clean Architecture API

A RESTful API built with Go using Clean Architecture principles, featuring Gin web framework, structured logging, Swagger documentation, configuration management, GORM ORM, and MySQL database.

## 🚀 Features

- **Clean Architecture**: Separation of concerns with distinct layers (entities, use cases, interfaces, and frameworks)
- **Gin Web Framework**: Fast and efficient HTTP routing and middleware
- **Structured Logging**: Comprehensive logging system
- **Swagger Documentation**: API documentation with Swagger/OpenAPI
- **Configuration Management**: Environment-based configuration
- **GORM**: Powerful ORM for database operations
- **MySQL**: Relational database
- **Docker**: Containerization support

## 📋 Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- MySQL 8.0
- Make (optional, for using Makefile commands)

## 🛠️ Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-clean-architecture.git
cd go-clean-architecture
```

2. Copy the environment file:
```bash
cp .env.example .env
```

3. Update the environment variables in `.env` file according to your needs.

## 📁 Project Structure

```
go-clean-app/
├── cmd/                        # Application entry points
│   └── server/
│       └── main.go            # Main application entry point
├── internal/                   # Private application code
│   ├── domain/                # Enterprise business rules
│   │   ├── entity/           # Business objects
│   │   └── repository/       # Repository interfaces
│   ├── usecase/              # Application business rules
│   │   └── user/            # User-related use cases
│   ├── interface/            # Interface adapters
│   │   ├── handler/         # HTTP handlers
│   │   └── route/           # API routes
│   ├── infrastructure/       # Frameworks and drivers
│   │   └── mysql/           # MySQL implementation
│   └── bootstrap/           # Dependency injection
├── pkg/                      # Public library code
│   ├── logger/              # Logging utilities
│   └── config/              # Configuration management
├── api/                      # API documentation
│   └── docs/                # Swagger/OpenAPI specs
├── .env                      # Environment variables
├── Dockerfile               # Docker configuration
├── docker-compose.yml       # Docker services
├── go.mod                   # Go module file
└── README.md               # Project documentation
```

## 🧪 Testing

Run the test suite:
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Run specific test
go test ./internal/usecase/...
```

## 🔧 Configuration

The application uses environment variables for configuration. Key configurations include:

- `APP_PORT`: Application port (default: 8080)
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name

## 📦 Dependencies

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Swagger](https://github.com/swaggo/swag) - API documentation
- [Zap](https://github.com/uber-go/zap) - Logging
- [dotenv](https://github.com/joho/godotenv) - Environment variables
- [Wire](https://github.com/google/wire) - Dependency injection

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Clean Architecture by Robert C. Martin
- Go best practices and community guidelines

### Layer Descriptions

#### 1. Domain Layer (`internal/domain/`)
- Contains enterprise business rules
- Independent of other layers
- Includes:
  - Entities (business objects)
  - Repository interfaces
  - Value objects
  - Domain events

#### 2. Use Case Layer (`internal/usecase/`)
- Implements application business rules
- Orchestrates the flow of data
- Depends only on domain layer
- Contains:
  - Business logic
  - Input/output ports
  - Use case implementations

#### 3. Interface Layer (`internal/interface/`)
- Adapts external interfaces to use cases
- Includes:
  - HTTP handlers
  - API routes
  - Request/Response DTOs
  - Input validation

#### 4. Infrastructure Layer (`internal/infrastructure/`)
- Implements interfaces defined in domain layer
- Contains:
  - Database implementations
  - External service clients
  - Framework-specific code

## 🏃‍♂️ Running the Application

### Using Docker

```bash
# Build and run the application
docker-compose up --build

# Run in detached mode
docker-compose up -d
```

### Local Development

1. Install dependencies:
```bash
go mod download
```

2. Run the application:
```bash
go run cmd/server/main.go
```

## 📚 API Documentation

Once the application is running, access the Swagger documentation at:
```
http://localhost:8080/swagger/index.html