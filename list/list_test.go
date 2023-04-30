package list

import "testing"

func TestInit(t *testing.T) {
	l := Init[int]()
	if l.length != 0 {
		t.Fatalf("Expected length 0 got: %d\n", l.length)
	}
	if l.first != nil {
		t.Fatalf("Expected first to be null pointer\n")
	}
}
func TestPush(t *testing.T) {
	l := Init[int]()
	l.Push(0)
	if *l.ElementAt(0) != 0 {
		t.Fatalf("Expected 0 got: %d\n", *l.ElementAt(0))
	}
}

func TestAddElementAt(t *testing.T) {
	l := Init[int]()

	l.Push(0)
	l.EmplaceBack(1)

	if *l.ElementAt(1) != 1 {
		t.Fatalf("Excpected 1 got: %d", *l.ElementAt(1))
	}
	if *l.ElementAt(0) != 0 {
		t.Fatalf("Expected 0 got: %d", *l.ElementAt(0))
	}
}

func TestAddAt(t *testing.T) {
	l := Init[int]()

	l.Push(0)
	l.EmplaceBack(2)
	l.AddAt(1, 1)

	if *l.ElementAt(1) != 1 {
		t.Fatalf("Expected 1 got: %d", *l.ElementAt(1))
	}
}

func TestReverse(t *testing.T) {
	l := Init[int]()

	l.Push(0)
	l.EmplaceBack(2)
	l.AddAt(1, 1)

	l.Reverse()

	if *l.ElementAt(0) != 2 {
		t.Fatalf("Expected 2, got: %d", *l.ElementAt(0))
	}
	if *l.ElementAt(1) != 1 {
		t.Fatalf("Expected 1, got: %d", *l.ElementAt(1))
	}
	if *l.ElementAt(2) != 0 {
		t.Fatalf("Expected 0, got: %d", *l.ElementAt(2))
	}
}

func TestReverseSingle(t *testing.T) {
	l := Init[int]()

	l.Push(0)
	l.Reverse()

	if *l.ElementAt(0) != 0 {
		t.Fatalf("Expected 0, got: %d", *l.ElementAt(0))
	}
}

func TestPop(t *testing.T) {
	l := Init[int]()

	l.Push(0)
	l.EmplaceBack(2)
	l.AddAt(1, 1)

	l.Pop()

	if *l.ElementAt(0) != 1 {
		t.Fatalf("After pop expected 1, got: %d\n", *l.ElementAt(0))
	}
}

func TestRemoveAt(t *testing.T) {
	l := Init[int]()

	l.Push(0)
	l.EmplaceBack(2)
	l.AddAt(1, 1)

	l.RemoveAt(1)
	if *l.ElementAt(1) != 2 {
		t.Fatalf("After pop expected 2, got: %d\n", *l.ElementAt(0))
	}
}

func TestAddAt0(t *testing.T) {
	l := Init[int]()

	l.AddAt(0, 0)

	if l.length != 1 {
		t.Fatalf("Was expecting length 1, got: %d\n", l.length)
	}

	if *l.ElementAt(0) != 0 {
		t.Fatalf("Expected 0, got: %d", *l.ElementAt(0))
	}
}
