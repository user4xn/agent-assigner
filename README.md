# 🧠 Agent Assigner App - Qiscus

A backend service built with Go that handles task queueing and consumption using [Asynq](https://github.com/hibiken/asynq) and Redis. This application supports two operational modes:
- **REST API server** (`rest`): used for adding tasks to the queue
- **Consumer service** (`consumer`): used for processing queued tasks

## 📦 Features

- Task queueing with Asynq and Redis
- Worker/consumer system with retry and concurrency control
- Configurable via `.env` file
- Dockerized for easy deployment

---

## 🚀 Getting Started with Docker

### 📁 1. Clone the repository

```bash
git clone https://github.com/user4xn/agent-assigner.git
cd agent-assigner
```

### ⚙️ 2. Setup the .env
Create a .env file in the root of the project and add the following:
```bash
# Application port
SERVER_PORT=8080

# Redis configuration
REDIS_HOST=localhost
REDIS_PASS=root
REDIS_PORT=6379
REDIS_DB=0

MAX_CUSTOMER_PER_AGENT=2

QISCUS_BASE_URL=https://omnichannel.qiscus.com
QISCUS_APP_ID_CODE=xxx
QISCUS_APP_SECRET=xxx

ASYNQ_PATTERN_CHAT_ASSIGNMENT=chat:assignment
ASYNQ_CONCURRENCY = 1 #num of workers
ASYNQ_RETRY_DELAY = 5 #in(s)
```

### 🐳 3. Run using Docker Compose
Make sure you have Docker and Docker Compose installed. Then, build and run the containers:
```bash
docker compose up --build -d
```

or run it manually, rest and consumer:
```bash
go run main.go
```
```bash
go run main.go -m=consumer
```

### 📌 4. Verify it's working
```bash
docker ps
```
## You should see :
- agent-redis
- agent-rest (API server)
- agent-consumer (task consumer)

### 📬 5. REST API Usage
The REST server is available at:
http://localhost:8080
webhook route was - /api/v1/agent/webhook-assign (POST)

Here's The cURL:
```bash
curl --location --request POST 'http://youraddess/api/v1/agent/webhook-assign'
```
