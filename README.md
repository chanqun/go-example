# go-example

main은 컴파일을 원할때 사용

function 을 대문자로 시작하면 export 하는 것


```go
package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

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

	// 3 function part One
	// function has multiple result
	fmt.Println(multiply(3, 4))
	
	totalLength, upperName := lenAndUpper("chanqun")
	fmt.Println(totalLength, upperName)
}
```