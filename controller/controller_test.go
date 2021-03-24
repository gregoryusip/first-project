// https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f

package controller

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregoryusip/first-project/mocks"
	"github.com/gregoryusip/first-project/models"
	"github.com/magiconair/properties/assert"
	"github.com/mitchellh/mapstructure"
)

type Expected struct {
	ID      int
	Message string
}

func TestAddProduct(t *testing.T) {

	// MOCK Controller
	controllers := gomock.NewController(t)
	defer controllers.Finish()

	// MOCK INTERFACE from MODELS
	MockInterface := mocks.NewMockProductModel(controllers)

	// INJECT MOCK INTERFACE into PRODUCT CONTROLLER
	productControllerTest := NewProductController(Dependencies{ProductORM: MockInterface})

	var id = 1

	// EXPECTATION RESULT
	expected := Expected{
		ID:      1,
		Message: "Product is inserted",
	}

	// PRODUCT TEST
	productTest := models.Products{
		ID:       1,
		Name:     "New Product",
		Price:    15000000,
		Quantity: 34,
	}

	// CONVERT PRODUCT STRUCT into JSON.RawMessage
	cvtProduct, err := json.Marshal(productTest)
	if err != nil {
		fmt.Println(err)
		return
	}

	// EXPECTED FUNCTION
	MockInterface.EXPECT().CreateProduct(productTest).Return(id)

	// CONVERT INTERFACE into STRUCT
	resultProductInterface, _ := productControllerTest.AddProduct(cvtProduct)
	resultProductExpected := Expected{}
	mapstructure.Decode(resultProductInterface, &resultProductExpected)

	// EQUAL RESULT with EXPECTED
	assert.Equal(t, resultProductExpected, expected)

}
