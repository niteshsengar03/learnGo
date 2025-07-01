// struct-and-methods
package main

import "fmt"

// no conecpt of class in go
type Product struct {
	name    string
	price   int
	company string
}

// func that act as constructor
func newProduct(name string, price int, company string) Product {
	p := Product{
		name:    name,
		price:   price,
		company: company,
	}
	return p // this is returning a copy of the product
	// here two times object is created first when originally creted
	// then a copy of it created and that is returned
}

func newProducts(name string, price int, company string) *Product {
	p := Product{
		name:    name,
		price:   price,
		company: company,
	}
	return &p // this is returning the address of P
}

func makeChangeinAcutalObject(p *Product) {
	p.name = "HeroLaal"
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
	// no direct concept of constructor
	// this is the way of creating object
	pr := Product{
		name:    "Nitesh",
		price:   1000,
		company: "dev",
	}
	fmt.Println(pr, pr.name)
	// creating new product by our function
	pr2 := newProduct("Alic", 100, "Bob")
	fmt.Println(pr2, pr2.name)
	pr3 := newProducts("Alice", 200, "Bobs")
	fmt.Println(pr3)
	makeChangeinAcutalObject(pr3)
	fmt.Println(pr3.name)

	pr3.display(true)
	pr3.display(false)

}
