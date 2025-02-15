
# 🚀 JarinDeveloper X LineBK Assignment

## 📋 Table of Contents
- [Prerequisites](#-prerequisites)
- [Technologies](#-technologies)
- [Project Structure](#-project-structure)
- [Database Schema Changes](#-database-schema-changes)
  - [Users Table](#-users-table-new-fields)
  - [Account Details Table](#-account-details-table)
  - [Account Flags Table](#-account-flags-table)
  - [Debit Card Status Table](#-debit-card-status-table)
  - [Debit Cards Table](#-debit-cards-table)
- [ER Diagram](#-er-diagram)
- [Database Initialization](#-database-initialization)
- [Environment Configuration](#-environment-configuration)
- [Backend Configuration](#-backend-configuration)
- [API Documentation](#-api-documentation)
- [Testing](#-testing)
- [Quick Start Guide](#-quick-start-guide)
- [Login Credentials](#-login-credentials)
- [K6 Performance Testing](#-k6-performance-testing)


## 📌 Prerequisites  
✅ **Docker**  
✅ **Docker Compose**  

---

## 🛠 Technologies  
- 🐹 **Golang**  
- ⚡ **Fiber**  
- 🛢 **SQLx & MySQL**  
- 🔥 **Redis**  
- 🐳 **Docker**  

---


# 📂 Project Structure

This project follows a modular structure to ensure scalability and maintainability. Below is an overview of the directory and file organization:

```
📦 project-root
│
├── 📂 cmd
│   └── main.go              # Entry point of the application
│
├── 📂 config                # Configuration settings
│   ├── cache.go             # Cache configuration
│   ├── config.go            # General application settings
│   └── database.go          # Database connection settings
│
├── 📂 docs                  # API documentation
│   ├── docs.go              # Swagger documentation setup
│   ├── swagger.json         # Auto-generated Swagger JSON
│   └── swagger.yaml         # OpenAPI specification
│
├── 📂 init-scripts          # Database initialization scripts
│   └── 01_schema.sql        # SQL script for creating tables
│
├── 📂 internal              # Core business logic (not for external import)
│   ├── 📂 account           # Account module
│   │   ├── dto.go           # Data Transfer Object definitions
│   │   ├── handler.go       # HTTP handlers
│   │   ├── mock_repository.go  # Mock repository (for testing)
│   │   ├── mock_service.go     # Mock service (for testing)
│   │   ├── model.go         # Account data models
│   │   ├── repository.go    # Repository layer for database operations
│   │   ├── routes.go        # API route definitions
│   │   ├── service.go       # Business logic
│   │   └── service_test.go  # Unit tests
│   │
│   ├── 📂 auth              # Authentication module
│   ├── 📂 banner            # Banner management module
│   ├── 📂 debit_cards       # Debit card processing
│   ├── 📂 transactions      # Transaction management
│   └── 📂 user              # User management module
│
├── 📂 pkg                   # Shared utilities and helper functions
│   ├── 📂 logs
│   │   └── logs.go          # Logging setup
│   │
│   ├── 📂 middleware
│   │   └── auth.go          # Authentication middleware
│   │
│   └── 📂 utils
│       ├── errs.go          # Custom error handling utilities
│       ├── jwt.go           # JWT token helper functions
│       ├── pagination.go    # Pagination utility
│       ├── response.go      # Standardized API responses
│       └── utils_test.go    # Unit tests for utilities
│
├── 📂 routes
│   └── routes.go            # Main API router
│
├── .dockerignore            # Ignore files for Docker build
├── .env                     # Environment variables file
├── .env-example             # Example environment variables template
├── .gitignore               # Ignore files for Git
├── docker-compose.yaml      # Docker Compose setup
├── Dockerfile               # Docker image configuration
├── go.mod                   # Go module dependencies
├── go.sum                   # Go module checksum file
├── Makefile                 # Build and automation commands
└── README.md                # Project documentation
```

---

### 📌 Notes:
- **`cmd/`** contains the main entry point (`main.go`).
- **`config/`** handles application configurations like database and caching.
- **`internal/`** houses core business logic, services, and handlers.
- **`pkg/`** contains reusable utilities like logging, middleware, and helper functions.
- **`routes/`** registers all API routes.
- **`init-scripts/`** includes SQL scripts to initialize the database when it starts.

This structure ensures a **clean separation of concerns**, making the project easy to **scale** and **maintain**. 🚀

---

## 🏗 Database Schema Changes  

### 🧑‍💻 Users Table (New Fields)  
| 🏷 Field Name | 🛠 Type | 🔎 Reason |
|--------------|--------|----------|
| 📞 `phone_number` | VARCHAR(100) | Store user contact information |
| 🖼 `profile_image` | VARCHAR(255) | Store user profile picture URL |
| 🔑 `pin_code` | VARCHAR(100) | Security and authentication feature |
| 🔒 `password` | VARCHAR(100) | User account authentication |
| ⏳ `created_at` | TIMESTAMP | Track account creation date |

### 🏦 Account Details Table  
| 🏷 Field Name | 🛠 Type | 🔎 Reason |
|--------------|--------|----------|
| 🏷 `account_nickname` | VARCHAR(100) | Optional user-defined account name |

### ⚠️ Account Flags Table  
| 🏷 Field Name | 🛠 Type | 🔎 Reason |
|--------------|--------|----------|
| 📌 `created_at` | TIMESTAMP | Record creation timestamp |
| 🔄 `updated_at` | TIMESTAMP | Track last update time |

### 💳 Debit Card Status Table  
| 🏷 Field Name | 🛠 Type | 🔎 Reason |
|--------------|--------|----------|
| 🚫 `blocked_reason` | VARCHAR(255) | Reason for blocking the card |

### 💰 Debit Cards Table  
| 🏷 Field Name | 🛠 Type | 🔎 Reason |
|--------------|--------|----------|
| 🏷 `card_type` | ENUM('virtual','physical') | Differentiate between virtual and physical cards |
| 📅 `issue_at` | TIMESTAMP | Card issuance date |
| ⏳ `expired_at` | TIMESTAMP | Card expiration date |

---

## 📊 ER Diagram  
![ER Diagram](https://storage.googleapis.com/wirtual-dev/Screenshot%202568-02-10%20at%2023.33.11.png)  

---


## 🗄 Database Initialization

⚡ **Minimal setup required!**  
✅ **Database and tables will be automatically created** when you run `docker-compose up`.  
✅ **Just import data into the tables manually** using the provided SQL file.


### 🔽 Importing Data into MySQL

Once the database and tables are ready, you can import data into the tables using the provided SQL file from **Google Drive**.

📄 **Download data import file:** [Click here to download](https://drive.google.com/file/d/1Ag2qDaFkFnm-VzNqvmlrjKGnPRPayL8G/view?usp=sharing)

### 🛠️ Example Command to Import Data

After downloading the SQL file, you can use the following command to import it into the MySQL container:

```bash
docker exec -i CONTAINER_NAME mysql -u root -pPASSWORD DATABASE_NAME < /path/to/insert.sql
```

**Replace the values as follows:**
- `CONTAINER_NAME` → Name of the MySQL container (check using `docker ps`)
- `PASSWORD` → Your MySQL root password
- `DATABASE_NAME` → Name of the database created in the container
- `/path/to/insert.sql` → Path to the downloaded SQL file (on your local machine)

### 📝 Example:
```bash
docker exec -i assignment_db mysql -u root -pjarindeveloper assignment < ~/Downloads/02_users.sql 
```

This will populate your database with the necessary initial data.  
You're now ready to run the system! 🚀

---

## ⚙️ Environment Configuration  

🚫 **No manual environment setup needed!**  
- All required environment variables are **pre-configured** in `docker-compose.yml`  
- Just run the following command, and you're ready to go:  

```bash
docker-compose up -d --build
```

---

## 🔥 Backend Configuration  

### 📍 Port Configuration  
- **Backend:** `8000`  
- **Redis:** `6379`  
- **MySQL:** `3306`

### 📌 Docker Compose Setup  
🔎 Ensure ports `8000`, `3306`, and `6379` are available before starting:  
```bash
sudo lsof -i :8000
sudo lsof -i :3306
sudo lsof -i :6379
```
✅ If the ports are occupied, stop conflicting services or modify `docker-compose.yml`.  

---

## 📜 API Documentation  

### 🛠 Swagger UI  
- URL: [`localhost:8000/swagger/`](http://localhost:8000/swagger/)  

### 🔑 Authentication Process  

#### 🏷 Login Requirements  
✅ **X-API-KEY**: `123` (pre-configured)  

#### 🔄 Authentication Flow  
1️⃣ Include `X-API-KEY` in the request  
2️⃣ Receive an `access_token`  
3️⃣ Use `access_token` in the header  
   ```http
   Authorization: Bearer [access_token]
   ```
4️⃣ Access the API via Swagger 🚀  

---

## 🧪 Testing  

### 🛠 Unit Testing  
```bash
go test ./...
```

---

## 🚀 Quick Start Guide  

1️⃣ **Clone the repository**  
2️⃣ **Run the backend**  
```bash
docker-compose up -d --build
```
3️⃣ **Open Swagger UI**  
   - [`localhost:8000/swagger/`](http://localhost:8000/swagger/)  

---

## 🔐 Login Credentials  

⚠️ **For testing purposes:**  
- **PIN & Password**: `123456`  
- The default `user_id` is pre-configured in Swagger for easy testing  

✅ **Login Options**  
1️⃣ **Login with PIN**  
2️⃣ **Login with Password**  

---

## 📊 K6 Performance Testing  

![K6 Performance Testing](https://storage.googleapis.com/wirtual-dev/Screenshot%202568-02-11%20at%2000.27.17.png)  

---

