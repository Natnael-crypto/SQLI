package models

import (
	"database/sql"
	"log"
	"sqli/initializers"
)

// Employee represents an employee record
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Product represents a product record
type Product struct {
	ID          int
	Name        string
	Category    string
	Price       float64
	Description string
}

// ProductVM represents a view model for Product
type ProductVM struct {
	Name        string
	Price       float64
	Description string
}

// GenerateViewModel converts a Product to ProductVM
func (p *Product) GenerateViewModel() ProductVM {
	return ProductVM{
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
	}
}

// SecureGetProductsByCategory retrieves products by category using a prepared statement
func SecureGetProductsByCategory(category string) ([]Product, error) {
	var (
		err      error
		stmt     *sql.Stmt
		rows     *sql.Rows
		products []Product
	)

	preparedString := "SELECT * FROM products WHERE category=?"
	log.Printf("Prepared query: %v\n", preparedString)

	stmt, err = initializers.DB.Prepare(preparedString)
	if err != nil {
		log.Printf("Error in prepare while trying to get products by category: %v\n", err)
		return nil, SomethingWentWrongErr
	}
	defer stmt.Close()

	rows, err = stmt.Query(category)
	if err != nil {
		log.Printf("Error in query while trying to get products by category: %v\n", err)
		return nil, SomethingWentWrongErr
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Description)
		if err != nil {
			log.Printf("Error in scan while iterating through products: %v\n", err)
			return nil, SomethingWentWrongErr
		}
		products = append(products, product)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v\n", err)
		return nil, SomethingWentWrongErr
	}

	return products, nil
}

// GetAllProducts retrieves all products from the database
func GetAllProducts() ([]Product, error) {
	var (
		err      error
		rows     *sql.Rows
		products []Product
	)

	query := "SELECT * FROM products"
	rows, err = initializers.DB.Query(query)
	if err != nil {
		log.Printf("Error in query while trying to get all products: %v\n", err)
		return nil, SomethingWentWrongErr
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Description)
		if err != nil {
			log.Printf("Error in scan while iterating through products: %v\n", err)
			return nil, SomethingWentWrongErr
		}
		products = append(products, product)
	}

	// Check for errors during iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v\n", err)
		return nil, SomethingWentWrongErr
	}

	return products, nil
}
