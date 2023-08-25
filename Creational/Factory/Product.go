package main

type iProduct interface {
	getPrice() string
	setPrice(price string)
	getName() string
	setName(name string)
}

type product struct {
	name  string
	price string
}

func (p *product) setName(name string) {
	p.name = name
}

func (p *product) setPrice(price string) {
	p.price = price
}

func (p *product) getPrice() string {
	return p.price
}

func (p *product) getName() string {
	return p.name
}
