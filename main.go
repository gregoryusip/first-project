package main

import (
	"encoding/json"

	"github.com/bitwurx/jrpc2"
)

type Product struct {
	Price    int `json:price`
	Quantity int `json:quantity`
	Name     int `json:name`
}

func (p *Product) FromPositional(params []interface{}) error {
	p.Price = int(params[0].(float64))
	p.Quantity = int(params[1].(float64))
	p.Name = int(params[2].(float64))
	return nil

}

func TotalProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	p := new(Product)
	if err := jrpc2.ParseParams(params, p); err != nil {
		return nil, err
	}
	return p.Price * p.Quantity, nil
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

func main() {
	v1 := jrpc2.NewMuxHandler()
	v1.Register("add", jrpc2.Method{Method: TotalProduct})
	v2 := jrpc2.NewMuxHandler()
	v2.Register("add", jrpc2.Method{Method: AddV2})
	s := jrpc2.NewMuxServer(":8080", nil)
	s.AddHandler("/rpc/v1", v1)
	s.AddHandler("/rpc/v2", v2)
	s.Start()
}
