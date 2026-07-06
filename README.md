# Nexus

Nexus is a production-ready microservices backend built in **Go (Golang)**. It demonstrates a modern event-driven architecture using **gRPC** for internal communication and **GraphQL** as the API Gateway. 

## 🏗️ Architecture & Data Flow

![Nexus Architecture Diagram](./assets/architecture.png)

Nexus models an e-commerce platform divided into three independent microservices that own their own databases:

1. **Account Service (PostgreSQL):** Manages user data.
2. **Catalog Service (Elasticsearch):** Manages products and powerful full-text search.
3. **Order Service (PostgreSQL + RabbitMQ):** Manages purchases. When a user buys something, it validates the user via gRPC, processes the order, and asynchronously publishes an event to RabbitMQ.

**The Gateway:** Clients never talk to the databases directly. They communicate exclusively with the **GraphQL Gateway**, which stitches the data together by firing off high-speed gRPC requests to the internal services.

## 🚀 Tech Stack

- **Language:** Go 1.20
- **API Gateway:** GraphQL (`gqlgen`)
- **Internal RPC:** gRPC & Protocol Buffers
- **Message Broker:** RabbitMQ
- **Databases:** PostgreSQL (Relational) & Elasticsearch v7 (Search)
- **Infrastructure:** Docker & Docker Compose

## 🚦 Quick Start (Local Development)

You only need **Docker** installed. No local Go environment is required.

1. **Spin up the cluster:**
   ```bash
   docker compose up -d --build
   ```
2. **Access the Dashboards:**
   - **GraphQL API & Playground:** [http://localhost:8000/playground](http://localhost:8000/playground)
   - **RabbitMQ Dashboard:** [http://localhost:15672](http://localhost:15672) *(Login: `guest` / `guest`)*

## 🔍 Example GraphQL API Usage

Because the backend uses a GraphQL Gateway, a frontend client can fetch data from all three microservices (Accounts, Orders, and Products) in a single unified request. 

Try pasting this in the playground to see how it resolves data across the internal gRPC services:

```graphql
query {
  accounts {
    id
    name
    orders {
      id
      createdAt
      totalPrice
      products {
        name
        price
        quantity
      }
    }
  }
}
```

## 🧹 Clean Up

To gracefully stop the containers and wipe the databases (useful if you want a fresh database instance):
```bash
docker compose down -v
```
