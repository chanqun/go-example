package main

import "fmt"

func main() {
	names := []string{"chanqun", "len"}

	names = append(names, "sung")

	fmt.Println(names)
}
