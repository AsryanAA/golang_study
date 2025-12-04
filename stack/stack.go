package main

import (
	"fmt"
	"slices"
)

type Stack struct {
	arr []int
}

func (s *Stack) Pop() int {
	result := s.arr[0]

	s.arr = s.arr[1:]

	return result
}

func (s *Stack) Push(i int) {
	s.arr = slices.Insert(s.arr, 0, i)
}

func main() {
	stask := &Stack{}

	stask.Push(42)
	stask.Push(16)
	stask.Push(22)

	fmt.Println(stask)

	stask.Pop()

	fmt.Println(stask)
}