package main

import "fmt"

func main() {
	a := 2
	b := a  // copy value
	c := &a // copy memory address

	a = 4

	fmt.Println(a, b, *c) // 4 2 4
}
