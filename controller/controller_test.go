package controller

import (
	"fmt"
	"testing"
)

func TestAddProduct(t *testing.T) {
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

	testingModel := new(Products)
	testingModel.ID = 10
	testingModel.Name = "Model"
	testingModel.Price = 80000
	testingModel.Quantity = 10

	for _, testController := range tests {
		testName := fmt.Sprintf("%s, %d, %d", testController.Name, testController.Price, testController.Quantity)
		t.Run(testName, func(t *testing.T) {
			ans := ProductRepository.ProductORM.CreateProduct(testingModel)
		})
	}
}
