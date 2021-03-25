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
	fmt.Printf("Expected Result for the Test: %v\n", expected)

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
	fmt.Printf("Product Models that being to be Tested: %v\n", productTest)

	// CONVERT PRODUCT STRUCT into JSON.RawMessage
	fmt.Println("Convert the Product Models into json.RawMessage type")
	cvtProduct, err := json.Marshal(productTest)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Product Models in json.RawMessage type: %v\n", cvtProduct)

	// EXPECTED FUNCTION
	MockInterface.EXPECT().CreateProduct(productTest).Return(id)

	// CONVERT INTERFACE into STRUCT
	fmt.Println("Convert the Result with Interface Type into Struct type")
	resultProductInterface, _ := productControllerTest.AddProduct(cvtProduct)
	resultProductExpected := Expected{}
	mapstructure.Decode(resultProductInterface, &resultProductExpected)
	fmt.Printf("Result Product in Struct type: %v\n", resultProductExpected)

	// EQUAL RESULT with EXPECTED
	fmt.Print("Success to Compare the Result with the Expected Result\n\n")
	assert.Equal(t, resultProductExpected, expected)

	// RETURN ERROR FOR UNIT TEST
	expectedError := Expected{
		ID:      20,
		Message: "Product is inserted",
	}
	fmt.Printf("Expected Error Result for the Test: %v\n\n", expectedError)

	cvtProduct, err = json.Marshal(productTest)
	if err != nil {
		fmt.Println(err)
		return
	}
	MockInterface.EXPECT().CreateProduct(productTest).Return(id)

	resultProductInterfaceError, _ := productControllerTest.AddProduct(cvtProduct)
	resultProductExpectedError := Expected{}
	mapstructure.Decode(resultProductInterfaceError, &resultProductExpectedError)

	assert.Equal(t, resultProductExpectedError, expectedError)
	fmt.Print("Fail to Compare the Result with the Expected Result\n\n")

}
