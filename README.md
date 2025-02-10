# JarinDeveloper X LineBK Assignment

## Prerequisites
- Docker
- Docker Compose


## Technologies
- Golang
- Fiber
- SQLx
- MySQL
- Redis
- Docker

## Database Schema Changes

### Users Table Additional Fields

| Field Name | Type | Reason |
|-----------|------|--------|
| `phone_number` | VARCHAR(100) | Allow user contact information |
| `profile_image` | VARCHAR(255) | Store user profile picture URL |
| `pin_code` | VARCHAR(100) | Authentication/Security feature |
| `password` | VARCHAR(100) | User account authentication |
| `created_at` | TIMESTAMP | Track user account creation time |

### Account Details Table
| Field Name | Type | Reason |
|-----------|------|--------|
| `account_nickname` | VARCHAR(100) | Optional user-defined account name |

### Account Flags Table
| Field Name | Type | Reason |
|-----------|------|--------|
| `created_at` | TIMESTAMP | Record creation timestamp |
| `updated_at` | TIMESTAMP | Track last update time |

### Debit Card Status Table
| Field Name | Type | Reason |
|-----------|------|--------|
| `blocked_reason` | VARCHAR(255) | Reason for card blocking |

### Debit Cards Table
| Field Name | Type | Reason |
|-----------|------|--------|
| `card_type` | ENUM('virtual','physical') | Distinguish card type |
| `issue_at` | TIMESTAMP | Card issuance date |
| `expired_at` | TIMESTAMP | Card expiration date |


## ER Diagram
![ER Diagram](https://storage.googleapis.com/wirtual-dev/Screenshot%202568-02-10%20at%2023.33.11.png)

## Database Initialization

### Schema and Initial Data

**Important Note**: 
- Large database schema and initial rows were imported directly to Google Cloud Platform
- No local initialization required during Docker setup
- This approach ensures clean, pre-populated database without manual import steps


## Environment Configuration

**No Manual Environment Setup Required**
- All necessary environment variables pre-configured in `docker-compose.yml`
- Just run `docker-compose up` and you're ready to go

## Backend Configuration

### Port Configuration
- **Backend Port**: 8080
- **Redis Port**: 6379

### Docker Compose Setup
- Ensure ports 8080 and 6379 are available on your local machine
- No additional environment setup required
- Only prerequisite is Docker installed

### Port Availability Check
```bash
# Check if ports 8080 and 6379 are free
sudo lsof -i :8080
sudo lsof -i :6379
```

**Note**: If ports are in use, stop conflicting services or change port mappings in `docker-compose.yml`

## API Documentation

### Swagger UI
- Access: `localhost:8080/swagger/`

### Authentication Process

#### Login Requirements
1. **X-API-KEY**
   - Pre-configured in environment variables
   - Value: `123`

#### Authentication Flow
1. Provide X-API-KEY during login
2. Receive `access_token`
3. Use `access_token` in Authorization header
   - Format: `Bearer [access_token]`

### Swagger Usage
- Include `X-API-KEY` and `Authorization` header in API requests
- Refer to Swagger UI for endpoint details


## Testing

### Unit Testing
```bash
go test ./...

```


## Quick Start Guide

### Steps
1. Pull repository code
2. Run `docker-compose up -d`
3. Open Swagger UI: `localhost:8080/swagger/`

## Login Credentials

**Important Note**: 
- For ALL users, both PIN and password are set to `123456`
- This is a default setting to simplify testing and review process

### Login Options
- Login with PIN
- Login with Password

### Note
- Default `user_id` pre-configured in Swagger for easy testing