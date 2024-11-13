package models

import (
	"fmt"
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

func (p *Product) GenerateViewModel() ProductVM{
	return ProductVM{
		Name: p.Name,
		Price: p.Price,
		Description: p.Description,
	}
}

type ProductVM struct {
	Name        string
	Price       float64
	Description string
}

func VulnGetProductsByCategory(category string) ([]Product, error) {
	queryString := fmt.Sprintf("SELECT * FROM products WHERE category = '%s'", category)
	fmt.Println(queryString)

	rows, queryErr := initializers.DB.Query(queryString)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Description)
		products = append(products, product)
	}
	for _, product := range products {

		fmt.Printf("%#v\n", product)
	}
	return products, nil

}
