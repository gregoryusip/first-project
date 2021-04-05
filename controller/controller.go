// https://codesource.io/build-a-crud-application-in-golang-with-postgresql/

package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bitwurx/jrpc2"
	"github.com/gregoryusip/first-project/models"
)

type Products struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type ProductControllerModel interface {
	AddProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject)
	ReadedProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject)
	UpdatedProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject)
	DeletedProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject)
}

type Dependencies struct {
	ProductORM models.ProductModel
}

func NewProductController(deps Dependencies) ProductControllerModel {
	return &ProductRepository{
		ProductORM: deps.ProductORM,
	}
}

type ProductRepository struct {
	ProductORM models.ProductModel
}

func (p *Products) FromPositional(params []interface{}) error {

	// id := int(params[0].(int))
	name := string(params[1].(string))
	price := int(params[2].(int))
	quantity := int(params[3].(int))

	// p.ID = id
	p.Name = name
	p.Price = price
	p.Quantity = quantity

	// p.Name = string(params[0].(string))
	// p.Price = int(params[1].(int))
	// p.Quantity = int(params[2].(int))

	return nil
}

func (p *ProductRepository) AddProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {

	produk := new(Products)
	// var produk Products

	if err := jrpc2.ParseParams(params, produk); err != nil {
		return nil, err
	}

	// insertID := models.ProductModel.CreateProduct(produk)
	insertID := p.ProductORM.CreateProduct(models.Products(*produk))
	// _, err := p.ProductORM.CreateProductPG("Testing", models.Products(*produk))

	res := response{
		ID:      insertID,
		Message: "Product is inserted",
	}

	return res, nil
}

func (p *ProductRepository) ReadedProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {

	produk, err := p.ProductORM.ReadProduct()
	// produk, err := p.ProductORM.ReadProductPG()

	// if err := jrpc2.ParseParams(params, p); err != nil {
	// 	return nil, err
	// }

	if err != nil {
		log.Fatalf("Can't take the data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = produk

	status := 1

	res := Response{
		Status:  status,
		Message: "Success",
		Data:    produk,
	}

	return res, nil
}

func (p *ProductRepository) UpdatedProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {

	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	log.Fatalf("Can't change from string to int. %v", err)
	// }

	produk := new(Products)

	if err := jrpc2.ParseParams(params, produk); err != nil {
		return nil, err
	}

	// updatedProducts := models.UpdateProduct(int64(id), p)
	updatedProduct := p.ProductORM.UpdateProduct(produk.Name, models.Products(*produk))

	msg := fmt.Sprintf("Product success to update.")

	res := response{
		ID:      int(updatedProduct),
		Message: msg,
	}

	return res, nil
}

func (p *ProductRepository) DeletedProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {

	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	log.Fatalf("Can't change from string to int. %v", err)
	// }

	produk := new(Products)

	if err := jrpc2.ParseParams(params, produk); err != nil {
		return nil, err
	}

	// deletedProducts := models.DeleteProduct(int64(id))
	deletedProduct := p.ProductORM.DeleteProduct(produk.Name)

	msg := fmt.Sprintf("Product success to delete.")

	res := response{
		ID:      int(deletedProduct),
		Message: msg,
	}

	return res, nil
}

type response struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []models.Products `json:"data"`
}
