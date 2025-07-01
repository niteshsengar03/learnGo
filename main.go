// struct-and-methods
package main

import "fmt"

type sellableProducts interface {
	buy()  
	getDiscount() int // return type of getDiscout() must be int
}

type Product struct {
	name    string
	price   int
	company string
}


func (p *Product) buy(){
	fmt.Println("Buying the product:",p.name,",from company:",p.company)
}

func (p *Product) getDiscount()int{
	discount := p.price * 20 / 100
	return discount
}

func getDiscountAndBuy(p sellableProducts){
	discount := p.getDiscount()
	if discount>30 {
		fmt.Println("Discount is good")
		p.buy();
	} else {
		fmt.Println("Discount is not good, not buying")
	}
}

func newProducts(name string, price int, company string) *Product {
	p := Product{
		name:    name,
		price:   price,
		company: company,
	}
	return &p // this is returning the address of P
}

// mebmer function of struct Product
func (p *Product) display(short bool) {
	fmt.Println()
	if short {
		fmt.Println("name is ", p.name)
		return
	}
	fmt.Println("Product details are:")
	fmt.Println("name: ", p.name)
	fmt.Println("price: ", p.price)
	fmt.Println("company: ", p.company)
}

func main() {
	pr := newProducts("Alice", 2, "Bobs")
	getDiscountAndBuy(pr)
	pr.display(false)
}
