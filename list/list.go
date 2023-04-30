package list

import "fmt"

type node[T any] struct {
	val  T
	next *node[T]
}

type List[T any] struct {
	length int
	first  *node[T]
}

func Init[T any]() List[T] {
	return List[T]{
		length: 0,
		first:  nil,
	}
}

func (l *List[T]) EmplaceBack(val T) {
	var lastNode *node[T]
	for current := l.first; current != nil; current = current.next {
		if current.next == nil {
			lastNode = current
		}
	}
	newNode := node[T]{
		val:  val,
		next: nil,
	}
	l.length++
	if lastNode == nil {
		l.first = &newNode
		return
	}
	lastNode.next = &newNode
}

func (l *List[T]) Push(val T) {
	current := l.first
	newNode := node[T]{
		val:  val,
		next: nil,
	}
	l.length++
	if current == nil {
		l.first = &newNode
		return
	}

	newNode.next = current
	l.first = &newNode
}

func (l List[T]) ElementAt(index int) *T {
	if index > l.length || index < 0 {
		return nil
	}

	x := 0
	for current := l.first; current != nil; current = current.next {
		if x == index {
			return &current.val
		}
		x++
	}
	return nil
}

func (l *List[T]) AddAt(index int, val T) {
	var prev *node[T]

	if index > l.length || index < 0 {
		return
	}

	node := node[T]{
		val:  val,
		next: nil,
	}

	prev = nil

	x := 0

	if l.first == nil {
		l.first = &node
		l.length++
		return
	}

	for current := l.first; current != nil; current = current.next {
		if x == index {
			if prev == nil {
				l.first = &node
			} else {
				prev.next = &node
			}
			node.next = current
			l.length++
			return
		}
		prev = current
		x++
	}
}

func (l *List[T]) Reverse() {
	var prev *node[T]
	var current *node[T]

	if l.first == nil {
		return
	}

	current = l.first

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	l.first = prev
}

func (l *List[T]) Pop() *T {
	current := l.first

	if current == nil {
		return nil
	}

	l.first = current.next
	current.next = nil
	l.length--
	return &current.val
}

func (l *List[T]) Print() {
	for current := l.first; current != nil; current = current.next {
		fmt.Println(*&current.val)
	}
}

func (l *List[T]) RemoveAt(index int) {
	var prev *node[T]

	if index > l.length || index < 0 {
		return
	}

	prev = nil

	x := 0
	for current := l.first; current != nil; current = current.next {
		if x == index {
			if prev == nil {
				l.first = current.next
			} else {
				prev.next = current.next
			}

			current.next = nil
			l.length--
			return
		}
		prev = current
		x++
	}
}
