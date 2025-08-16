package controllers

import (
	"net/http"
	"strings"

	"item-comparison-meli.com/m/api/models"
	"item-comparison-meli.com/m/api/services"

	"github.com/gin-gonic/gin"
)

// ProductController handles HTTP requests related to products
type ProductController struct {
	productService *services.ProductService
}

// NewProductController creates a new ProductController instance
func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// @Summary Get all products
// @Description Retrieve a list of all products
// @Tags Products
// @Produce json
// @Success 200 {object} models.ProductResponse
// @Failure 404 {object} models.ProductResponse
// @Router /products [get]
func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, models.ProductResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ProductResponse{
		Success: true,
		Data:    products,
	})
}

// @Summary Compare products
// @Description Compare products by their IDs
// @Tags Products
// @Produce json
// @Param ids query string true "Comma-separated list of product IDs"
// @Success 200 {object} models.ProductResponse
// @Failure 400 {object} models.ProductResponse
// @Failure 404 {object} models.ProductResponse
// @Router /products/compare [get]
func (pc *ProductController) CompareProducts(c *gin.Context) {
	idsParam := c.Query("ids")
	if idsParam == "" {
		c.JSON(http.StatusBadRequest, models.ProductResponse{
			Success: false,
			Message: "ids parameter is required",
		})
		return
	}

	ids := strings.Split(idsParam, ",")
	if len(ids) == 0 {
		c.JSON(http.StatusBadRequest, models.ProductResponse{
			Success: false,
			Message: "at least one product ID must be provided",
		})
		return
	}

	products, err := pc.productService.GetProductsByIDs(ids)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ProductResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ProductResponse{
		Success: true,
		Data:    products,
	})
}

// @Summary Get product by ID
// @Description Retrieve a single product by its ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.ProductResponse
// @Failure 400 {object} models.ProductResponse
// @Failure 404 {object} models.ProductResponse
// @Router /products/{id} [get]
func (pc *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, models.ProductResponse{
			Success: false,
			Message: "product ID is required",
		})
		return
	}

	product, err := pc.productService.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ProductResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ProductResponse{
		Success: true,
		Data:    product,
	})
}
