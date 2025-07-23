# SwiftPlay Backend API

> 🎮 ❤️ **Find Your Player Two** - Where gamers connect and compete together

[![Go Version](https://img.shields.io/badge/Go-1.23.1-blue.svg)](https://golang.org)
[![Framework](https://img.shields.io/badge/Framework-Gin-green.svg)](https://gin-gonic.com/)
[![ORM](https://img.shields.io/badge/ORM-GORM-red.svg)](https://gorm.io/)
[![Database](https://img.shields.io/badge/Database-PostgreSQL-blue.svg)](https://postgresql.org)

## 📋 Overview

This is the RESTful API for the Swiftplay app designed specifically for the gaming community. It enables players to create profiles, find compatible teammates based on rank and preferences, match with other players, and communicate through secure messaging. Built with performance, scalability, and security in mind.

## ✨ Features

- **🎯 Gaming Profiles** - Comprehensive player profiles with multi-game rank integration
- **🎮 Multi-Game Support** - Currently supports Valorant, with more games planned
- **🔍 Smart Matching** - Connect players based on game ranks, location, and preferences  
- **💬 Secure Messaging** - Real-time communication between matched players
- **📱 Mobile-First API** - Designed specifically for React Native mobile application
- **🛡️ Authentication** - Secure user registration and login system
- **📊 Auto-Migration** - Database schema management with GORM
- **🚀 High Performance** - Built with Gin framework for optimal speed

## 🛠 Tech Stack

### Backend API (This Repository)
- **Backend Framework**: [Gin](https://gin-gonic.com/) - High-performance HTTP web framework
- **Database ORM**: [GORM](https://gorm.io/) - Developer-friendly ORM library
- **Database**: PostgreSQL 15+
- **Language**: Go 1.23.1+
- **Authentication**: JWT (Ready for implementation)
- **Architecture**: Clean Architecture with separated concerns

### Frontend Mobile App (Separate Repository)
- **Mobile Framework**: React Native
- **Platforms**: iOS & Android
- **State Management**: Redux/Context API
- **Navigation**: React Navigation

## 📋 Prerequisites

Before running this application, ensure you have:

- **Go 1.23.1+** installed ([Download](https://golang.org/dl/))
- **PostgreSQL 15+** running ([Download](https://www.postgresql.org/download/))
- **Git** for version control

## 🚀 Quick Start

### 1. Clone & Setup
```bash
# Clone the repository
git clone https://github.com/1ShoukR/swiftplay-backend.git
cd swiftplay-backend

# Install dependencies
go mod download
```

### 2. Database Setup
```bash
# Create PostgreSQL database
createdb swiftplay_db

# Or using psql
psql -c "CREATE DATABASE swiftplay_db;"
```

### 3. Environment Configuration
```bash
# Create environment file
touch .env

# Add your configuration (see Environment Variables section)
```

### 4. Run the Application
```bash
# Development mode
go run cmd/main.go

# Or build and run
go build -o swiftplay cmd/main.go
./swiftplay
```

The API server will start on **http://localhost:8081** 🎉

## 🧪 Quick Test

Once your server is running, test the API with these curl commands:

### **Health Check**
```bash
curl http://localhost:8081/health
```
**Expected Response:**
```json
{"status":"OK"}
```

### **Create a Gaming User**
```bash
curl -X POST http://localhost:8081/api/users/create \
  -H "Content-Type: application/json" \
  -d '{
    "user": {
      "username": "valorant_player",
      "email": "player@swiftplay.com",
      "password_hash": "your_hashed_password_here",
      "auth_level": "user"
    },
    "profile": {
      "first_name": "Alex",
      "last_name": "Chen",
      "bio": "Diamond Valorant player looking for competitive teammates",
      "city": "San Francisco",
      "country": "USA",
      "game_ranks": {
        "valorant": "Diamond 2",
        "csgo": "Legendary Eagle Master"
      }
    }
  }'
```

**Expected Response:**
```json
{
  "message": "User and gaming profile created successfully",
  "user": {
    "user_id": 1,
    "username": "valorant_player",
    "email": "player@swiftplay.com",
    "auth_level": "user",
    "created_at": "2025-01-22T22:47:38Z"
  },
  "profile": {
    "profile_id": 1,
    "user_id": 1,
    "first_name": "Alex",
    "game_ranks": {"valorant": "Diamond 2", "csgo": "Legendary Eagle Master"}
  },
  "supported_games": ["valorant", "csgo", "apex_legends", "lol", "rocket_league"]
}
```

### **Test Authentication Endpoints**
```bash
# Login endpoint (placeholder)
curl http://localhost:8081/api/auth/login

# Register endpoint (placeholder)  
curl -X POST http://localhost:8081/api/auth/create
```

## 📁 Project Structure

```
swiftplay-backend/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── database/            # Database connection & configuration
│   │   └── database.go      # GORM setup with auto-migration
│   ├── handlers/            # HTTP request handlers
│   │   ├── auth.go          # Authentication endpoints
│   │   └── user.go          # User management endpoints
│   ├── models/              # Data models and schemas
│   │   ├── user.go          # User & Profile models
│   │   ├── match.go         # Matching system model
│   │   └── message.go       # Messaging model
│   └── server/              # Server configuration
│       ├── server.go        # Gin server setup
│       └── routes/          # API routing
│           ├── router.go    # Main router configuration
│           └── auth.go      # Authentication routes
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksums
└── README.md               # Project documentation
```

## 🔧 Environment Variables

Create a `.env` file in the root directory:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=swiftplay_db
DB_SSLMODE=disable

# Server Configuration  
PORT=8081
GIN_MODE=release        # Use 'debug' for development

# Security (Add these for production)
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRY=24h

# API Configuration
API_VERSION=v1
RATE_LIMIT=100          # Requests per minute
```

## 🗄 Database Schema

### Core Models

#### Users Table
```sql
CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    auth_level VARCHAR(20) DEFAULT 'user',  -- user, admin, super_admin, engineer
    soft_delete BOOLEAN DEFAULT false,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
```

#### Profiles Table
```sql
CREATE TABLE profiles (
    profile_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    gender VARCHAR(20),
    date_of_birth TIMESTAMPTZ,
    bio TEXT,
    city VARCHAR(100),
    country VARCHAR(100),
    game_ranks JSONB,        -- Store ranks for multiple games
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);
```

#### Matches Table
```sql
CREATE TABLE matches (
    match_id BIGSERIAL PRIMARY KEY,
    user_id_1 BIGINT REFERENCES users(user_id),
    user_id_2 BIGINT REFERENCES users(user_id),
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMPTZ
);
```

#### Messages Table
```sql
CREATE TABLE messages (
    message_id BIGSERIAL PRIMARY KEY,
    match_id BIGINT REFERENCES matches(match_id),
    sender_id BIGINT REFERENCES users(user_id),
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ,
    read_at TIMESTAMPTZ
);
```

## 🔌 API Endpoints

### Base URL: `http://localhost:8081`

### Health Check
```http
GET /health
```
**Response:**
```json
{
  "status": "OK"
}
```

### Authentication
```http
GET  /api/auth/login          # User login (placeholder)
POST /api/auth/create         # User registration (placeholder)
```

### User Management
```http
POST /api/users/create        # Create user with profile
```

**Create User Request:**
```json
{
  "user": {
    "username": "gaming_pro",
    "email": "player@example.com",
    "password_hash": "hashed_password_here",
    "auth_level": "user"
  },
  "profile": {
    "first_name": "John",
    "last_name": "Doe",
    "bio": "Competitive gamer looking for teammates",
    "city": "Los Angeles",
    "country": "USA",
    "game_ranks": {
      "valorant": "Immortal 2"
    }
  }
}
```

**Success Response:**
```json
{
  "message": "User and profile created successfully",
  "user": {
    "user_id": 1,
    "username": "valorant_pro",
    "email": "player@example.com",
    "created_at": "2025-01-22T22:47:38Z"
  },
  "profile": {
    "profile_id": 1,
    "user_id": 1,
    "first_name": "John",
         "game_ranks": {"valorant": "Immortal 2"}
  }
}
```

## 🧪 Testing

### Run Tests
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Test specific package
go test ./internal/handlers -v
```

### Manual API Testing

#### Using cURL
```bash
# Health check
curl http://localhost:8081/health

# Auth endpoints
curl http://localhost:8081/api/auth/login
curl -X POST http://localhost:8081/api/auth/create

# Create user with profile
curl -X POST http://localhost:8081/api/users/create \
  -H "Content-Type: application/json" \
  -d '{
         "user": {
       "username": "test_player",
       "email": "test@example.com", 
       "password_hash": "hashed_password",
       "auth_level": "user"
     },
         "profile": {
       "first_name": "Test",
       "game_ranks": {"valorant": "Gold 3"},
       "bio": "Looking for teammates!"
     }
  }'
```

## 🐳 Docker Support

### Dockerfile
```dockerfile
FROM golang:1.23.1-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
EXPOSE 8081

CMD ["./main"]
```

### Docker Compose
```yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8081:8081"
    environment:
      - DB_HOST=db
      - DB_NAME=swiftplay_db
    depends_on:
      - db
      
  db:
    image: postgres:15
    environment:
      POSTGRES_DB: swiftplay_db
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
```

## 🚀 Deployment

### Production Checklist
- [ ] Set secure environment variables
- [ ] Enable HTTPS with TLS certificates
- [ ] Configure CORS for frontend domains
- [ ] Implement rate limiting
- [ ] Set up monitoring and logging
- [ ] Configure database connection pooling
- [ ] Enable database backups
- [ ] Set up CI/CD pipeline
- [ ] Configure load balancing (if needed)

### Environment Setup
```bash
# Set production environment
export GIN_MODE=release
export DB_SSLMODE=require

# Use environment-specific configs
export DB_HOST=your-production-db-host
export JWT_SECRET=your-super-secure-secret
```

## 🛡️ Security Features

- **Role-Based Access Control**: Multi-level authentication system
  - `user` - Standard players (default for new registrations)
  - `admin` - Administrative access
  - `super_admin` - Full system access
  - `engineer` - Development and system maintenance access
- **Password Hashing**: Secure password storage (ready for bcrypt implementation)
- **SQL Injection Protection**: GORM ORM prevents SQL injection attacks
- **Input Validation**: JSON binding with validation
- **Database Transactions**: Atomic operations for data consistency
- **Soft Deletes**: Users are soft-deleted, not permanently removed

## 📊 Performance Features

- **Gin Framework**: High-performance HTTP router
- **Database Indexing**: Optimized queries with proper indexing
- **Connection Pooling**: Efficient database connection management
- **Auto-Migration**: Schema management without downtime
- **Graceful Shutdown**: Proper cleanup on server termination

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork** the repository
2. **Create** your feature branch (`git checkout -b feature/amazing-feature`)
3. **Commit** your changes (`git commit -m 'Add some amazing feature'`)
4. **Push** to the branch (`git push origin feature/amazing-feature`)
5. **Open** a Pull Request

### Development Guidelines
- Run `go fmt` before committing
- Ensure all tests pass (`go test ./...`)
- Follow Go naming conventions
- Add tests for new features
- Update documentation for API changes

### Code Style
```bash
# Format code
go fmt ./...

# Lint code (install golangci-lint first)
golangci-lint run

# Vet code
go vet ./...
```

## 🔮 Roadmap

### Phase 1 (Current)
- [x] Basic API structure with Gin + GORM
- [x] User and Profile management
- [x] Database auto-migration
- [x] Basic authentication endpoints

### Phase 2 (Next)
- [ ] JWT authentication implementation
- [ ] Password hashing with bcrypt
- [ ] Input validation and sanitization
- [ ] Rate limiting middleware

### Phase 3 (Future)
- [ ] Multi-game support expansion (CS2, Apex Legends, LoL, etc.)
- [ ] Advanced matching algorithm implementation
- [ ] Real-time messaging with WebSockets
- [ ] File upload for profile pictures
- [ ] Push notifications for React Native app
- [ ] Email verification system

### Phase 4 (Advanced)
- [ ] Machine learning for better matching across games
- [ ] Integration with game APIs (Riot, Steam, etc.)
- [ ] Advanced analytics and metrics dashboard
- [ ] Tournament and event creation features
- [ ] Voice chat integration
- [ ] Team formation for competitive gaming

## 📞 Support & Contact

- **GitHub Issues**: [Report bugs or request features](https://github.com/1ShoukR/swiftplay-backend/issues)
- **Discussions**: [Community discussions](https://github.com/1ShoukR/swiftplay-backend/discussions)
- **Project Lead**: [@1ShoukR](https://github.com/1ShoukR)

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **[Gin Framework](https://gin-gonic.com/)** - High-performance HTTP framework
- **[GORM](https://gorm.io/)** - Fantastic ORM library for Go
- **[PostgreSQL](https://postgresql.org)** - Powerful, open-source database
- **[React Native](https://reactnative.dev/)** - Mobile development framework
- **Gaming Community** - Inspiration for building connections through shared gaming experiences

---

<div align="center">

**Made with ❤️ and 🎮 for the gaming community**

[⬆ Back to Top](#swiftplay-backend-api)

</div>