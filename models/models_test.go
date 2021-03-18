package models

import (
	"fmt"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	type Produk struct {
		ID       int
		Name     string
		Price    int
		Quantity int
	}

	var tests = []struct {
		Name     string
		Price    int
		Quantity int
	}{
		{"Bola Basket", 250000, 400},
		{"Meja Makan", 450000, 100},
		{"TV LED", 1500000, 250},
		{"Playstation 5", 8000000, 50},
	}

	// testingModel := Products{10, "Playstation 5", 8000000, 50}
	testingModel := new(Products)
	testingModel.ID = 10
	testingModel.Name = "Model"
	testingModel.Price = 80000
	testingModel.Quantity = 10

	testName := fmt.Sprintf("%s, %d, %d", testingModel.Name, testingModel.Price, testingModel.Quantity)
	t.Run(testName, func(t *testing.T) {
		ans := ProductModel.CreateProduct(testingModel)

		if ans < 0 {
			t.Errorf("Some error is happened")
		}
	})

	// for _, testModels := range tests {
	// 	testName := fmt.Sprintf("%s, %d, %d", testModels.Name, testModels.Price, testModels.Quantity)
	// 	t.Run(testName, func(t *testing.T) {
	// 		ans := ProductModel.CreateProduct(testingModel)

	// 		if ans < 0 {
	// 			t.Errorf("Some error is happened")
	// 		}
	// 	})
	// }
}
