package main

import "fmt"

func main() {
	ch := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch <- "data"
	}()

	go func() {
		ch2 <- "data2"
	}()
	select {
	case msgFromChannel := <-ch:
		fmt.Println(msgFromChannel)
	case msgFromAnotherChannel := <-ch2:
		fmt.Println(msgFromAnotherChannel)
	}
}
