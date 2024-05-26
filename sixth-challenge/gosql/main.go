package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "gosql"
)

var (
	db  *sql.DB
	err error
)

type Product struct {
	ID         int
	Product    string
	Created_at string
	Updated_at string
}

type Variant struct {
	ID         int
	Variant    string
	Quantity   int
	Product_id int
	Created_at string
	Updated_at string
}

func main() {
	mysqlInfo := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	db, err = sql.Open("mysql", mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	createProduct()
	updateProduct()
	getProductById()
	createVariant()
	updateVariantById()
	deleteVariantById()
	getProductWithVariant()
}

func createProduct() {
	sqlStatement := `INSERT INTO products (product, created_at, updated_at) VALUES (?, ?, ?)`
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := createdAt

	result, err := db.Exec(sqlStatement, "celana", createdAt, updatedAt)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Printf("New product ID: %d\n", id)
}

func updateProduct() {
	sqlStatement := `UPDATE products SET product = ?, updated_at = ? WHERE id = ?`
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	result, err := db.Exec(sqlStatement, "baju", updated_at, 1)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount (product): ", count)
}

func getProductById() {
	var product Product

	sqlStatement := `SELECT id, product, created_at, updated_at FROM products WHERE id = ?`
	row := db.QueryRow(sqlStatement, 1)
	err := row.Scan(&product.ID, &product.Product, &product.Created_at, &product.Updated_at)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No product found with the given ID")
		} else {
			panic(err)
		}
		return
	}

	fmt.Printf("ID: %d\nProduct: %s\nCreated At: %s\nUpdated At: %s\n", product.ID, product.Product, product.Created_at, product.Updated_at)
}

func createVariant() {
	sqlStatement := `INSERT INTO variants (variant, quantity, product_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	updatedAt := createdAt

	result, err := db.Exec(sqlStatement, "kemeja", 13, 1, createdAt, updatedAt)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Printf("New variant ID: %d\n", id)
}

func updateVariantById() {
	sqlStatement := `UPDATE variants SET variant = ?, quantity = ?, product_id = ?, updated_at = ? WHERE id = ?`
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	result, err := db.Exec(sqlStatement, "crewneck", 19, 1, updated_at, 6)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount (variant): ", count)
}

func deleteVariantById() {
	sqlStatement := `DELETE from variants WHERE id = ?`

	result, err := db.Exec(sqlStatement, 9)
	if err != nil {
		panic(err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted variant amount:", count)
}

func getProductWithVariant() {
	sqlStatement := `SELECT p.product, v.variant FROM products p LEFT JOIN variants v ON p.id = v.product_id`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product string
		var variant sql.NullString

		err := rows.Scan(&product, &variant)
		if err != nil {
			panic(err)
		}

		if variant.Valid {
			fmt.Printf("Product: %s, Variant: %s\n", product, variant.String)
		} else {
			fmt.Printf("Product: %s, Variant: NULL\n", product)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
