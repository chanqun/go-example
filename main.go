package main

import (
	"fmt"
	"time"
)

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
