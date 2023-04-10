package queue

import "fmt"

type Node struct {
	next *Node
	val  float64
}

type Queue struct {
	size int
	head *Node
	tail *Node
}

func (q *Queue) AddNode(val float64) {
	n := Node{val: val}
	if q.head == nil {
		q.head = &n
		q.tail = &n
	} else {
		q.tail.next = &n
		q.tail = &n
	}
}

func (q *Queue) AddNodeInFront(val float64) {
	n := Node{val: val}
	tmp := q.head
	q.head = &n
	q.head.next = tmp
}

func (q Queue) String() string {
	s := "["
	n := q.head
	if n == nil {
		return "[]"
	}

	for n != nil {
		s += fmt.Sprintf("%f", n.val)
		if n.next != nil {
			s += ", "
		}

		n = n.next
	}
	return s + "]"
}

func (q *Queue) MoveAlong(val float64) {
	if q.head == nil {
		return
	}

	n := q.head
	for n.next.next != nil {
		n = n.next
	}
	n.next = nil
	q.AddNodeInFront(val)
}
