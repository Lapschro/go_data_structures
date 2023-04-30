package stack

import "fmt"

type Stack[T any] struct {
	top *node[T]
}

type node[T any] struct {
	next *node[T]
	val  T
}

func (s *Stack[T]) Push(val T) {
	n := newNode(val)
	n.next = s.top
	s.top = &n
}

func (s *Stack[T]) Pop() T {
	current := s.top
	s.top = current.next
	current.next = nil
	return current.val
}

func (s Stack[T]) Print() {
	fmt.Println(s)
}

func (s Stack[T]) Top() T {
	var zv T

	if s.top == nil {
		return zv
	}

	return s.top.val
}

func Init[T any]() Stack[T] {
	return Stack[T]{
		top: nil,
	}
}

func newNode[T any](val T) node[T] {
	return node[T]{
		next: nil,
		val:  val,
	}
}
