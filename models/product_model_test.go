package models

import (
	"sqli/initializers"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGenerateViewModel(t *testing.T) {
	product := Product{
		ID:          1,
		Name:        "Laptop",
		Category:    "Electronics",
		Price:       1200.00,
		Description: "High-end gaming laptop",
	}

	vm := product.GenerateViewModel()

	assert.Equal(t, product.Name, vm.Name)
	assert.Equal(t, product.Price, vm.Price)
	assert.Equal(t, product.Description, vm.Description)
}

func TestSecureGetProductsByCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	rows := sqlmock.NewRows([]string{"id", "name", "category", "price", "description"}).
		AddRow(1, "Laptop", "Electronics", 1200.00, "High-end gaming laptop").
		AddRow(2, "Mouse", "Electronics", 25.00, "Wireless mouse")

	mock.ExpectPrepare("SELECT \\* FROM products WHERE category=\\?").
		ExpectQuery().
		WithArgs("Electronics").
		WillReturnRows(rows)

	products, err := SecureGetProductsByCategory("Electronics")

	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Laptop", products[0].Name)
	assert.Equal(t, "Mouse", products[1].Name)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSecureGetProductsByCategory_ErrorInQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectPrepare("SELECT \\* FROM products WHERE category=\\?").
		ExpectQuery().
		WithArgs("Electronics").
		WillReturnError(SomethingWentWrongErr)

	products, err := SecureGetProductsByCategory("Electronics")

	assert.Error(t, err)
	assert.Nil(t, products)
	assert.Equal(t, SomethingWentWrongErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	rows := sqlmock.NewRows([]string{"id", "name", "category", "price", "description"}).
		AddRow(1, "Laptop", "Electronics", 1200.00, "High-end gaming laptop").
		AddRow(2, "Mouse", "Electronics", 25.00, "Wireless mouse")

	mock.ExpectQuery("SELECT \\* FROM products").
		WillReturnRows(rows)

	products, err := GetAllProducts()

	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Laptop", products[0].Name)
	assert.Equal(t, "Mouse", products[1].Name)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllProducts_ErrorInQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	initializers.DB = db

	mock.ExpectQuery("SELECT \\* FROM products").
		WillReturnError(SomethingWentWrongErr)

	products, err := GetAllProducts()

	assert.Error(t, err)
	assert.Nil(t, products)
	assert.Equal(t, SomethingWentWrongErr, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
