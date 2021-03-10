// // https://medium.com/@OmisNomis/creating-an-rpc-server-in-go-3a94797ab833 & https://medium.com/rungo/building-rpc-remote-procedure-call-network-in-go-5bfebe90f7e9

// package middleware

// import (
// 	"encoding/json"

// 	"github.com/bitwurx/jrpc2"

// 	"github.com/gregoryusip/first-project/models"
// )

// // type Product struct {
// // 	Name     *string `json:"name"`
// // 	Price    *int64  `json:"price"`
// // 	Quantity *int64  `json:"quantity"`
// // }

// // type EditProduct struct {
// // 	Name        string `json:"name"`
// // 	NewName     string `json:"newname"`
// // 	NewPrice    int    `json:"price"`
// // 	NewQuantity int    `json:"quantity"`
// // }

// // var product = new(models.Product)

// // type Products int

// var productSlice []models.Product

// func (p *models.Products) GetProduct(name string, reply *models.Product) error {
// 	var found models.Product

// 	for _, v := range productSlice {
// 		if v.Name == name {
// 			found = v
// 		}
// 	}

// 	*reply = found

// 	return nil
// }

// func (p *models.Product) FromPositional(params []interface{}) error {

// 	// name := string(params[0].(string))
// 	name := params[0].(string)
// 	price := params[1].(int64)
// 	quantity := params[2].(int64)

// 	p.Name = &name
// 	p.Price = &price
// 	p.Quantity = &quantity

// 	return nil
// }

// func CreateProduct(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {

// 	p := new(models.Product)

// 	if err := jrpc2.ParseParams(params, p); err != nil {
// 		return nil, err
// 	}

// 	return nil
// }

// func (p *Products) EditProduct(product EditProduct, reply *Product) error {

// 	var editedProduct Product

// 	for i, v := range productSlice {
// 		if v.Name == product.Name {
// 			productSlice[i] = Product{product.NewName, product.NewPrice, product.NewQuantity}
// 			editedProduct = Product{product.NewName, product.NewPrice, product.NewQuantity}

// 		}
// 	}

// 	*reply = editedProduct

// 	return nil
// }

// func (p *Products) DeleteProduct(product Product, reply *Product) error {

// 	var deletedProduct Product

// 	for i, v := range productSlice {
// 		if v.Name == product.Name && v.Price == product.Price && v.Quantity == product.Quantity {
// 			productSlice = append(productSlice[:i], productSlice[i+1:]...)
// 			deletedProduct = product
// 			break
// 		}
// 	}

// 	*reply = deletedProduct

// 	return nil
// }
