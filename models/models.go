// https://blog.afrizalmy.com/read/membuat-crud-golang-rest-api-dengan-postgresql/

package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-pg/pg"
	"github.com/gregoryusip/first-project/config/database"
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
	CreateProduct(produk Products) int
	UpdateProduct(name string, produk Products) int64
	DeleteProduct(name string) int64
	ReadProductPG() ([]Products, error)
	CreateProductPG(produk Products) ([]Products, error)
	UpdateProductPG(name string, produk Products) ([]Products, error)
	DeleteProductPG(name string) error
}

type Dependencies struct {
	Db   *sql.DB
	DbPg *pg.DB
}

func NewProductModel(deps Dependencies) ProductModel {
	return &ProductRepository{
		DbPg: deps.DbPg,
		// Db: deps.Db,
	}
}

// func NewProductModel(Db *sql.DB) ProductModel {
// 	return ProductRepository{Db: Db}
// }

type ProductRepository struct {
	Db   *sql.DB
	DbPg *pg.DB
	Db2  database.Ormer
}

func (p *ProductRepository) CreateProductPG(produk Products) ([]Products, error) {

	// sqlStatement := `INSERT INTO products (name, price, quantity) VALUES ($1, $2, $3) RETURNING id`

	// var id int

	// err := p.Db.QueryRow(sqlStatement, produk.Name, produk.Price, produk.Quantity).Scan(&id)
	// _, err := p.Db2.Model(&produk).TableExpr("products").Insert()
	_, err := p.Db2.Model(&produk).TableExpr("products").Insert()

	if err != nil {
		return nil, err
	}

	return p.ReadProductPG()
}

func (p *ProductRepository) CreateProduct(produk Products) int {
	// p.Db.Query("INSERT INTO product")
	sqlStatement := `INSERT INTO products (name, price, quantity) VALUES ($1, $2, $3) RETURNING id`

	var id int

	err := p.Db.QueryRow(sqlStatement, produk.Name, produk.Price, produk.Quantity).Scan(&id)

	if err != nil {
		log.Fatalf("Can't execute the Query. %v", err)
	}

	// fmt.Printf("Insert data single record %d", id)

	return id
}

func (p *ProductRepository) ReadProductPG() ([]Products, error) {
	var products []Products

	// sqlStatement := fmt.Sprintf(`
	// 	SELECT
	// 		*
	// 	FROM
	// 		products
	// `)

	// _, err := p.Db2.QueryContext(ctx, &products, sqlStatement)
	err := p.DbPg.Model(&products).Select()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) ReadProduct() ([]Products, error) {

	var products []Products

	sqlStatement := `SELECT * FROM products`

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

func (p *ProductRepository) UpdateProductPG(name string, produk Products) ([]Products, error) {

	// sqlStatement := fmt.Sprintf(`UPDATE products SET name='%s', price='%d', quantity='%d' WHERE name='%s'`, produk.Name, produk.Price, produk.Quantity, produk.Name)

	// _, err := p.Db2.ExecContext(ctx, sqlStatement)
	_, err := p.Db2.Model(&produk).TableExpr("products").Where("name = ?", name).Update()

	if err != nil {
		return nil, err
	}

	return p.ReadProductPG()
}

func (p *ProductRepository) UpdateProduct(name string, produk Products) int64 {

	sqlStatement := `UPDATE products SET name=$1, price=$2, quantity=$3 WHERE name=$1`

	res, err := p.Db.Exec(sqlStatement, name, produk.Price, produk.Quantity)

	if err != nil {
		log.Fatalf("Can't execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error when checking data rows to update. %v", err)
	}

	fmt.Printf("Total rows/record to update %v\n", rowsAffected)

	return rowsAffected
}

func (p *ProductRepository) DeleteProductPG(name string) error {
	var products []Products

	// sqlStatement := fmt.Sprintf(`DELETE FROM products WHERE name = '%s'`, name)

	// _, err := p.Db2.ExecContext(ctx, sqlStatement)
	_, err := p.Db2.Model(&products).Where("name = ?", name).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) DeleteProduct(name string) int64 {

	sqlStatement := `DELETE FROM products WHERE name=$1`

	res, err := p.Db.Exec(sqlStatement, name)

	if err != nil {
		log.Fatalf("Can't execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Can't find the data. %v", err)
	}

	fmt.Printf("Total rows/record that deleted. %v", rowsAffected)

	return rowsAffected
}

type EditProduct struct {
	Name        string `json:"name"`
	NewName     string `json:"newname"`
	NewPrice    int    `json:"price"`
	NewQuantity int    `json:"quantity"`
}
