package main

type appliances struct {
	product
}

func createAppliances(name string, price string) iProduct {

	return &appliances{
		product{
			name:  name,
			price: price,
		},
	}
}
