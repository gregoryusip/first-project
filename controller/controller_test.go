package controller

import (
	"fmt"
	"testing"

	"github.com/gregoryusip/first-project/config"
	"github.com/gregoryusip/first-project/models"
)

func TestAddProduct(t *testing.T) {
	db := config.CreateConnection()

	productORM := models.NewProductModel(models.Dependencies{
		Db: db,
	})

	productController := NewProductController(Dependencies{
		ProductORM: productORM,
	})

	produk := Products{
		Name:     "New Product",
		Price:    15000000,
		Quantity: 34,
	}

	testName := fmt.Sprintf("%s, %d, %d", produk.Name, produk.Price, produk.Quantity)
	t.Run(testName, func(t *testing.T) {
		err := productController.AddProduct(models.Products(*produk))

		if err < 0 {
			t.Errorf("Some error is happened")
		}

	})
}
