package main

import "fmt"

func Filter[T any](arr []T, f func(T)bool) []T {
	result := []T{}

	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5}

	b := a[2:]

	fmt.Println(a, b)

	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	b[2] = 42

	fmt.Println(&a[2], &b[0])

	fmt.Println(a, b)

	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
}