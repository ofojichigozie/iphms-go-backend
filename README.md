# IPHMS Go Backend

The Intelligent Patient Health Monitoring System (IPHMS) backend is a Go-based REST API that serves as the core of the IPHMS platform. It handles user authentication, user management, and health vitals data collection from IoT devices.

## 🚀 Features

- **User Authentication**: Secure JWT-based authentication system
- **Role-Based Access Control**: Admin and regular user roles with appropriate permissions
- **User Management**: CRUD operations for user accounts
- **Health Vitals Tracking**: Record and retrieve health metrics from IoT devices
- **IoT Device Integration**: Secure API endpoints for IoT device data submission
- **Database Integration**: PostgreSQL with GORM ORM for data persistence
- **CORS Support**: Configured for frontend integration with support for multiple origins

## 📋 Prerequisites

- Go 1.24 or higher
- PostgreSQL database
- Git

## 🔧 Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ofojichigozie/iphms-go-backend.git
   cd iphms-go-backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the root directory with the following variables:
   ```
   # Database connection string
   DB_URL=postgres://username:password@localhost:5432/iphms_db

   # JWT configuration
   JWT_SECRET=your_jwt_secret_key
   JWT_EXPIRES_IN=24h
   JWT_REFRESH_EXPIRES_IN=168h

   # Server configuration (optional)
   PORT=8080
   ```

4. Run database migrations:
   ```bash
   go run migrate/migrate.go
   ```

5. Start the server:
   ```bash
   go run main.go
   ```

   Alternatively, for development with hot reload:
   ```bash
   air
   ```

## 🏗️ Project Structure

```
iphms-go-backend/
├── controllers/       # Request handlers
├── dtos/              # Data Transfer Objects
├── initializers/      # Application initialization
├── middleware/        # HTTP middleware
├── models/            # Database models
├── repositories/      # Database access layer
├── responses/         # Standardized API responses
├── routes/            # API route definitions
├── services/          # Business logic
├── utils/             # Helper functions
├── migrate/           # Database migration
├── main.go            # Application entry point
└── .env               # Environment variables (not in repo)
```

## 🔐 Authentication

The API uses JWT (JSON Web Tokens) for authentication. There are two types of authentication:

1. **User Authentication**: Using email and password
2. **IoT Device Authentication**: Using device ID and secret

### User Authentication Flow

1. Register a user account (`POST /auth/register`)
2. Login to get access and refresh tokens (`POST /auth/login`)
3. Use the access token in the Authorization header for subsequent requests:
   ```
   Authorization: Bearer <access_token>
   ```
4. Refresh the token when it expires (`POST /auth/refresh`)

### IoT Device Authentication

IoT devices authenticate using custom headers:
```
X-Device-ID: <device_id>
X-Device-Secret: <device_secret>
```

## 🌐 API Endpoints

### Authentication

- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login and get tokens
- `POST /auth/refresh` - Refresh access token

### Users

- `GET /users` - Get all users (admin only)
- `GET /users/:id` - Get user by ID
- `POST /users` - Create a new user
- `PATCH /users/:id` - Update a user
- `DELETE /users/:id` - Delete a user

### Vitals

- `POST /vitals` - Record new vitals (IoT device)
- `GET /vitals` - Get all vitals (filtered by user role)
- `GET /vitals/:id` - Get vitals by ID
- `DELETE /vitals/:id` - Delete vitals

## 🔄 Data Models

### User

```go
type User struct {
    gorm.Model
    Name        string
    Email       string
    Password    string
    DateOfBirth string
    DeviceId    string
    Role        string // "admin" or "user"
}
```

### Vitals

```go
type Vitals struct {
    gorm.Model
    Temperature    float32
    Humidity       float32
    PulseRate      float32
    LightIntensity float32
    UserID         uint
    User           User
}
```

## 🛠️ Development

### Hot Reload

The project includes configuration for [Air](https://github.com/cosmtrek/air), which provides hot reload functionality during development:

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Database Migrations

To update the database schema:

```bash
go run migrate/migrate.go
```

## 🔒 Security Features

- Password hashing using bcrypt
- JWT token-based authentication
- Role-based access control
- HTTPS support
- CORS configuration for frontend access

## 🤝 Integration with IoT Devices

The system is designed to work with Arduino-based IoT devices that collect health metrics:

- Temperature
- Humidity
- Pulse rate
- Light intensity

IoT devices send data to the `/vitals` endpoint using device authentication.

## 📝 License

[MIT License](LICENSE)

## 👥 Contributors

- Your Name - Initial work

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [JWT Go](https://github.com/golang-jwt/jwt)
