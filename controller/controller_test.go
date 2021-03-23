// https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f

package controller

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregoryusip/first-project/config"
	"github.com/gregoryusip/first-project/mocks"
	"github.com/gregoryusip/first-project/models"
)

// type Products struct {
// 	ID       int    `json:"id"`
// 	Name     string `json:"name"`
// 	Price    int    `json:"price"`
// 	Quantity int    `json:"quantity"`
// }

func TestAddProduct(t *testing.T) {
	db := config.CreateConnection("../")

	productORM := models.NewProductModel(models.Dependencies{
		Db: db,
	})

	productController := NewProductController(Dependencies{
		ProductORM: productORM,
	})

	controller := gomock.NewController(t)

	defer controller.Finish()

	MockInterface := mocks.NewMockProductControllerModel(controller)
	// MockInterface := mocks.NewMockProductModel(controller)

	var produk models.Products
	MockInterface.EXPECT().AddProduct(produk)

	// var produk Products
	// MockInterface.EXPECT().AddProduct(produk).

}
