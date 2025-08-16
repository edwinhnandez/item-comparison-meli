package services

import (
	"os"
	"testing"
)

func TestNewProductService(t *testing.T) {
	// Create temporary test data file
	testData := `[
		{
			"id": "1",
			"name": "Test Product",
			"imageUrl": "http://example.com/image.jpg",
			"description": "Test description",
			"price": 99.99,
			"rating": 4.5,
			"specifications": {"color": "red"}
		}
	]`

	tmpFile, err := os.CreateTemp("", "test_products_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(testData); err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}
	tmpFile.Close()

	service, err := NewProductService(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to create service: %v", err)
	}

	products, err := service.GetAllProducts()
	if err != nil {
		t.Fatalf("Failed to get products: %v", err)
	}

	if len(products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(products))
	}

	if products[0].ID != "1" {
		t.Errorf("Expected product ID '1', got '%s'", products[0].ID)
	}
}

func TestProductService_GetProductByID(t *testing.T) {
	testData := `[{"id": "1", "name": "Test Product", "price": 99.99}]`

	tmpFile, err := os.CreateTemp("", "test_products_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	tmpFile.WriteString(testData)
	tmpFile.Close()

	service, _ := NewProductService(tmpFile.Name())

	product, err := service.GetProductByID("1")
	if err != nil {
		t.Fatalf("Failed to get product: %v", err)
	}

	if product.ID != "1" {
		t.Errorf("Expected product ID '1', got '%s'", product.ID)
	}

	_, err = service.GetProductByID("999")
	if err == nil {
		t.Error("Expected error for non-existent product")
	}
}
