package main

import (
	"fmt"
	"github.com/chanqun/go-example/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	definition, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(definition)
}
