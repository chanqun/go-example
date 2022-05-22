package main

import (
	"fmt"
)

func main() {
	// 1
	fmt.Println("Hello world!") // formatting package

	// 2
	const name string = "chanqun" // go는 type 언어
	const flag bool = true

	var changeName string = "chanqun"
	changeName = "cq"
	otherName := "sung"

	fmt.Println(changeName)
	fmt.Println(otherName)

	// 3
}
