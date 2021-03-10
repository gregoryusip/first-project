package controller

import (
	"encoding/json"

	"github.com/bitwurx/jrpc2"
	"github.com/gregoryusip/first-project/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    []models.Product `json:"data"`
}

func (p *models.Product) FromPositional(params []interface{}) error {
	p.Name = string(params[0].(string))
	p.Price = int(params[1].(int))
	p.Quantity = int(params[2].(int))
	return nil
}

func AddProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	p := new(models.Product)

	if err := jrpc2.ParseParams(params, p); err != nil {
		return nil, err
	}

	insertID := models.CreateProduct(p)

	res := response{
		ID:      insertID,
		Message: "Product is inserted",
	}

	return res, nil
}
