package main

import (
	"fmt"
	"reflect"
)

type Stack []interface{}

func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

type OutStack struct {
	Stack
}

func (os *OutStack) Put(elem int) {
	os.Stack.Push(elem)
}

func main() {
	// You can push an int and a string to the same stack
	// myStack := &Stack{}
	// myStack.Push(10)
	// myStack.Push("hello")
	// fmt.Println(*myStack...)
	// fmt.Println("Emptying stack")
	// a := myStack.Pop().(int)
	// fmt.Println(a)
	// b := myStack.Pop().(int)
	// fmt.Println(b)

	// //create outer stack
	// myOuterStack := &OutStack{Stack{}}
	// myOuterStack.Put(1)
	// // myOuterStack.Put("As")
	// myOuterStack.IsEmpty()

	s := []interface{}{true, 0, "yes"}
	e := s[0]
	fmt.Println(e, reflect.TypeOf(e))
	//this does not works
	// if e {
	// 	fmt.Println("The condition is true")
	// }
	//this can be done but may create panic
	// if e.(bool) {
	// 	fmt.Println("The condition is true")
	// }
	if e == true {
		fmt.Println("The condition is true")
	}

}

// type Stack[T any] []T

// func (s *Stack[T]) Push(v T) {
//     *s = append(*s, v)
// }

// func (s *Stack[T]) Pop() T {
//     if s.IsEmpty() {
//         var zero T // Returns the zero value of the type
//         if s.IsEmpty() {
//             return zero
//         }
//     }
//     top := (*s)[len(*s)-1]
//     *s = (*s)[:len(*s)-1]
//     return top
// }

// func (s *Stack[T]) IsEmpty() bool {
//     return len(*s) == 0
// }

// func main() {
//     // You must specify the type when creating the stack
//     intStack := &Stack[int]{}
//     intStack.Push(10)

//     stringStack := &Stack[string]{}
//     stringStack.Push("hello")

//     // This will cause a compile-time error!
//     // intStack.Push("world")
// }







// Go has no automatic type conversion 
// Interface values are of type interface
// until the dynamic type is accessed
