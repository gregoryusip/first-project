// https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f

package controller

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregoryusip/first-project/mocks"
	"github.com/gregoryusip/first-project/models"
)

type ProductsJSON struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

// type response struct {
// 	ID      int
// 	Message string
// }

func TestAddProduct(t *testing.T) {

	controllers := gomock.NewController(t)
	defer controllers.Finish()

	MockInterface := mocks.NewMockProductModel(controllers)

	// produkTest := ProductRepository{ProductORM: MockInterface}
	produkTest := NewProductController(Dependencies{ProductORM: MockInterface})

	var id = 1

	produk1 := models.Products{
		ID:       1,
		Name:     "New Product",
		Price:    15000000,
		Quantity: 34,
	}

	// produk2 := []byte(`
	// {
	// 	"id": 1,
	// 	"name": "Meja Belajar",
	// 	"price": 150000,
	// 	"quantity": 10,
	// }
	// `)

	testProduct := ProductsJSON{ID: 1, Name: "New Product", Price: 1500000, Quantity: 34}
	resultProduk, err := json.Marshal(testProduct)
	if err != nil {
		fmt.Println(err)
		return
	}

	MockInterface.EXPECT().CreateProduct(testProduct).Return(id)

	result, _ := produkTest.AddProduct(resultProduk)
	fmt.Println(testProduct)
	fmt.Println(result)

	// assert.Equal(t, testProduct, result)

}
