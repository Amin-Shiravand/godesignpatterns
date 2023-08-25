package main

import "fmt"

func newProduct(productType string, name string, price string) (iProduct, error) {

	if productType == "Appliances" {
		return createAppliances(name, price), nil
	} else if productType == "Furnitures" {
		return createFurnitures(name, price), nil
	}

	return nil, fmt.Errorf("no such product type")
}
