package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"item-comparison-meli.com/m/api/services"

	"github.com/gin-gonic/gin"
)

func setupTestController(t *testing.T) *ProductController {
	testData := `[
		{"id": "1", "name": "Product 1", "price": 99.99},
		{"id": "2", "name": "Product 2", "price": 149.99}
	]`

	tmpFile, err := os.CreateTemp("", "test_products_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	t.Cleanup(func() { os.Remove(tmpFile.Name()) })

	tmpFile.WriteString(testData)
	tmpFile.Close()

	service, err := services.NewProductService(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to create service: %v", err)
	}

	return NewProductController(service)
}

func TestProductController_GetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := setupTestController(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/products", nil)

	controller.GetProducts(c)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestProductController_GetProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := setupTestController(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/products/1", nil)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	controller.GetProduct(c)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestProductController_CompareProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	controller := setupTestController(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/products/compare?ids=1,2", nil)

	controller.CompareProducts(c)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}
