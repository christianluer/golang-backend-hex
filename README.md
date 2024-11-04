# golang-backend-hex

structure of the project
```
go-auth-app/
├── cmd/
│   └── main.go                # Entry point
├── domain/
│   └── user.go                # Core business logic & interfaces
├── service/
│   └── auth_service.go        # Application logic for authentication
├── handler/
│   └── auth_handler.go        # HTTP handlers
├── infrastructure/
│   └── repository/
│       └── user_repo.go       # Data storage/repository interface implementation
│   └── security/
│       └── jwt.go             # JWT generation/validation logic
└── go.mod

```