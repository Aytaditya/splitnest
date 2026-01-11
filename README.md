# SplitNest – Expense Sharing Backend (Microservices)

SplitNest is a **Splitwise-like expense sharing backend system** built using **Go microservices**, **SQLite**, **RabbitMQ**, and an **NGINX API Gateway**.

The project focuses on **clean service boundaries**, **correct expense math**, **transaction safety**, and **event-driven architecture**, rather than over-engineering.

---

## 🚀 Features

- User registration & login (JWT-based authentication)
- Group creation and member management
- Expense creation with **equal split logic**
- Net balance calculation per group
- Asynchronous notifications using RabbitMQ
- Event-driven extensibility
- API Gateway using NGINX
- Independent database per service
- Dockerized setup

---

## 🧱 Architecture Overview

### Core Services
- **User Service** – user registration, login, identity
- **Group Service** – group creation & membership
- **Expense Service** – expense tracking & balance calculation

### Supporting Services
- **Notification Service** – consumes events asynchronously
- **API Gateway** – routes requests to services (NGINX)

---

## 🔁 Communication Pattern

- **Synchronous (HTTP)**  
  Used for core business logic (e.g., Expense → Group)

- **Asynchronous (RabbitMQ)**  
  Used for side effects like notifications

---

## 📦 Tech Stack

- **Language:** Go
- **Databases:** SQLite (one per service)
- **Messaging:** RabbitMQ
- **Auth:** JWT
- **Gateway:** NGINX
- **Containerization:** Docker & Docker Compose

---

## 📂 Project Structure

```text
.
├── api-gateway
│   └── nginx.conf
├── docker-compose.yaml
├── user-service
├── group-service
├── expense-service
├── notification-service
└── Makefile
````

Each service follows a clean internal structure:

* `cmd/` – entry point
* `internal/http` – HTTP handlers
* `internal/storage` – database logic
* `internal/middleware` – JWT validation
* `internal/types` – request/response models

---

## 🔐 Authentication

* JWT-based authentication
* Token passed via header:

```
Authorization: Bearer <JWT_TOKEN>
```

Each service validates JWT independently.

---

## 📡 API Documentation

### 🧑 User Service (Port: `8081`)

#### Health Check

```
GET /
```

---

#### Register User

```
POST /register
```

```json
{
  "username": "aditya",
  "email": "aditya@gmail.com",
  "password": "secret123"
}
```

---

#### Login

```
POST /login
```

```json
{
  "email": "aditya@gmail.com",
  "password": "secret123"
}
```

Response:

```json
{
  "token": "<jwt-token>"
}
```

---

#### Find User by Username (Internal)

```
GET /find-user/{username}
```

---

## 👥 Group Service (Port: `8082`)

#### Health Check

```
GET /
```

---

#### Create Group

```
POST /create-group
```

Headers:

```
Authorization: Bearer <JWT>
```

```json
{
  "name": "Goa Trip"
}
```

---

#### Add Member to Group

```
POST /add-members/{groupId}
```

```json
{
  "username": "rohan"
}
```

---

#### Get User Groups

```
GET /user-groups
```

Returns all groups where the user is a **member or admin**.

---

#### Get Group Members

```
GET /group-members/{groupId}
```

Response:

```json
[
  { "user_id": 1 },
  { "user_id": 2 },
  { "user_id": 4 }
]
```

Used internally by Expense Service.

---

## 💸 Expense Service (Port: `8083`)

#### Health Check

```
GET /
```

---

#### Add Expense

```
POST /add-expense/{groupId}
```

```json
{
  "amount": 1500
}
```

Behavior:

* Equal split among group members
* Transaction-safe writes
* Publishes `expense.created` event to RabbitMQ

---

#### Get Group Expenses

```
GET /get-expense/{groupId}
```

---

## 🧮 Balance Logic

For each expense:

```
net_balance = money_paid - money_owed
```

* Positive → user gets money
* Negative → user owes money
* Balances always sum to **0**

Balances are **incrementally updated**, not recomputed.

---

## 🔔 Events (RabbitMQ)

### Exchange

* Name: `events`
* Type: `topic`

### Published Events

#### `expense.created`

```json
{
  "event": "expense.created",
  "group_id": 3,
  "expense_id": 2,
  "paid_by": 4,
  "amount": 900
}
```

#### `group.member_added`

```json
{
  "event": "group.member_added",
  "group_id": 3,
  "user_id": 4
}
```

Events are published **after successful DB transaction commit**.

---

## 🛠 Running the Project

### Prerequisites

* Docker
* Docker Compose

### Start all services

```bash
docker compose up --build
```

### Stop services

```bash
docker compose down
```

---

## 🧪 Future Enhancements (Not Implemented)

* Analytics Service (event-driven)
* Settlement optimization (who pays whom)
* gRPC for internal calls
* Idempotency keys
* Observability & metrics
