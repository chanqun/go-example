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

func repeatMe(words ...string) {
	fmt.Println(words)
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

	repeatMe("chanqun", "nico") // [chanqun nico]
}
```

### naked return

```go
func lenAndUpper(name string) (length int, uppercase string) {
defer fmt.Println("I'm done")

length = len(name)
uppercase = strings.ToUpper(name)
return
}
```

### defer

execute code after function end

### only have for

```go
func superAdd(numbers ...int) int {
    total := 0

    for _, number := range numbers {
        total += number
    }

    return total
}
```

### if else
if 안에 변수를 선언해서 잠시만 사용할 수 있다
variable expression

```go
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}
```

