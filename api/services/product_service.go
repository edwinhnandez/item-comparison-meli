package services

import (
	"encoding/json"
	"errors"
	"io"
	"item-comparison-meli.com/m/api/models"
	"os"
	"sync"
)

// ProductService handles product-related business logic
type ProductService struct {
	products []models.Product
	mu       sync.Mutex
}

// NewProductService creates a new ProductService instance
func NewProductService(dataFile string) (*ProductService, error) {
	ps := &ProductService{}
	if err := ps.loadProducts(dataFile); err != nil {
		return nil, err
	}
	return ps, nil
}

// loadProducts reads products from JSON file
func (s *ProductService) loadProducts(dataFile string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Open(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &s.products); err != nil {
		return err
	}

	return nil
}

// GetAllProducts returns all products
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	if len(s.products) == 0 {
		return nil, errors.New("no products available")
	}
	return s.products, nil
}

// GetProductsByIDs returns products matching the provided IDs
func (s *ProductService) GetProductsByIDs(ids []string) ([]models.Product, error) {
	var result []models.Product
	foundIDs := make(map[string]bool)

	for _, id := range ids {
		for _, product := range s.products {
			if product.ID == id {
				if !foundIDs[id] {
					result = append(result, product)
					foundIDs[id] = true
				}
				break
			}
		}
	}

	if len(result) == 0 {
		return nil, errors.New("no products found with the provided IDs")
	}

	return result, nil
}

// GetProductByID returns a single product by ID
func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	for _, product := range s.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}
