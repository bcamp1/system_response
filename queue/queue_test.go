package queue

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	q := Queue{size: 5}
	q.AddNode(2.5)
	q.AddNode(4.3)
	q.AddNode(4.3)
	fmt.Println(q)
}

func TestAddNodeInFront(t *testing.T) {
	q := Queue{size: 5}
	q.AddNodeInFront(17)
	q.AddNodeInFront(33)
	fmt.Println(q)
}

func TestMoveAlong(t *testing.T) {
	q := Queue{size: 3}

	fmt.Printf("Before: %s\n", q)
	q.MoveAlong(0)
	q.MoveAlong(13)
	fmt.Printf("After: %s\n", q)

}
