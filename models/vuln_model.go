package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
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

func InsecureQuery(category string) ([]Product, error) {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "192.168.92.43:3306",
		DBName: "sqlidb",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS employees(id int AUTO_INCREMENT PRIMARY KEY NOT NULL, name VARCHAR(200) NOT NULL, salary DECIMAL(7,2) NOT NULL)")
	if err != nil {
		return nil, err
	}
	_, err = db.Query("CREATE TABLE IF NOT EXISTS products(id int AUTO_INCREMENT PRIMARY KEY NOT NULL, name VARCHAR(200) NOT NULL, category VARCHAR(50) NOT NULL, price DECIMAL(6,2) NOT NULL, description TEXT)")
	if err != nil {
		return nil, err
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS credentials(id int AUTO_INCREMENT PRIMARY KEY NOT NULL, name VARCHAR(200) NOT NULL, password VARCHAR(200) NOT NULL)")
	if err != nil {
		return nil, err
	}

	queryString := fmt.Sprintf("SELECT * FROM products WHERE category = '%s'", category)
	fmt.Println(queryString)
	rows, queryErr := db.Query(queryString)
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

func some()bool {
	return true
}