# Go Loan Application API

A RESTful API for loan applications built with Go, Echo framework, PostgreSQL, and Redis.

## Features

- 🏦 Loan application processing
- ✅ Input validation with custom validators
- 🗄️ PostgreSQL database integration
- 🚀 Redis caching
- 📝 Comprehensive logging
- 🐳 Docker containerization
- 🔧 Clean architecture

## API Endpoints

### Apply for Loan
- **POST** `/api/v1/loans`
- **Content-Type**: `application/json`

#### Request Body
```json
{
  "fullName": "Somkanit Jitsanook",
  "monthlyIncome": 5000,
  "loanAmount": 10000,
  "loanPurpose": "home",
  "age": 25,
  "phoneNumber": "0851234567",
  "email": "demo@example.com"
}
```

#### Response
```json
{
  "id": "uuid-generated",
  "status": "approved",
  "message": "Congratulations! Your loan application has been approved.",
  "createdAt": "2025-01-18T10:30:00Z"
}
```

## Project Structure

```
├── app/
│   ├── internal/
│   │   ├── config.go          # Configuration management
│   │   ├── validation.go      # Custom validators
│   │   └── loan/
│   │       ├── loan.go        # Data structures
│   │       ├── handler.go     # HTTP handlers
│   │       └── service.go     # Business logic
│   └── main.go               # Application entry point
├── config/
│   └── config.yaml           # Configuration file
├── docker-compose.yml        # Docker services
├── Makefile                  # Build commands
├── example_request.json      # API test example
├── go.mod                    # Go modules
└── README.md
```

## Prerequisites

- Go 1.21+
- Docker & Docker Compose
- Make (optional)

## Quick Start

### 1. Clone the repository
```bash
git clone <repository-url>
cd go-miniproject
```

### 2. Start infrastructure services
```bash
docker-compose up -d
```

### 3. Install dependencies
```bash
go mod tidy
```

### 4. Run the application
```bash
# Using Make
make run

# Or directly with Go
go run app/main.go
```

### 5. Test the API
```bash
curl -X POST http://localhost:8080/api/v1/loans \
  -H "Content-Type: application/json" \
  -d @example_request.json
```

## Available Commands

```bash
# Build the application
make build

# Run the application
make run

# Run tests
make test

# Format code
make fmt

# Run code analysis
make vet

# Clean build artifacts
make clean

# Install dependencies
make deps

# Run all checks
make check

# Build for production
make build-prod
```

## Configuration

Edit `config/config.yaml`:

```yaml
server:
  port: "8080"
database:
  host: "localhost"
  port: 5432
  user: "loan_user"
  password: "loan_password"
  dbname: "loan_db"
redis:
  host: "localhost"
  port: 6379
```

## Docker Services

- **PostgreSQL**: Database on port 5432
- **Redis**: Cache on port 6379
- **Adminer**: Database admin UI on port 8081

### Access Adminer
- URL: http://localhost:8081
- System: PostgreSQL
- Server: postgres
- Username: loan_user
- Password: loan_password
- Database: loan_db

## Loan Approval Logic

- **Approved**: Monthly income ≥ $3,000 and loan amount ≤ 5x monthly income
- **Under Review**: Monthly income ≥ $2,000 but doesn't meet approval criteria
- **Rejected**: Monthly income < $2,000

## Development

### Project Dependencies
- **Echo v4**: Web framework
- **Validator v10**: Input validation
- **UUID**: Unique ID generation
- **Viper**: Configuration management

### Code Style
- Follow Go conventions
- Use `gofmt` for formatting
- Run `go vet` for static analysis
- Write tests for business logic

## Environment Variables

```bash
# Server configuration
SERVER_PORT=8080

# Database configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=loan_user
DB_PASSWORD=loan_password
DB_NAME=loan_db

# Redis configuration
REDIS_HOST=localhost
REDIS_PORT=6379
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## License

This project is licensed under the MIT License.