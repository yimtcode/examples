package queue

import "testing"

func TestQueue(t *testing.T) {
	q := new(Queue)
	q.Push(1)

	if q.Len() != 1 {
		t.Fatalf("Len() != 1")
	}


	num, ok := q.Pop()
	if !ok || num != 1 {
		t.Fatalf("Pop() != 1")
	}

}