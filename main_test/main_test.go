package main_test

import (
	"testing"

	"github.com/gregoryusip/first-project/controller"
	"github.com/gregoryusip/first-project/models"
)

type MockRepository struct{}

func NewMockRepository() controller.ProductControllerModel {
	return MockRepository{}
}

func (mock MockRepository) Store(entity models.Products) error {
	return nil
}

func TestInsertModels(t *testing.T) {
	productORM := NewMockRepository()
	productController := controller.NewProductController(controller.Dependencies{
		ProductORM: productORM,
	})

	if err := productController.AddProduct(productORM); err != nil {
		t.Errorf("Got %v, expect nil when inserting new following")
	}
}
