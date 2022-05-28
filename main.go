package main

import (
	"fmt"
	"time"
)

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
