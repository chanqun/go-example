package something

import "fmt"

func sayBye() {
	fmt.Println("Bye")
}

func SayHello() { //대문자만 export 가능
	fmt.Println("Hello")
}
