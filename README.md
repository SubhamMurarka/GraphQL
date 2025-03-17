# Go GraphQL Service for Best Material and Supplier Search

## Overview
This project is a Go-based GraphQL service that fetches and filters material and supplier details from external APIs. It helps customers find the best materials based on price, rating, and quality within their locality.

## Tech Stack
- **Golang** (Go)
- **gqlgen** (GraphQL library)
- **Redis** (Caching)
- **Goroutines** (Concurrent API calls for MaterialsAPI, SupplierAPI)
- **Repository Pattern** (For clean architecture)
- **Docker** (For containerization)
- **Testing**

## Features
- **GraphQL API**: Query-based data retrieval.
- **Efficient Filtering & Sorting**:
  - Filters materials based on `maxPrice`.
  - Sorts materials by `rating > quality`.
  - Fetches supplier details based on `materialType`, `materialName`, and `location`.
- **Optimized API Calls**:
  - Two goroutines fetch materials and suppliers concurrently.
- **Redis Caching**: Speeds up frequent API requests.

## API Endpoints
### Query: Best Material Search
```graphql
query {
  bestMaterial(request: {
    materialType: "Brick",
    price: 500.0,
    locality: "New York"
  }) {
    material {
      materialName
      materialType
      price
      rating
      Quality
    }
    supplier {
      supplierName
      supplierLocation
      stockAvailability
    }
  }
}
```

## Project Structure
```
├── Config
│   ├── Config.go
│   └── redis.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── gqlgen.yml
├── graph
│   ├── generated.go
│   ├── model
│   │   └── models_gen.go
│   ├── resolver.go
│   ├── schema.graphqls
│   ├── schema.resolvers.go
│   └── validate
│       └── validate.go
├── models
│   └── models.go
├── repository
│   ├── cache.go
│   └── repository.go
├── server.go
├── service
│   └── service.go
└── Testing
    └── server_test.go

```

## How to Run

### Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/SubhamMurarka/GraphQL.git
   cd project-root
   ```
2. Start the project:
   ```sh
    docker-compose up -d --build
   ```
3. Access on localhost:8085
