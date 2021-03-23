// https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f

package controller

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregoryusip/first-project/mocks"
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

	MockInterface := mocks.NewMockProductControllerModel(controllers)
	// MockInterface := mocks.NewMockProductModel(controller)

	// produkTest := Dependencies{
	// 	ProductORM: MockInterface,
	// }

	produkTest := &ProductRepository{ProductORM: MockInterface}

	var id = 1
	res := response{
		ID:      id,
		Message: "Mocking Success!",
	}

	// produk := models.Products{
	// 	Name:     "New Product",
	// 	Price:    15000000,
	// 	Quantity: 34,
	// }

	produk := []byte(`
	{
		"name": "New Product",
		"price": 200000,
		"quantity": 10,
	}
	`)

	// MockInterface.EXPECT().AddProduct()
	MockInterface.EXPECT().AddProduct(produk).Return(res, nil)

	// result := ProductRepository.ProductORM.CreateProduct(models.Products(*&produk))
	// result := Dependencies.ProductORM.CreateProduct()
	// assert.Equal(t, res.ID, result)

}
