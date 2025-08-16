# How GenAI and Modern Tools Were Used

1. Code Generation: Used GitHub Copilot for boilerplate code generation and suggestions.
2. Error Handling Patterns: Consulted ChatGPT for best practices in Go error handling.
3. Documentation: Used swaggo to automatically generate Swagger docs from code annotations.

# AI Tool Prompts Used in Development

## 1. Project Setup Prompts

**Prompt:**  
"Generate a Go project structure for a RESTful API that compares product items. Include models, services, and controllers separation. The API should use Gin framework and have Swagger documentation."

**Prompt:**  
"Create a sample products.json file with 3 example products for comparison. Include fields: id, name, imageUrl, description, price, rating, and specifications (which should be a nested object with technical details)."

## 2. Code Generation Prompts

**Prompt:**  
"Write a Go service layer for product comparison that loads data from a JSON file and provides methods to:  
1. Get all products  
2. Get products by IDs (for comparison)  
3. Get single product by ID  
Include proper error handling and thread safety."

**Prompt:**  
"Generate a Gin controller for product comparison with these endpoints:  
- GET /products - list all  
- GET /products/compare?ids=1,2 - compare specific products  
- GET /products/{id} - get single product  
Use consistent JSON responses with success status, message, and data fields."

## 3. Error Handling Prompts

**Prompt:**  
"Show best practices for error handling in Go REST APIs using Gin. Include examples for:  
- Not found cases  
- Bad requests  
- Internal server errors  
- Consistent error response format"

**Prompt:**  
"Improve this Go error handling to include more context and follow best practices:  
`if err != nil { return err }`"

## 4. Documentation Prompts

**Prompt:**  
"Create a comprehensive README.md for a Go product comparison API that includes:  
- Project description  
- API endpoints  
- Setup instructions  
- Example requests/responses  
- Technology stack"

## 5. Deployment Prompts

**Prompt:**  
"Create a Dockerfile for a Go Gin application that:  
1. Uses multi-stage builds  
2. Has proper health checks  
3. Configures through environment variables  
4. Is secure (non-root user, etc.)"

**Prompt:**  
"Generate GitHub Actions workflow for a Go API that:  
1. Runs tests on push  
2. Builds a Docker image  
3. Deploys to AWS ECS  
Include all necessary steps and secrets handling."

## AI Tool Usage Notes

1. All generated code was reviewed and modified as needed
2. Business logic was validated against requirements
3. Architectural decisions were made consciously, not just accepted from AI
4. Security considerations were added beyond initial suggestions
5. The prompts evolved through iteration to get better results