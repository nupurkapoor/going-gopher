
package main

import "fmt"

func main() {

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c) //outputs 0, shouldnt it be 0.0 ??
	fmt.Printf("%v \n", d)

	fmt.Println()
}