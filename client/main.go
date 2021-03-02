package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Name  string
	Price int
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection Error: ", err)
	}

	a := Item{"Baju", 90000}
	b := Item{"Celana", 50000}
	c := Item{"Kemeja", 140000}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("Database: ", db)

	client.Call("API.EditItem", Item{"Celana", 40000}, &reply)
	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &reply)
	fmt.Println("Database: ", db)

	client.Call("API.GetByName", "Baju", &reply)
	fmt.Println("First Item: ", reply)
}
