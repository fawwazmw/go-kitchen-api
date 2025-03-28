# Go Kitchen API

A microservice-based kitchen order management system built with Go, gRPC, and HTTP APIs.

## Project Overview

Go Kitchen API is a distributed system consisting of two main services:

1. **Orders Service**: Manages order creation and retrieval
2. **Kitchen Service**: Provides a web interface to interact with the orders

The system uses gRPC for inter-service communication and HTTP for external API endpoints.

## Architecture

```
┌───────────────┐     gRPC      ┌───────────────┐
│ Kitchen       │◄──────────────┤ Orders        │
│ Service       │               │ Service       │
│ (:1024)       │──────────────►│ (:9000 gRPC)  │
└───────────────┘               │ (:8000 HTTP)  │
                                └───────────────┘
```

### Technology Stack

- **Go**: Primary programming language
- **Protocol Buffers**: For defining service contracts
- **gRPC**: For efficient inter-service communication
- **HTTP**: For external API endpoints
- **HTML Templates**: For basic web UI

## Project Structure

```
├── protobuf/
│   └── orders.proto         # Protocol Buffer definitions
├── services/
│   ├── common/              # Shared code
│   │   ├── genproto/        # Generated Protocol Buffer code
│   │   └── util/            # Utility functions
│   ├── kitchen/             # Kitchen Service
│   │   ├── main.go          # Entry point
│   │   └── http.go          # HTTP server implementation
│   └── orders/              # Orders Service
│       ├── handler/         # Request handlers
│       │   └── orders/      # Order-specific handlers
│       ├── service/         # Business logic
│       ├── types/           # Type definitions
│       ├── grpc.go          # gRPC server implementation
│       ├── http.go          # HTTP server implementation
│       └── main.go          # Entry point
├── go.mod                   # Go module file
├── go.sum                   # Go sum file
└── Makefile                 # Build and run commands
```

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Protocol Buffers compiler (protoc)
- Go plugins for Protocol Buffers

### Installation

1. Clone the repository:

```bash
git clone https://github.com/fawwazmw/go-kitchen-api.git
cd go-kitchen-api
```

2. Generate Protocol Buffer code:

```bash
make gen
```

### Running the Services

1. Start the Orders Service:

```bash
make run-orders
```

This will start the Orders Service with:
- HTTP server on port 8000
- gRPC server on port 9000

2. Start the Kitchen Service:

```bash
make run-kitchen
```

This will start the Kitchen Service with:
- HTTP server on port 1024

## API Documentation

### Orders Service - HTTP API

#### Create an Order

```
POST /orders
```

**Request Body:**

```json
{
  "customerID": 24,
  "productID": 3123,
  "quantity": 2
}
```

**Response:**

```json
{
  "status": "success"
}
```

### Orders Service - gRPC API

The Orders Service exposes the following gRPC methods:

#### CreateOrder

Creates a new order in the system.

**Request:**
```protobuf
message CreateOrderRequest {
  int32 customerID = 1;
  int32 productID = 2;
  int32 quantity = 3;
}
```

**Response:**
```protobuf
message CreateOrderResponse {
  string status = 1;
}
```

#### GetOrders

Retrieves orders for a specific customer.

**Request:**
```protobuf
message GetOrdersRequest {
  int32 customerID = 1;
}
```

**Response:**
```protobuf
message GetOrderResponse {
  repeated Order orders = 1;
}
```

### Kitchen Service - Web UI

The Kitchen Service provides a web interface to view orders at:

```
http://localhost:1024/
```

## Development

### Adding New Features

1. Define new Protocol Buffer messages and service methods in `protobuf/orders.proto`
2. Regenerate Protocol Buffer code with `make gen`
3. Implement new handlers in the appropriate service
4. Update the service interfaces as needed

### Data Model

```protobuf
message Order {
  int32 OrderID = 1;
  int32 CustomerID = 2;
  int32 ProductID = 3;
  int32 Quantity = 4;
}
```

## Known Issues and Limitations

- Currently uses in-memory storage, not suitable for production
- Limited error handling and validation
- No authentication or authorization mechanisms

## Future Enhancements

- Add persistent storage (database)
- Implement user authentication
- Add product service to manage kitchen products
- Improve error handling and validation
- Add logging and monitoring
- Containerize services with Docker

## License

This project is licensed under the [MIT License](https://opensource.org/license/MIT) - see the LICENSE file for details.