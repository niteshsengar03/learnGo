package main

import "fmt"

type Complex struct {
	realNo int
	imagNo int
}

func createComplexNo(realNo int, imagNo int) *Complex {
	c := Complex{
		realNo: realNo,
		imagNo: imagNo,
	}
	return &c
}

func (c *Complex) add(a *Complex) *Complex {
	addImag := c.imagNo + a.imagNo
	addReal := c.realNo + a.realNo
	afterAdd := Complex{
		realNo: addReal,
		imagNo: addImag,
	}
	return &afterAdd
}

func (c *Complex) display() {
	fmt.Println(c.realNo, "+", c.imagNo, "i")
}

func main() {
	c1 := createComplexNo(1, 2)
	c2 := createComplexNo(1, 2)
	c1.display()
	c2.display()
	c3 := c1.add(c2)
	c3.display()
}
