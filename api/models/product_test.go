package models

import (
	"encoding/json"
	"testing"
)

func TestProduct_JSONMarshaling(t *testing.T) {
	product := Product{
		ID:          "1",
		Name:        "Test Product",
		ImageURL:    "http://example.com/image.jpg",
		Description: "Test description",
		Price:       99.99,
		Rating:      4.5,
		Specifications: map[string]interface{}{
			"color": "red",
			"size":  "large",
		},
	}

	data, err := json.Marshal(product)
	if err != nil {
		t.Fatalf("Failed to marshal product: %v", err)
	}

	var unmarshaled Product
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal product: %v", err)
	}

	if unmarshaled.ID != product.ID {
		t.Errorf("Expected ID %s, got %s", product.ID, unmarshaled.ID)
	}
	if unmarshaled.Name != product.Name {
		t.Errorf("Expected Name %s, got %s", product.Name, unmarshaled.Name)
	}
	if unmarshaled.Price != product.Price {
		t.Errorf("Expected Price %f, got %f", product.Price, unmarshaled.Price)
	}
}

func TestProductResponse_JSONMarshaling(t *testing.T) {
	response := ProductResponse{
		Success: true,
		Message: "Success",
		Data:    "test data",
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal response: %v", err)
	}

	var unmarshaled ProductResponse
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if unmarshaled.Success != response.Success {
		t.Errorf("Expected Success %t, got %t", response.Success, unmarshaled.Success)
	}
}
