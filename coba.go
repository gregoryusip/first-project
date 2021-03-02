package main

// import (
// 	"encoding/json"
// 	"errors"

// 	"github.com/bitwurx/jrpc2"
// )

// type AddParams struct {
// 	X *float64 `json:"x"`
// 	Y *float64 `json:"y"`
// }

// func (ap *AddParams) FromPositional(params []interface{}) error {
// 	if len(params) != 2 {
// 		return errors.New("Exactly two integers are required")
// 	}

// 	x := params[0].(float64)
// 	y := params[1].(float64)

// 	ap.X = &x
// 	ap.Y = &y

// 	return nil
// }

// func Add(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
// 	p := new(AddParams)

// 	if err := jrpc2.ParseParams(params, p); err != nil {
// 		return nil, err
// 	}

// 	if p.X == nil || p.Y == nil {
// 		return nil, &jrpc2.ErrorObject{
// 			Code:    jrpc2.InvalidParamsCode,
// 			Message: jrpc2.InvalidParamsMsg,
// 			Data:    "exactly two integers are required",
// 		}
// 	}

// 	return *p.X + *p.Y, nil
// }

// func main() {
// 	// Hello world, the web server

// 	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
// 	// 	io.WriteString(w, "Hello, world!\n")
// 	// }

// 	// http.HandleFunc("/hello", helloHandler)
// 	// log.Println("Listing for requests at http://localhost:8000/hello")
// 	// log.Fatal(http.ListenAndServe(":8000", nil))

// 	s := jrpc2.NewServer(":8888", "/api/v1/rpc", nil)

// 	s.Register("add", jrpc2.Method{Method: Add})

// 	s.Start()

// }

import (
	"encoding/json"

	"github.com/bitwurx/jrpc2"
)

type AddV1Params struct {
	X int `json:x`
	Y int `json:y`
}

func (p *AddV1Params) FromPositional(params []interface{}) error {
	p.X = int(params[0].(float64))
	p.Y = int(params[1].(float64))
	return nil
}

func AddV1(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	p := new(AddV1Params)
	if err := jrpc2.ParseParams(params, p); err != nil {
		return nil, err
	}
	return p.X + p.Y, nil
}

type AddV2Params struct {
	Args []float64 `json:args`
}

func (p *AddV2Params) FromPositional(params []interface{}) error {
	p.Args = params[0].([]float64)
	return nil
}

func AddV2(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	p := new(AddV2Params)
	if err := jrpc2.ParseParams(params, p); err != nil {
		return nil, err
	}
	return p.Args[0] + p.Args[1], nil
}

func main2() {
	v1 := jrpc2.NewMuxHandler()
	v1.Register("add", jrpc2.Method{Method: AddV1})
	v2 := jrpc2.NewMuxHandler()
	v2.Register("add", jrpc2.Method{Method: AddV2})
	s := jrpc2.NewMuxServer(":8080", nil)
	s.AddHandler("/rpc/v1", v1)
	s.AddHandler("/rpc/v2", v2)
	s.Start()
}
