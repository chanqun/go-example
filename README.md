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

### *

```go
func (a *Account) Deposit(amount int) { // don't copy Account
a.balance += amount
}
```

### error

error or nil
```go
// Withdraw withdraw
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("can't withdraw you are poor")
	}

	a.balance -= amount
	return nil
}

if err != nil {
    log.Fatalln(err) // print and exit
}
```

### 동기
```go
package main

import (
	"errors"
	"fmt"
	http "net/http"
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
		"https://www.yahoo.com",
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	var results = make(map[string]string)

	for _, url := range urls {
		err := hitUrl(url)

		var result = "OK"
		if err != nil {
			result = "FAILED"
		}

		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed = errors.New("request failed")

func hitUrl(url string) error {
	fmt.Println("Checking", url)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}

```

### goroutine

```go
func main() {
	go sexyCount("chanqun")
	sexyCount("sungmin")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

```

### channel
```go
func main() {
	c := make(chan bool)
	people := [2]string{"chanqun", "sungmin"}

	for _, person := range people {
		go isSexy(person, c)
	}

	result := <-c
	result2 := <-c

	fmt.Println(result, result2)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 3)
	c <- true
}
```


최종 메인 코드

```go
package main

import (
	"github.com/chanqun/go-example/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("jobs.csv")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment("jobs.csv", "jobs.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
```