package stack

import (
	"testing"
)

func TestStack_Len(t *testing.T) {
	s := new(Stack)
	s.Push(1)

	if s.Len() != 1 {
		t.Errorf("Len() = %v, want %v", s.Len(), 1)
	}
}

func TestStack_Pop(t *testing.T) {
	s := new(Stack)
	s.Push(1)
	s.Push(2)

	obj := s.Pop()
	if obj != 2 {
		t.Errorf("Pop() = %v, want %v", obj, 2)
	}

}
