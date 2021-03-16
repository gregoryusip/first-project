// https://blog.afrizalmy.com/read/membuat-crud-golang-rest-api-dengan-postgresql/

package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Products struct {
	ID       int    `json:"id" pg:"id,pk"`
	Name     string `json:"name" pg:"name"`
	Price    int    `json:"price" pg:"price"`
	Quantity int    `json:"quantity" pg:"quantity"`
}

func NewProducts(Name string, Price int, Quantity int) Products {
	return Products{0, Name, Price, Quantity}
}

type ProductModel interface {
	ReadProduct() ([]Products, error)
	CreateProduct(produk Products) error
}

type Dependencies struct {
	Db *sql.DB
}

func NewProductModel(deps Dependencies) ProductModel {
	return &ProductRepository{
		Db: deps.Db,
	}
}

// func NewProductModel(Db *sql.DB) ProductModel {
// 	return ProductRepository{Db: Db}
// }

type ProductRepository struct {
	Db *sql.DB
}

func (p *ProductRepository) CreateProduct(produk Products) error {
	// p.Db.Query("INSERT INTO first-project")
	sqlStatement := `INSERT INTO product (name, price, quantity) VALUES ($1, $2, $3) RETURNING id`

	var id int

	err := p.Db.QueryRow(sqlStatement, produk.Name, produk.Price, produk.Quantity).Scan(&id)

	if err != nil {
		log.Fatalf("Can't execute the Query. %v", err)
	}

	fmt.Println("Insert data single record %v", id)

	return nil
}

func (p *ProductRepository) ReadProduct() ([]Products, error) {

	var products []Products

	sqlStatement := `SELECT * FROM product`

	rows, err := p.Db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Can't execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var product Products

		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)

		if err != nil {
			log.Fatalf("Can't take the data. %v", err)
		}

		products = append(products, product)
	}

	return products, err
}

type EditProduct struct {
	Name        string `json:"name"`
	NewName     string `json:"newname"`
	NewPrice    int    `json:"price"`
	NewQuantity int    `json:"quantity"`
}
