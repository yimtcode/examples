package queue

type node struct {
	Data interface{}
	Next *node
}

type Queue struct {
	head *node
	tail *node
	len  int
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Push(obj interface{}) {
	q.len++
	n := &node{Data: obj, Next: nil}
	if q.head == nil {
		q.head = n
	}
	if q.tail != nil {
		q.tail.Next = n
	}
	q.tail = n
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.len == 0 {
		return nil, false
	}

	q.len--
	head := q.head
	q.head = q.head.Next
	if q.tail == head {
		q.tail = nil
	}

	return head.Data, true
}
