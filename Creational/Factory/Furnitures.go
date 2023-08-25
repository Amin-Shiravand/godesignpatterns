package main

type furnitures struct {
	product
}

func createFurnitures(name string, price string) iProduct {

	return &furnitures{
		product: product{
			name:  name,
			price: price,
		},
	}
}
