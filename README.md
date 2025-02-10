
# 🚀 JarinDeveloper X LineBK Assignment  

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

⚡ **No manual setup required!**  
✅ The database schema and initial data are already imported to **Google Cloud Platform**  
✅ No need to import data locally  

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

### 📌 Docker Compose Setup  
🔎 Ensure ports `8000` and `6379` are available before starting:  
```bash
sudo lsof -i :8000
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
docker-compose up -d
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
