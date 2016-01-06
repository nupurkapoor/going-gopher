package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	var h, i, j = 1, 2, 3 //infer types

	var k, l, m = 1, false, "three" //infer mixed-up types

	fmt.Println(h, i, j) // prints 1 2 3

	fmt.Println(k, l, m) // prints 1 false three

	fmt.Println(a) // am not sure how this is different than using %v, looks same! ??
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Println(d,e,f,g) // prints all in same line

	fmt.Println("####### %T version #######")

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
}