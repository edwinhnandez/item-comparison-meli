package models

// Product represents a product with comparison attributes
// swagger:model Product
type Product struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	ImageURL       string                 `json:"imageUrl"`
	Description    string                 `json:"description"`
	Price          float64                `json:"price"`
	Rating         float64                `json:"rating"`
	Specifications map[string]interface{} `json:"specifications"`
}

// ProductResponse standard API response format
// swagger:model ProductResponse
// swagger:response ProductResponse
type ProductResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
