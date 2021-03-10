package main

import (
	"database/sql"

	"github.com/bitwurx/jrpc2"
	"github.com/gregoryusip/first-project/controller"
	_ "github.com/lib/pq"
)

func main() {
	// DATABASE CONNECTION
	connection := "user=postgres dbname=connect-db password=mysecret host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// OBJECT
	// products := new(models.Products)

	// HANDLER
	v1 := jrpc2.NewMuxHandler()
	// v1.Register("add", jrpc2.Method{Method: TotalProduct})
	v1.Register("addProduct", jrpc2.Method{Method: controller.AddProduct})
	v2 := jrpc2.NewMuxHandler()
	v2.Register("add", jrpc2.Method{Method: AddV2})
	s := jrpc2.NewMuxServer(":8080", nil)
	s.AddHandler("/rpc/v1", v1)
	s.AddHandler("/rpc/v2", v2)
	s.Start()
}
