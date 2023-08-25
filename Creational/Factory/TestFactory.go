package main

import "fmt"

func main() {
	product1, _ := newProduct("Appliances", "Refrigerator", "800")
	product2, _ := newProduct("Furnitures", "Sofa", "400")

	productDetails(product1)
	productDetails(product2)
}

func productDetails(product iProduct) {
	fmt.Println("product name", product.getName())
	fmt.Println("product price", product.getPrice())
}
