package queue

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	q := Queue{Size: 5}
	q.AddNode(2.5)
	q.AddNode(4.3)
	q.AddNode(4.3)
	fmt.Println(q)
}

func TestAddNodeInFront(t *testing.T) {
	q := Queue{Size: 5}
	q.AddNodeInFront(17)
	q.AddNodeInFront(33)
	fmt.Println(q)
}

func TestMoveAlong(t *testing.T) {
	q := Queue{Size: 3}
	q.AddNode(3)
	//q.AddNode(4)
	//q.AddNode(5)
	fmt.Printf("Before: %s\n", q)
	q.MoveAlong(22)
	fmt.Printf("After: %s\n", q)
	q.MoveAlong(33)
	fmt.Printf("After: %s\n", q)
	q.MoveAlong(44)
	fmt.Printf("After: %s\n", q)

}
