package models

import (
	"fmt"
	"testing"

	"github.com/gregoryusip/first-project/config"
)

type Produk struct {
	ID       int
	Name     string
	Price    int
	Quantity int
}

func TestCreateProduct(t *testing.T) {
	db := config.CreateConnection()

	productORM := NewProductModel(Dependencies{
		Db: db,
	})

	// productController := controller.ProductControllerModel(controller.Dependencies{
	// 	ProductORM: productORM,
	// })

	produk := Products{
		Name:     "New Product",
		Price:    15000000,
		Quantity: 34,
	}

	// err := productORM.CreateProduct(produk)

	testName := fmt.Sprintf("%s, %d, %d", produk.Name, produk.Price, produk.Quantity)
	t.Run(testName, func(t *testing.T) {
		err := productORM.CreateProduct(produk)

		if err < 0 {
			t.Errorf("Some error is happened")
		}

	})
}
