package main

import "fmt"

type stack[T any] struct {
	elements []T
}

func (s *stack[T]) push(elem T) {

	s.elements = append(s.elements, elem)
}

func (s *stack[T]) pop() (T, bool) {

	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true

}

func (s *stack[T]) isEmpty() bool {

	return len(s.elements) == 0
}

func (s stack[T]) printall() {

	if len(s.elements) == 0 {
		fmt.Println("The stack is empty.")
		return
	}

	for _, element := range s.elements {
		fmt.Println(element)
	}
}

func main() {

	intStack := stack[int]{}
	fmt.Println("Pushing Values to stack.")
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.push(4)
	intStack.printall()
	fmt.Println("Poping Values from stack.")
	fmt.Println(intStack.pop())
	intStack.printall()
	fmt.Println("Is Stack Empty: ", intStack.isEmpty())

}
