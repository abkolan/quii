package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")
	time.Sleep(2 * time.Second)
	fmt.Println("Hey")
}
