package tree

import "testing"

func TestTreeInit(t *testing.T) {
	tree := Init[int]()

	if tree.root != nil {
		t.Fatalf("Tree was not properly instantiated")
	}
}

func TestTreeAdd1(t *testing.T) {
	tree := Init[int]()

	tree.Add(1)

	if *tree.root.val != 1 {
		t.Fatalf("Expected root node to be 1 got:%d", *tree.root.val)
	}
}

func TestTreeAdd(t *testing.T) {
	tree := Init[int]()

	tree.Add(1)

	if *tree.root.val != 1 {
		t.Fatalf("Expected root node to be 1 got:%d", *tree.root.val)
	}

	tree.Add(0)

	if *tree.root.left.val != 0 {
		t.Fatalf("Expected 0 got: %d", *tree.root.left.val)
	}

	tree.Add(2)

	if *tree.root.right.val != 2 {
		t.Fatalf("Expected 2 got: %d", *tree.root.right.val)
	}

	tree.Add(3)

	if *tree.root.right.right.val != 3 {
		t.Fatalf("Expected 3 got: %d", *tree.root.right.right.val)
	}
}
