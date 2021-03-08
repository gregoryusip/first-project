// https://medium.com/@OmisNomis/creating-an-rpc-server-in-go-3a94797ab833 & https://medium.com/rungo/building-rpc-remote-procedure-call-network-in-go-5bfebe90f7e9

package middleware

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type EditProduct struct {
	Name        string `json:"name"`
	NewName     string `json:"newname"`
	NewPrice    int    `json:"price"`
	NewQuantity int    `json:"quantity"`
}

type Products int

var productSlice []Product

func (p *Products) GetProduct(name string, reply *Product) error {
	var found Product

	for _, v := range productSlice {
		if v.Name == name {
			found = v
		}
	}

	*reply = found

	return nil
}

func (p *Products) MakeProduct(product Product, reply *Product) error {

	productSlice = append(productSlice, product)

	*reply = product

	return nil
}

func (p *Products) EditProduct(product EditProduct, reply *Product) error {

	var editedProduct Product

	for i, v := range productSlice {
		if v.Name == product.Name {
			productSlice[i] = Product{product.NewName, product.NewPrice, product.NewQuantity}
			editedProduct = Product{product.NewName, product.NewPrice, product.NewQuantity}

		}
	}

	*reply = editedProduct

	return nil
}

func (p *Products) DeleteProduct(product Product, reply *Product) error {

	var deletedProduct Product

	for i, v := range productSlice {
		if v.Name == product.Name && v.Price == product.Price && v.Quantity == product.Quantity {
			productSlice = append(productSlice[:i], productSlice[i+1:]...)
			deletedProduct = product
			break
		}
	}

	*reply = deletedProduct

	return nil
}
