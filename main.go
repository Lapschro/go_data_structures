package main

import (
	"fmt"
	"gokita/ds/list"
	"gokita/ds/stack"
)

func main() {
	s := stack.Init[int]()
	s.Push(5)
	s.Print()
	fmt.Printf("%d\n", s.Top())
	s.Push(10)
	s.Print()
	fmt.Printf("%d\n", s.Top())
	s.Print()
	fmt.Printf("%d\n", s.Pop())
	s.Print()
	fmt.Printf("%d\n", s.Top())

	l := list.Init[int]()
	l.AddAt(0, -1)
	l.Push(0)
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	l.Print()

	l.RemoveAt(0)
	l.Print()

	fmt.Printf("Hello world!\n")
}
