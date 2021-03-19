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
	db := config.CreateConnection()

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
