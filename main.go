package main

import (
	"github.com/bitwurx/jrpc2"
	"github.com/gregoryusip/first-project/config"
	"github.com/gregoryusip/first-project/controller"
	_ "github.com/lib/pq"
)

func main() {
	// DATABASE CONNECTION
	db := config.CreateConnection()

	// connection := "user=postgres dbname=connect-db password=mysecret host=localhost sslmode=disable"
	// db, err := sql.Open("postgres", connection)

	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// db := postgres.New(&pg.Options{
	// 	User:"postgres",
	// 	Password:"postgres",
	// 	Database:"first-project"
	// })

	// defer db.Close()

	// db.AddQueryHook(postgres.DBLogger{})

	// OBJECT
	// products := new(models.Products)

	// HANDLER
	// v1.Register("add", jrpc2.Method{Method: TotalProduct})
	v1 := jrpc2.NewMuxHandler()
	v1.Register("addProduct", jrpc2.Method{Method: controller.AddProduct})

	v2 := jrpc2.NewMuxHandler()
	v2.Register("updateProduct", jrpc2.Method{Method: controller.UpdatedProduct})

	v3 := jrpc2.NewMuxHandler()
	v3.Register("deleteProduct", jrpc2.Method{Method: controller.De})

	v4 := jrpc2.NewMuxHandler()
	v4.Register("readProduct", jrpc2.Method{Method: controller.ReadedProduct})

	v5 := jrpc2.NewMuxHandler()
	v5.Register("welcome", jrpc2.Method{Method: controller.Welcome})

	s := jrpc2.NewMuxServer(":8080", nil)
	s.AddHandler("/rpc/v1", v1)
	s.AddHandler("/rpc/v2", v2)
	s.AddHandler("/rpc/v3", v3)
	s.AddHandler("/rpc/v4", v4)
	s.AddHandler("/rpc/v5", v5)
	s.Start()
}
