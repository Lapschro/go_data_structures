package stack

import "testing"

func TestInit(t *testing.T) {
	s := Init[int]()
	s.Push(5)

	if s.Top() != 5 {
		t.Errorf("Init expecting %d, got %d", 5, s.Top())
	}
}

func TestPush(t *testing.T) {
	s := Init[int]()
	s.Push(5)
	s.Push(10)

	if s.Top() != 10 {
		t.Errorf("After push expected %d, got %d", 10, s.Top())
	}
}

func TestPop(t *testing.T) {
	s := Init[int]()
	s.Push(5)
	s.Push(10)
	s.Pop()

	if s.Top() != 5 {
		t.Errorf("After pop expected %d, got %d", 5, s.Top())
	}

}
