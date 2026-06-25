package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Product struct provides the data model for productstore
type Product struct {
	ID          int
	Title       string
	Description string
	Price       float32
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/productstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	product := Product{
		Title:       "Amazon Prime",
		Description: "Amazon Prime - Online",
		Price:       23.99,
	}
	// Insert a product
	createProduct(product)
	// Read all product records
	getProducts()
	// Read a product by given id
	getProductByID(1)
}

// createProduct inserts product values into product table
func createProduct(prd Product) {
	result, err := db.Exec("INSERT INTO products(title, description, price) VALUES($1, $2, $3)", prd.Title, prd.Description, prd.Price)
	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, err := result.LastInsertId()
	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Product with id=%d created successfully (%d row affected)\n", lastInsertID, rowsAffected)
}

// getProducts reads all records from the product table
func getProducts() {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No Records Found")
			return
		}
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var products []*Product
	for rows.Next() {
		prd := &Product{}
		err := rows.Scan(&prd.Title, &prd.Description, &prd.Price, &prd.ID)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prd)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, pr := range products {
		fmt.Printf("%d, %s, %s, $%.2f\n", pr.ID, pr.Title, pr.Description, pr.Price)
	}
}

func getProductByID(id int) {
	var product string
	err := db.QueryRow("SELECT title FROM products WHERE id=$1", id).Scan(&product)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Printf("No product with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Product is %s\n", product)
	}
}
