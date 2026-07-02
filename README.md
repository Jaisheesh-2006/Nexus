# Nexus

Nexus is a high-performance, production-ready microservices backend built in **Go (Golang)**. It demonstrates a modern event-driven architecture using **gRPC** for internal communication and **GraphQL** as the API Gateway. 

## 🚀 Tech Stack

- **Language:** Go 1.20
- **API Gateway:** GraphQL
- **RPC:** gRPC & Protocol Buffers
- **Message Broker:** RabbitMQ
- **Databases:** PostgreSQL (Relational) & Elasticsearch v7 (Search)
- **Infrastructure:** Docker & Docker Compose

## 🏗️ Architecture

Nexus models an e-commerce platform divided into three independent microservices:

1. **Account Service (PostgreSQL):** Manages users.
2. **Catalog Service (Elasticsearch):** Manages products and powerful full-text search.
3. **Order Service (PostgreSQL + RabbitMQ):** Manages purchases. When an order is created, it asynchronously publishes an event to RabbitMQ that the Account service listens to.

Clients communicate only with the **GraphQL Gateway**, which stitches the data together by firing off high-speed gRPC requests to the internal services.

## 🚦 Quick Start

You only need Docker installed to run this project.

1. **Start the cluster:**
   ```bash
   docker compose up -d --build
   ```
2. **Access the API:**
   - GraphQL Playground: [http://localhost:8000/playground](http://localhost:8000/playground)
   - RabbitMQ Dashboard: [http://localhost:15672](http://localhost:15672) *(Login: guest/guest)*
   
## 🧹 Clean Up

To gracefully stop the containers and wipe the databases:
```bash
docker compose down -v
```

## 📝 License
MIT License
