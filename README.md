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

