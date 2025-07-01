package app

import "fmt"

type Complex struct {
	realNo int
	imagNo int
}

func CreateComplexNo(realNo int, imagNo int) *Complex {
	c := Complex{
		realNo: realNo,
		imagNo: imagNo,
	}
	return &c
}

func (c *Complex) Add(a *Complex) *Complex {
	addImag := c.imagNo + a.imagNo
	addReal := c.realNo + a.realNo
	afterAdd := Complex{
		realNo: addReal,
		imagNo: addImag,
	}
	return &afterAdd
}

func (c *Complex) Display() {
	fmt.Println(c.realNo, "+", c.imagNo, "i")
}

