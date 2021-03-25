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
	fmt.Sprintln("Expected Result for the Test")

	// expected := []Expected{
	// 	{
	// 		ID:      1,
	// 		Message: "Product is inserted",
	// 	},
	// 	{
	// 		ID:      20,
	// 		Message: "Product is inserted",
	// 	},
	// }

	// PRODUCT TEST
	// productTest := []models.Products{
	// 	{
	// 		ID:       1,
	// 		Name:     "New Product",
	// 		Price:    15000000,
	// 		Quantity: 34,
	// 	},
	// 	{
	// 		ID:       2,
	// 		Name:     "New Product 2",
	// 		Price:    240000,
	// 		Quantity: 90,
	// 	},
	// }

	productTest := models.Products{
		ID:       1,
		Name:     "New Product",
		Price:    15000000,
		Quantity: 34,
	}
	fmt.Sprintln("Product Models that being to be Tested")

	// CONVERT PRODUCT STRUCT into JSON.RawMessage
	cvtProduct, err := json.Marshal(productTest)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Sprintln("Convert the Product Models into json.RawMessage type")

	// EXPECTED FUNCTION
	MockInterface.EXPECT().CreateProduct(productTest).Return(id)

	// CONVERT INTERFACE into STRUCT
	resultProductInterface, _ := productControllerTest.AddProduct(cvtProduct)
	resultProductExpected := Expected{}
	mapstructure.Decode(resultProductInterface, &resultProductExpected)
	fmt.Sprintln("Convert the Result with Interface Type into Struct type")

	// EQUAL RESULT with EXPECTED
	assert.Equal(t, resultProductExpected, expected)
	fmt.Sprintln("Success to Compare the Result with the Expected Result")

	// RETURN ERROR FOR UNIT TEST
	expected2 := Expected{
		ID:      20,
		Message: "Product is inserted",
	}
	productTest2 := models.Products{
		ID:       1,
		Name:     "New Product",
		Price:    15000000,
		Quantity: 34,
	}
	cvtProduct2, err := json.Marshal(productTest2)
	if err != nil {
		fmt.Println(err)
		return
	}
	MockInterface.EXPECT().CreateProduct(productTest2).Return(id)

	resultProductInterface2, _ := productControllerTest.AddProduct(cvtProduct2)
	resultProductExpected2 := Expected{}
	mapstructure.Decode(resultProductInterface2, &resultProductExpected2)

	assert.Equal(t, resultProductExpected2, expected2)
	fmt.Sprintln("Fail to Compare the Result with the Expected Result")

}
