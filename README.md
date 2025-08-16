# Item Comparison MELI

A tool for comparing items on Mercado Libre (MELI), helping users analyze features, prices, and reviews to make informed purchasing decisions.

## Features

- Search and compare multiple Mercado Libre items
- View detailed product information
- Price and feature comparison
- User-friendly interface

## Getting Started

### Prerequisites

- Docker
- Go (v1.24+ recommended) - for local development

### Installation & Running

#### Using Docker (Recommended)

```bash
git clone https://github.com/edwinhnandez/item-comparison-meli.git
cd item-comparison-meli

# Build and run with Docker
docker buildx build -t item-comparison-api .
docker run -p 8080:8080 item-comparison-api
```

### If you are using colima and see the following error
```bash
ERROR: BuildKit is enabled but the buildx component is missing or broken.
       Install the buildx component to build images with BuildKit:
       https://docs.docker.com/go/buildx/
```
You can run the following command to fix it:
```bash
brew install docker-buildx
```

#### Local Development

```bash
git clone https://github.com/edwinhnandez/item-comparison-meli.git
cd item-comparison-meli
go mod download
go run main.go
```

### API Documentation

Once running, access Swagger documentation at:
- http://localhost:8080/swagger/index.html

### API Endpoints

- `GET /api/products` - Get all products
- `GET /api/products/{id}` - Get product by ID
- `GET /api/products/compare?ids=1,2` - Compare products by IDs

## Usage

1. Start the application using Docker or Go
2. Access the API endpoints directly or via Swagger UI
3. Use the comparison endpoint to analyze multiple products

### Environment Variables

- `GIN_MODE` - Set to `release` for production (default: `debug`)
- `PORT` - Server port (default: `8080`)
- `DATA_FILE` - Path to products JSON file (default: `data/products.json`)

### Testing

```bash
go test -v ./...
```

## License

This project is licensed under the MIT License.

## Acknowledgments

- Open source community


## API Design Diagram Explanation

[Sequence Diagram] (https://mermaid.live/edit#pako:eNptkl1PgzAUhv9Kc640YRuDfdFETXRq9ELN5pXiRUOPswottsU4l_13y8rY_OCCwDnP-563B1aQKY5AweB7hTLDqWALzYpUEneVTFuRiZJJS85ygdL-rc9UZVH_wytptcrz_3pz1B8iw7-NKbPMV_3dD-0cH_splFye35MeK0Wv1IpXmTW9TBXOAU8EN0f9IPI6jzvdLgX1RaLrg5rmILu2Y5tUlCzQ3jX2p8urqTl4dMZPh17SUI6vwzpXZJxs03RfjZKeq7udfdcZ2kpLwvK8xX86_op7IXJ3BuS_6L3Ie3u5ULpg1jpaoymVNPhjEbX1ZpWUXM9vb1oIAlhowYE-s9xgAAU6m_odVrU-BfuCBaZA3SNn-i2FVK6dyH2sB6UKoFZXTqZVtXhpTaqSM7v9kVoEJUd9pippgQ43DkBX8Ak06Xf74yQah3E0iOJ4MApgCTTqd-PxIEmGSTgKh-FkFK8D-NrMDLuT8XD9DY006DI)

### Components
1. Client
- The frontend or API consumer making HTTP requests

2. API Server
- The Go backend application with Gin framework
- Gin Router: Routes incoming requests to controllers
- Product Controller: Handles HTTP logic and responses
- Product Service: Contains business logic
- Product Model: Data structure definitions

3. JSON Data File
Persistent storage (replaces database)

4. Swagger UI
Interactive API documentation

### Request Flow
1. List Products
- Client → GET /api/products → Router → Controller → Service → Model → JSON Data

2. Compare Products
- Client → GET /api/products/compare?ids=1,2 → Router → Controller → Service (filters by IDs) → Model → JSON Data

3. Get Single Product
- Client → GET /api/products/:id → Router → Controller → Service (finds by ID) → Model → JSON Data

4. API Documentation
- Client → GET /swagger → Swagger UI (shows interactive docs)

### Key Design Features
1. Layered Architecture
- Presentation Layer (Controller)
- Business Logic Layer (Service)
- Data Access Layer (Model + JSON file)

2. Separation of Concerns
- Routing, business logic, and data handling are separated
- JSON file acts as a simple database substitute
3. Documentation
Built-in Swagger UI for API exploration
4. Data Flow
- All data originates from the JSON file
- Service layer handles all data operations
- Controller prepares HTTP responses