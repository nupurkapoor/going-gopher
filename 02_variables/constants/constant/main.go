package main

import "fmt"

const p string = "death & taxes without string"
const q = "death & taxes with string"

func main() {

	const (
		r int = 42
		s int = 45
	)

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("r - ", r)
	fmt.Println("r - ", s)
}

// a CONSTANT is a simple unchanging value