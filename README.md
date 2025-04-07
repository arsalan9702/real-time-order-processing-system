# Real-Time Order Processing System

A scalable, event-driven application for processing orders in real-time using Go, Kafka, and PostgreSQL.

## Overview

This system implements a microservices architecture to handle order processing with high throughput and reliability. It consists of:

- **Producer Service**: REST API that accepts orders and publishes them to Kafka
- **Consumer Service**: Processes orders from Kafka and updates their status in the database
- **Kafka**: Message broker for reliable communication between services
- **PostgreSQL**: Persistent storage for order data

## Architecture

```
┌─────────────┐         ┌─────────┐         ┌─────────────┐
│             │ publish │         │ consume │             │
│  Producer   ├────────►│  Kafka  ├────────►│  Consumer   │
│  Service    │         │ (Topic) │         │  Service    │
│             │         │         │         │             │
└─────┬───────┘         └─────────┘         └─────┬───────┘
      │                                           │
      │ write                                     │ update
      ▼                                           ▼
┌─────────────────────────────────────────────────────┐
│                                                     │
│                    PostgreSQL                       │
│                                                     │
└─────────────────────────────────────────────────────┘
```

## Prerequisites

- Go 1.23+
- PostgreSQL
- Apache Kafka
- Docker (optional, for containerization)

## Environment Setup

1. Create a `.env` file in the root directory with the following variables:

```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=orders
DB_PORT=5432
```

2. Ensure Kafka is running on `localhost:9092`

## Installation

Clone the repository and install dependencies:

```bash
git clone https://github.com/arsalan9702/real-time-order-processing-system.git
cd real-time-order-processing-system
go mod download
```

## Running the Application

### Start the Producer Service

```bash
go run ./producer.go
```

The producer will start on port 8080.

### Start the Consumer Service

```bash
go run ./consumer.go
```

## API Endpoints

### Create Order

```
POST /orders
```

Request body:

```json
{
  "user_id": 123,
  "product_id": 456,
  "quantity": 2
}
```

Response:

```json
{
  "ID": 1,
  "CreatedAt": "2023-07-14T10:30:00Z",
  "UpdatedAt": "2023-07-14T10:30:00Z",
  "DeletedAt": null,
  "user_id": 123,
  "product_id": 456,
  "quantity": 2,
  "status": "Pending"
}
```

## Data Seeding

The project includes two seeding scripts:

### Basic Seeding

```bash
go run ./seed.go
```

This will generate 10 random orders and send them to the system.

### Advanced Seeding

```bash
go run ./scripts.go --count 1000 --verbose
```

Parameters:

- `--count`: Number of orders to generate (default: 1000)
- `--url`: Base URL of the API (default: "http://localhost:8080")
- `--min-user`: Minimum user ID (default: 1)
- `--max-user`: Maximum user ID (default: 1000)
- `--min-product`: Minimum product ID (default: 1)
- `--max-product`: Maximum product ID (default: 100)
- `--min-quantity`: Minimum quantity (default: 1)
- `--max-quantity`: Maximum quantity (default: 10)
- `--verbose`: Print detailed output (default: false)

## Project Structure

```
├── config/
│   └── config.go         # Database configuration
├── consumer/
│   └── main.go           # Consumer service entry point
├── internal/
│   └── kafka/
│       └── kafka.go      # Kafka producer/consumer implementation
├── models/
│   └── order.go          # Order model definition
├── producer/
│   ├── handler.go        # HTTP handler for orders
│   └── main.go           # Producer service entry point
├── scripts/
│   └── advanced_seeding.go  # Advanced data seeding script
├── seed/
│   └── main.go           # Basic data seeding script
└── .env                  # Environment variables (not tracked by git)
```

## Technical Details

### Order Processing Flow

1. Client submits an order to the Producer's REST API
2. Producer saves the order to PostgreSQL with "Pending" status
3. Producer publishes the order to Kafka topic "orders"
4. Consumer receives the order from Kafka
5. Consumer processes the order and updates its status to "Processed" in PostgreSQL

### Database Schema

The system uses GORM for object-relational mapping with the following schema:

**Orders Table**:
- ID (primary key)
- UserID (integer, not null)
- ProductID (integer, not null)
- Quantity (integer, not null)
- Status (varchar, default: "pending")
- CreatedAt (timestamp)
- UpdatedAt (timestamp)
- DeletedAt (timestamp, nullable)

## Performance Testing

Use the advanced seeding script to simulate high load on the system:

```bash
go run ./scripts --count 10000
```

The script will report throughput metrics after completion.

## Future Enhancements

- Add authentication and authorization
- Implement retry mechanism for failed order processing
- Add monitoring and alerting
- Create Docker Compose setup for easy deployment
- Add unit and integration tests
