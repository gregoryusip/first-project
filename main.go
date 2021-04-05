package main

import (
	"github.com/bitwurx/jrpc2"
	"github.com/gregoryusip/first-project/config"
	"github.com/gregoryusip/first-project/controller"
	"github.com/gregoryusip/first-project/models"
	_ "github.com/lib/pq"
)

func main() {
	// DATABASE CONNECTION
	// db := config.CreateConnection()
	db := config.CreateConnection("../first-project")
	// db := config.CreateConnectionPG()

	// config, err := config.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("Cannot load config:", err)
	// }

	// conn, err := sql.Open(config.DBDriver, config.DBSource)
	// if err != nil {
	// 	log.Fatal("Cannot connect to db:", err)
	// }

	// store := db.NewStore(conn)

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
	server.Register("product.UpdateProduct", jrpc2.Method{Method: productController.UpdatedProduct})
	server.Register("product.DeleteProduct", jrpc2.Method{Method: productController.DeletedProduct})

	s := jrpc2.NewMuxServer(":8080", nil)
	s.AddHandler("/rpc/create", server)
	s.AddHandler("/rpc/read", server)
	s.AddHandler("/rpc/update", server)
	s.AddHandler("/rpc/delete", server)

	s.Start()
}
