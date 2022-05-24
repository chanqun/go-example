package main

import "fmt"

func main() {
	mapA := map[string]string{"name": "chanqun", "age": "30"}

	for key, value := range mapA {
		fmt.Println(key, value)
	}
}
