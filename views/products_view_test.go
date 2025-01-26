package views

import (
	"bytes"
	"html/template"
	"sqli/initializers"
	"sqli/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsRender_Success(t *testing.T) {
	var buf bytes.Buffer

	// Mock the HTML template for products rendering
	initializers.Template = template.Must(template.New("products.html").Parse("User: {{.Username}}, Products Count: {{len .Products}}"))

	userVM := UserVM{
		Username: "john_doe",
		Products: []models.ProductVM{
			{Name: "Laptop", Price: 1500.00, Description: "Gaming laptop"},
			{Name: "Headphones", Price: 200.00, Description: "Noise-canceling headphones"},
		},
	}

	ProductsRender(&buf, userVM)

	expectedOutput := "User: john_doe, Products Count: 2"
	assert.Equal(t, expectedOutput, buf.String())
}

func TestProductsRender_TemplateExecutionFailure(t *testing.T) {
	var buf bytes.Buffer

	// Mock a faulty template that references an invalid field to trigger an error
	initializers.Template = template.Must(template.New("products.html").Parse("{{.InvalidField}}"))

	userVM := UserVM{
		Username: "john_doe",
		Products: []models.ProductVM{},
	}

	ProductsRender(&buf, userVM)

	// Expect an empty output since the template execution should fail
	assert.Empty(t, buf.String())
}

func TestProductsRender_EmptyProducts(t *testing.T) {
	var buf bytes.Buffer

	// Mock template handling empty product lists
	initializers.Template = template.Must(template.New("products.html").Parse("User: {{.Username}}, Products Count: {{len .Products}}"))

	userVM := UserVM{
		Username: "john_doe",
		Products: []models.ProductVM{}, // No products
	}

	ProductsRender(&buf, userVM)

	expectedOutput := "User: john_doe, Products Count: 0"
	assert.Equal(t, expectedOutput, buf.String())
}