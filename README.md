# Golang Backend with Hexagonal Architecture
This repository implements a backend application in Golang using the hexagonal (or ports and adapters) architecture. This architecture promotes separation of concerns, allowing for easy testing, maintenance, and replacement of components.

## Project Structure

The project is organized into three primary layers, following the hexagonal architecture principles:

```bash
.
├── cmd                   # Application entry point and wiring (Driver Enry Point)
│   └── main.go           # Main application file
├── domain                # Core business logic (Entities, Repositories, Services) (Domain Layer)
│   ├── model             # Domain models (Entities)
│   └── repository        # Repository interfaces for persistence (Ports)
├── infrastructure        # Implementations of interfaces (Repositories, HTTP handlers, etc.) (Adapters)
│   ├── security          # security related code (Secondary Adapters)
│   └── persistence       # Database access layer (Secondary Adapters)
├── service               # Business services implementing use cases (Application Layer)
├── handler               # HTTP handlers (Primary Adapters)
├── router                # Application router
├── config                # Configuration files
│   └── database.go       # Database configuration
├── go.mod                # Go module file
├── go.sum                # Go sum file
├── docker-compose.yml    # Docker Compose configuration
├── Dockerfile            # Dockerfile for building the application
└── Dockerfile.dev        # Dockerfile for development
├── .env                  # Environment variables
└── README.md             # Project documentation
```

## Directory Breakdown
1. cmd/ - Driver / Entry Point: The entry point of the application. This layer initiates the application by configuring dependencies, starting the server, and registering routes.
2. domain/ - Domain Layer: domain/model - Entities: Defines core business entities, such as User. These models represent the domain and encapsulate business rules. domain/repository - Ports: Defines repository interfaces that specify the operations needed for persistence (e.g., UserRepository). Ports are interfaces that can be implemented by any adapter.

3. infrastructure/ - Adapters: Implements the interfaces defined in the domain layer to connect to external systems like databases or APIs.
infrastructure/persistence - Secondary Adapters: Contains concrete implementations of repository interfaces for persistence, such as database access layers for SQL or NoSQL databases. infrastructure/security - Secondary Adapter: Implements security-related code, such as token generation or hashing, which may be used across the application.

4. handler/ - Primary Adapters: Contains HTTP handlers that serve as the interface between the client and the service layer. Handlers receive requests, process them, and call the appropriate service methods.

5. service/ - Application Layer: This layer contains business services implementing the application’s use cases, independent of the infrastructure. It coordinates business rules and orchestrates the use of different domain interfaces.
UserService - Defines use cases for user operations such as registering, retrieving, updating, and deleting users. This service is an intermediary between handlers (Primary Adapters) and repositories (Secondary Adapters), adhering to the Port interfaces in domain/repository.




## Key Files
main.go: The main file that sets up the application, initializes dependencies (such as the database connection and router), and starts the HTTP server.
service/user_service.go: Implements the business logic for handling user-related functionality. The service uses dependency injection to interact with UserRepository for data persistence.
infrastructure/handler/user_handler.go: Defines the HTTP endpoints related to user operations, like registering and retrieving users. It communicates with UserService to process requests.
Example of Registering a New User
Request: A client sends a POST request to /register with a JSON payload containing username and password.
Handler: The UserHandler decodes the request, validates it, and passes the data to UserService.
Service: The UserService processes the request, checking if the username already exists and, if not, creating a new user.
Repository: The UserService uses UserRepository to save the new user to the database.
Response: The UserHandler returns a success response with the created user’s information.
Error Handling
The project includes custom error handling for JSON responses, using jsonErrorResponse to return standardized error messages and status codes. For instance, if a username already exists, UserService returns a specific error that is handled by UserHandler to send a 409 Conflict status.

Dependencies
This project uses the following key dependencies:

Gorilla Mux: For HTTP request routing.
Go-Playground Validator: For validating incoming request payloads.
Database Driver: Add your preferred driver here (e.g., PostgreSQL, MySQL, or MongoDB).
Getting Started
Install Dependencies:

bash
Copiar código
go mod tidy
Run the Application:

bash
Copiar código
go run main.go
Endpoints:

POST /register: Register a new user.
GET /users/{id}: Retrieve a user by ID.
Future Improvements
Add more comprehensive tests for each layer.
Integrate authentication and authorization.
Add more service methods and refine err



## Layered Role Summary
### Primary Adapters: 
Modules that serve as entry points into the application, handling user interactions and triggering use cases. For example, handler handles HTTP requests and calls service methods.
### Application Layer:
Implements business logic and coordinates interactions between ports and adapters. Services in this layer define and manage application-specific rules.
### Ports: 
Defined as interfaces within the domain layer. Ports outline the essential operations required by the application, which can be fulfilled by any adapter.
### Secondary Adapters:
Implement the interfaces defined in the domain layer, interacting with external resources such as databases, APIs, or security tools.
### Domain Layer:
The core of the application, where the essential business entities, rules, and logic reside. It remains isolated from infrastructure, allowing adapters to plug in as needed.

#### This structure, using hexagonal architecture terminology, reinforces clean boundaries between layers and ensures the core business logic remains independent of infrastructure and external dependencies.