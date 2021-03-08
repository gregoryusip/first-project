// LINK : https://codesource.io/build-a-crud-application-in-golang-with-postgresql/ & https://bambang-sso.medium.com/building-crud-api-using-golang-and-postgresql-469c3352f774

package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gregoryusip/first-project/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected!")

	return db

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := insertProduct(product)

	res := response{
		ID:      insertID,
		Message: "Product created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	user, err := getProduct(int64(id))

	if err != nil {
		log.Fatalf("Unable to get usr. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	users, err := getAllProduct()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	json.NewEncoder(w).Encode(users)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := updateProduct(int64(id), product)

	msg := fmt.Sprintf("Product updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	deletedRows := deleteProduct(int64(id))

	msg := fmt.Sprintf("Product deleted successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// 	HANDLER FUNCTION
func insertProduct(product models.Product) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO products (name, price, quantity) VALUES ($1, $2, $3) RETURNING productid`

	var id int64

	err := db.QueryRow(sqlStatement, product.Name, product.Price, product.Quantity).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func getProduct(id int64) (models.Product, error) {
	db := createConnection()

	defer db.Close()

	var product models.Product

	sqlStatement := `SELECT * FROM products WHERE productid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return product, err
}

func getAllProduct() ([]models.Product, error) {
	db := createConnection()

	defer db.Close()

	var products []models.Product

	sqlStatement := `SELECT * FROM products`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		products = append(products, product)
	}

	return products, err
}

func updateProduct(id int64, product models.Product) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `UPDATE products SET name=$2, price=$3, quantity=$6 WHERE productid=$1`

	res, err := db.Exec(sqlStatement, id, product.Name, product.Price, product.Quantity)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func deleteProduct(id int64) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM products WHERE productid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
