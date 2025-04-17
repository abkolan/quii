package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func MapValues[T constraints.Ordered](values []T, f func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := f(v)
		newValues = append(newValues, newValue)
	}

	return newValues
}
func main() {

	result := MapValues([]int{1, 2, 3}, func(i int) int {
		return i * i
	})

	fmt.Println(result)

}
