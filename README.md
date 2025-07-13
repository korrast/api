# Korrast API

A REST API for the Korrast project management application built with Go, Gin, and PostgreSQL. 

## 🏗️ Architecture Overview

```mermaid
graph TB
    subgraph "Client Layer"
        C[Client Applications]
    end
    
    subgraph "API Layer"
        direction TB
        R[Router/Gin] --> M[Middleware]
        M --> H[Handlers]
        H --> S[Services]
        
        subgraph "Handlers"
            AH[AuthHandler]
            TH[TableHandler]
        end
        
        subgraph "Services"
            AS[AuthService]
            TS[TableService]
        end
        
        subgraph "Middleware"
            AM[Auth Middleware]
            CORS[CORS Middleware]
        end
    end
    
    subgraph "Data Layer"
        direction TB
        S --> DM[Database Manager]
        DM --> Q[Query Functions]
        Q --> DB[(PostgreSQL)]
        
        subgraph "Database Queries"
            UQ[User Queries]
            TQ[Table Queries]
            LQ[Link Queries]
        end
        
        subgraph "Models"
            UM[User Model]
            TM[Table Model]
            TASK[Task Model]
            COL[Column Model]
            LAB[Label Model]
            MIL[Milestone Model]
        end
    end
    
    subgraph "External"
        JWT[JWT Tokens]
        ENV[Environment Variables]
    end
    
    C --> R
    H --> DM
    AM --> JWT
    DM --> ENV
    Q --> UQ
    Q --> TQ
    Q --> LQ
    
    classDef handler fill:#e1f5fe
    classDef service fill:#f3e5f5
    classDef data fill:#e8f5e8
    classDef external fill:#fff3e0
    
    class AH,TH handler
    class AS,TS service
    class DM,Q,UQ,TQ,LQ,UM,TM,TASK,COL,LAB,MIL data
    class JWT,ENV external
```

## 📁 Project Structure

```
api/
├── main.go                # Application entry point
├── Dockerfile             # Container configuration
├── docker-compose.yaml    # Service orchestration
├── go.mod                 # Go module dependencies
├── 
├── api/                   # Legacy API handlers (to be refactored)
├── dto/                   # Data Transfer Objects
├── handler/               # HTTP request handlers
├── service/               # Business logic layer
├── middleware/            # HTTP middleware (auth, CORS, etc.)
├── database/              # Database connection and queries
├── model/                 # Data models
├── request/               # Request parsing utilities
└── response/              # Response formatting utilities
```

## 🔧 API Endpoints

### Authentication
- `POST /register` - Register a new user
- `POST /login` - Authenticate and get JWT token

### Tables (Protected)
- `POST /api/tables` - Create a new table
- `GET /api/tables` - Get all user tables

### Models Structure

```mermaid
erDiagram
    User ||--o{ Table : owns
    Table ||--o{ Column : contains
    Table ||--o{ Label : has
    Table ||--o{ Milestone : tracks
    Column ||--o{ Task : contains
    Task ||--o{ Label : tagged_with
    Task ||--o| Milestone : belongs_to
    
    User {
        uuid id PK
        string username
        string password
        Table[] tables
    }
    
    Table {
        uuid id PK
        string title
        Column[] columns
        Label[] labels
        Milestone[] milestones
    }
    
    Column {
        uuid id PK
        string name
        string color
        int task_number
        Task[] tasks
    }
    
    Task {
        uuid id PK
        string title
        string description
        Label[] labels
        uuid milestone_id FK
    }
    
    Label {
        uuid id PK
        string title
        string color
    }
    
    Milestone {
        uuid id PK
        string title
        string description
        date end_date
    }
```

## 🛠️ Setup & Installation

### Prerequisites
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL (if running without Docker)

### Environment Variables

Create a `.env` file in the API directory:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=korrast-dev
DB_PASSWORD=123456
DB_NAME=korrast_db

# JWT Configuration
SECRET_TOKEN=your-super-secret-jwt-key-change-this-in-production

# Optional: Enable stub mode for testing
# STUB_MODE=true
```

### Quick Start with Docker

1. **Start the entire stack:**
   ```bash
   docker-compose up -d
   ```

2. **Check service health:**
   ```bash
   docker-compose ps
   docker-compose logs korrast-api
   ```

3. **API will be available at:** `http://localhost:8080`

### Manual Setup

1. **Clone and setup:**
   ```bash
   git clone <repository>
   cd korrast/api
   go mod download
   ```

2. **Start PostgreSQL:**
   ```bash
   cd ../database
   docker-compose up -d
   ```

3. **Run the API:**
   ```bash
   cd ../api
   go run main.go
   ```

### Build Scripts

- **Linux/Mac:** `./build.sh`
- **Windows:** `build.bat`

## 🧪 Testing the API

### Register a new user
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}'
```

### Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "password123"}'
```

### Create a table (requires JWT token)
```bash
curl -X POST http://localhost:8080/api/tables \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"title": "My Project"}'
```

### Create a column
```bash
curl -X POST http://localhost:8080/api/columns \
  -H "Content-Type: application/json" \
  -H "Athorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"table_id": "table_id", "title": "New columns", `optinnal`: "color": "#FAE536"}' 
```

## 📄 License

MIT License - see LICENSE file for details.

## Author

This project is entirely made by me (ASTOLFI Vincent). I suggest you to check on my github profile if you want to see the other project I've done for my studies or the ones I do in my free time. 
