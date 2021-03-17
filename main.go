package main

import (
	"github.com/bitwurx/jrpc2"
	"github.com/gregoryusip/first-project/config"
	"github.com/gregoryusip/first-project/controller"
	"github.com/gregoryusip/first-project/models"
	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "mysecret"
// 	dbname   = "first_project_db"
// )

func main() {
	// DATABASE CONNECTION
	db := config.CreateConnection()

	// psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err := sql.Open("postgres", psqlConn)

	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Sucessfully Connected!")

	// SINGLETON
	productORM := models.NewProductModel(models.Dependencies{
		Db: db,
	})

	productController := controller.NewProductController(controller.Dependencies{
		ProductORM: productORM,
	})

	server := jrpc2.NewMuxHandler()
	server.Register("product.CreateProduct", jrpc2.Method{Method: productController.AddProduct})
	server.Register("product.ReadProduct", jrpc2.Method{Method: productController.ReadedProduct})

	s := jrpc2.NewMuxServer(":8080", nil)
	s.AddHandler("/rpc/create", server)
	s.AddHandler("/rpc/read", server)

	s.Start()
}
