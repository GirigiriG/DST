package queue

import (
	"fmt"

	"github.com/GirigiriG/DST/adt"
)

type Queue []*adt.Node

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value *adt.Node) {
	*q = append(*q, value)
}

func (q *Queue) Dequeue() (*adt.Node, error) {
	if len(*q) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	node := (*q)[0]
	*q = (*q)[1:]

	return node, nil
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Size() int {
	return len(*q)
}
