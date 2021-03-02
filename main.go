package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Name  string
	Price int
}

type API int

var database []Item

func (a *API) GetDB(name string, reply *[]Item) error {
	*reply = database

	return nil
}

func (a *API) GetByName(name string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Name == name {
			getItem = val
		}
	}

	*reply = getItem

	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item

	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.Name == edit.Name {
			database[idx] = Item{edit.Name, edit.Price}
			changed = edit
		}
	}
	*reply = changed

	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.Name == item.Name && val.Price == item.Price {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	*reply = del

	return nil
}

func main() {
	// Hello world, the web server

	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }

	// http.HandleFunc("/hello", helloHandler)
	// log.Println("Listing for requests at http://localhost:8000/hello")
	// log.Fatal(http.ListenAndServe(":8000", nil))

	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener Error", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}

	// fmt.Println("Initial Database: ", database)

	// a := Item{"Baju", 90000}
	// b := Item{"Celana", 50000}
	// c := Item{"Kemeja", 140000}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)

	// fmt.Println("Second Database: ", database)

	// DeleteItem(b)
	// fmt.Println("Third Database: ", database)

	// EditItem("Baju", Item{"T-shirt", 70000})
	// fmt.Println("Fourth Database: ", database)

	// x := GetByName("T-shirt")
	// y := GetByName("Kemeja")
	// fmt.Println(x, y)
}
