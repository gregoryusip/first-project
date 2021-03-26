// https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f

package controller

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gregoryusip/first-project/mocks"
	"github.com/gregoryusip/first-project/models"
	"github.com/mitchellh/mapstructure"
	. "github.com/smartystreets/goconvey/convey"
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

	// var id = 1

	Convey("Create Product Models to be Tested", t, func() {
		productTest := models.Products{
			ID:       1,
			Name:     "New Product",
			Price:    1500000,
			Quantity: 34,
		}
		Convey("Convert Product Models into json.RawMessage Type", func() {
			cvtProduct, err := json.Marshal(productTest)
			if err != nil {
				fmt.Println(err)
				return
			}
			Convey("Call Mock Function with Product Models as a Parameters", func() {
				MockInterface.EXPECT().CreateProduct(productTest).Return(productTest.ID)
				Convey("Convert the Result with Interface Type into Struct Type", func() {
					resultProductInterface, errExp := productControllerTest.AddProduct(cvtProduct)
					resultProductExpected := Expected{}
					mapstructure.Decode(resultProductInterface, &resultProductExpected)
					Convey("Success Compare the Result with the Expected Result", func() {
						expected := Expected{
							ID:      1,
							Message: "Product is inserted",
						}

						So(errExp, ShouldBeNil)
						So(resultProductExpected, ShouldResemble, expected)
					})
					Convey("Error Test When Pass the Empty Product Models", func() {
						productTest2 := models.Products{}
						cvtProduct2, err := json.Marshal(productTest2)
						if err != nil {
							fmt.Println(err)
							// cvtProduct2 = nil
							return
						}
						MockInterface.EXPECT().CreateProduct(productTest2).Return(productTest.ID)
						resultProductInterface2, errExp2 := productControllerTest.AddProduct(cvtProduct2)
						resultProductExpected2 := Expected{}
						mapstructure.Decode(resultProductInterface2, &resultProductExpected2)

						var expectedError int
						if resultProductExpected2.ID == 0 && errExp2 == nil {
							expectedError = 1
						}

						So(expectedError, ShouldNotBeNil)
					})
				})
			})

		})
	})

}
