package queue

import (
	"fmt"
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Node struct {
	next *Node
	val  float64
}

type Queue struct {
	Size int
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

func (q *Queue) MoveAlong(val float64) float64 {
	if q.head == nil {
		q.head = &Node{val: 0}
		return 0
	}

	n := q.head
	prev := n
	count := 1
	for n.next != nil {
		count++
		prev = n
		n = n.next
	}

	if count > q.Size {
		panic("COUNT greater than size")
	}

	if count == q.Size {
		q.tail = prev
		temp := q.tail.next.val
		q.tail.next = nil
		q.AddNodeInFront(val)
		return temp
	}

	q.AddNodeInFront(val)
	return 0
}

func (q Queue) Graph(screen *ebiten.Image, startX, startY, magY, spacingX float32, color color.Color) {
	if q.head == nil {
		return
	}

	n := q.head
	index := 0
	for n.next != nil {
		x1 := startX + float32(index)*spacingX
		x2 := startX + float32(index+1)*spacingX

		y1 := startY - (float32(n.val) * magY)
		y2 := startY - (float32(n.next.val) * magY)

		vector.StrokeLine(screen, x1, y1, x2, y2, 2.0, color, false)

		index++
		n = n.next
	}
}
