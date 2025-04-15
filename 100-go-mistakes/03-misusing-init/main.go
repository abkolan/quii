package main

import "fmt"

var a = func() int {
	fmt.Println("in var")
	return 0
}()

func init() {
	fmt.Println("in init")
}

func main() {
	fmt.Println("in main")
}
