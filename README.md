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

### switch

### pointers
```go
package main

import "fmt"

func main() {
	a := 2
	b := a  // copy value
	c := &a // copy memory address

	a = 4
	
	// *b = 20 using b update a

	fmt.Println(a, b, *c) // 4 2 4
}
```

### array
크기를 명시해야 한다
```go
names := [5]string{"chanqun"}
```

### slice
without length

```go
names := []string{"chanqun"}
namse = append(names, "sung")
```
append return new slice

### map

```go
mapA := map[string]string{"name": "chanqun", "age": "30"}
	
for key, value := range mapA {
    fmt.Println(key, value)		
}
```

### struct

```go
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

```

### Account 객체

```go
package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner}

	return &account
}
```
