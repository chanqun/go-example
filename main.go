package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "don"}
	chanqun := person{age: 18, name: "chanqun", favFood: favFood}

	fmt.Println(chanqun)
}
