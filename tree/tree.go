package tree

import (
	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	val   *T
	right *node[T]
	left  *node[T]
}

type Tree[T constraints.Ordered] struct {
	root *node[T]
}

func Init[T constraints.Ordered]() Tree[T] {
	return Tree[T]{
		root: nil,
	}
}

func (t *Tree[T]) Add(val T) {
	n := node[T]{
		val:   &val,
		right: nil,
		left:  nil,
	}

	if t.root == nil {
		t.root = &n
		return
	}

	t.root.add(val)
}

func initNode[T constraints.Ordered](val T) node[T] {
	return node[T]{
		val:   &val,
		right: nil,
		left:  nil,
	}
}

func (n *node[T]) add(val T) {
	if *(n.val) > val {
		if n.left == nil {
			node := initNode[T](val)
			n.left = &node
		} else {
			n.left.add(val)
		}
	} else {
		if n.right == nil {
			node := initNode[T](val)
			n.right = &node
		} else {
			n.right.add(val)
		}
	}
}
