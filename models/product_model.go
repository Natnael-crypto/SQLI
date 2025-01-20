package models

import (
	"database/sql"
	"fmt"
	"log"
	"sqli/initializers"
)

type Employee struct {
	ID     int
	Name   string
	salary float64
}

type Product struct {
	ID          int
	Name        string
	Category    string
	Price       float64
	Description string
}

func (p *Product) GenerateViewModel() ProductVM {
	return ProductVM{
		Name:        p.Name,
		Price:       p.Price,
		Description: p.Description,
	}
}

type ProductVM struct {
	Name        string
	Price       float64
	Description string
}


func SecureGetProductsByCategory(category string) ([]Product, error) {
	var (
		err      error
		stmt     *sql.Stmt
		rows     *sql.Rows
		products []Product
	)
	preparedString := fmt.Sprintf("SELECT * FROM products WHERE category=?")
	log.Printf("preparedString: %v\n", preparedString)

	stmt, err = initializers.DB.Prepare(preparedString)
	if err != nil {
		log.Printf("error occured in prepare while trying to get products by category: %v\n", err)
		return nil, SomethingWentWrongErr
	}
	defer stmt.Close()

	rows, err = stmt.Query(category)
	if err != nil {
		log.Printf("error occured in query while trying to get products by category: %v\n", err)
		return nil, SomethingWentWrongErr
	}
	for rows.Next() {
		var product Product
		rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Description)
		products = append(products, product)
	}

	return products, nil
}

func GetAllProducts() ([]Product, error) {
	var (
		err      error
		rows     *sql.Rows
		products []Product
	)

	query := "SELECT * FROM products"
	rows, err = initializers.DB.Query(query)
	if err != nil {
		log.Printf("error occurred in query while trying to get all products: %v\n", err)
		return nil, SomethingWentWrongErr
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Description)
		if err != nil {
			log.Printf("error occurred in scan while iterating through products: %v\n", err)
			return nil, SomethingWentWrongErr
		}
		products = append(products, product)
	}

	return products, nil
}
