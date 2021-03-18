package models

import (
	"fmt"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	type Produk struct {
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

	// var testModels = struct {
	// 	ID       int
	// 	Name     string
	// 	Price    int
	// 	Quantity int
	// }{
	// 	ID:       10,
	// 	Name:     "Playstation 5",
	// 	Price:    8000000,
	// 	Quantity: 50,
	// }

	testModelss := Produk{"Playstation 5", 8000000, 50}

	for _, testModels := range tests {
		testName := fmt.Sprintf("%s, %d, %d", testModels.Name, testModels.Price, testModels.Quantity)
		t.Run(testName, func(t *testing.T) {
			ans := ProductModel.CreateProduct(testModelss)

			if ans < 0 {
				t.Errorf("Some error is happened")
			}
		})
	}
}
