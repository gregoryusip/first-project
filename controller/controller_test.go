// https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f

package controller

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregoryusip/first-project/config"
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
	db := config.CreateConnection("../")

	productORM := models.NewProductModel(models.Dependencies{
		Db: db,
	})

	productControllers := NewProductController(Dependencies{
		ProductORM: productORM,
	})

	if productControllers != nil {
		return
	}

	var p *ProductRepository

	controllers := gomock.NewController(t)

	defer controllers.Finish()

	MockInterface := mocks.NewMockProductControllerModel(controllers)
	// MockInterface := mocks.NewMockProductModel(controller)

	var id = 1
	res := response{
		ID:      id,
		Message: "Mocking Success!",
	}

	produk := Products{
		Name:     "New Product",
		Price:    15000000,
		Quantity: 34,
	}

	// MockInterface.EXPECT().AddProduct()
	MockInterface.EXPECT().AddProduct(produk).Return(res, nil)

	result := p.ProductORM.CreateProduct(models.Products(*&produk))
	// result := Dependencies.ProductORM.CreateProduct()
	assert.Equal(t, res.ID, result)

}
