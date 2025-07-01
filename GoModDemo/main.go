package main

import (
	app "GoModDemo/app"
)

// if you want a packge to access a method or variable from other package name it first Letter uppercase it's like public in java
// and if you don't name with firt letter lowercase


func main() {
	c1 := app.CreateComplexNo(1, 2)
	c2 := app.CreateComplexNo(1, 2)
	c1.Display()
	c2.Display()
	c3 := c1.Add(c2)
	c3.Display()
}
