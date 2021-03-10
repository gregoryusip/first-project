// https://blog.afrizalmy.com/read/membuat-crud-golang-rest-api-dengan-postgresql/

package models

import (
	"fmt"
	"log"

	"github.com/gregoryusip/first-project/config"
	_ "github.com/lib/pq"
)

type Product struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
}

type EditProduct struct {
	Name        string `json:"name"`
	NewName     string `json:"newname"`
	NewPrice    int    `json:"price"`
	NewQuantity int    `json:"quantity"`
}

// type Products int

func CreateProduct(product Product) int64 {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO product (name, price, quantity) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, product.Name, product.Price, product.Quantity).Scan(&id)

	if err != nil {
		log.Fatalf("Can't exec the Query. %v", err)
	}

	fmt.Println("Insert data single record %v", id)

	return id
}

func UpdateProduct(id int64, product Product) int64 {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `UPDATE product SET name=$2, price=$1000, quantity=$5 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, product.Name, product.Price, product.Quantity)

	if err != nil {
		log.Fatalf("Can't exec the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error when checking data rows to update. %v", err)
	}

	fmt.Printf("Total rows/record to update %v\n", rowsAffected)

	return rowsAffected
}
