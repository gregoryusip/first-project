// https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f

package controller

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregoryusip/first-project/mocks"
	"github.com/gregoryusip/first-project/models"
	"github.com/magiconair/properties/assert"
)

// type Products struct {
// 	ID       int    `json:"id"`
// 	Name     string `json:"name"`
// 	Price    int    `json:"price"`
// 	Quantity int    `json:"quantity"`
// }

// type response struct {
// 	ID      int
// 	Message string
// }

func TestAddProduct(t *testing.T) {

	controllers := gomock.NewController(t)
	defer controllers.Finish()

	MockInterface := mocks.NewMockProductModel(controllers)
	// MockInterface := mocks.NewMockProductModel(controller)

	// produkTest := Dependencies{
	// 	ProductORM: MockInterface,
	// }

	// produkTest := ProductRepository{ProductORM: MockInterface}

	var id = 1

	produk := models.Products{
		Name:     "New Product",
		Price:    15000000,
		Quantity: 34,
	}

	// MockInterface.EXPECT().AddProduct()
	MockInterface.EXPECT().CreateProduct(produk).Return(id)

	result := MockInterface.CreateProduct(produk)
	// result := Dependencies.ProductORM.CreateProduct()
	assert.Equal(t, id, result)

}
