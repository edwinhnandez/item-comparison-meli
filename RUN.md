# How to Run the Item Comparison API

## Prerequisites
- Go 1.16 or higher installed
- Git (optional)

## Installation

1. Clone the repository (or create the files manually):
   ```bash
   git clone https://github.com/edwinhnandezc/item-comparison-meli.git
   cd item-comparison-meli

2. Install dependencies
    ```bash
    go mod tidy
    ```

## Run the API 
1. Start the server: 
    ```bash
    go run main.go
    ```
2. The API will be available at http://localhost:8080

## API Endpoints 

- `GET /api/products`: List all products
  - Example: `http://localhost:8080/api/products` 
- `GET /api/products/compare?ids=1,2` : Compare specific products
    - Example: `http://localhost:8080/api/products/compare?ids=1,2`
- `GET /api/products/:id`: Get a single product details
    - Example: `http://localhost:8080/api/products/1`
- `GET /swagger/index.html` - API documentation (Swagger UI)

## API Testing

1. Run Unit tests
    ```bash
    go test ./...
    ```
## Environment Variables
- `PORT`: Port number for the server (default: 8080)
- `DATA_FILE`: Path to products JSON file (default:"data/products.json")

