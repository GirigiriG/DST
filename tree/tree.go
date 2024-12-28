package tree

import (
	"fmt"

	"github.com/GirigiriG/DST/adt"
	"github.com/GirigiriG/DST/queue"
)

func Insert(root *adt.Node, data int) *adt.Node {
	if root == nil {
		return &adt.Node{Value: data}
	}

	if data < root.Value {
		root.Left = Insert(root.Left, data)
	}

	if data > root.Value {
		root.Right = Insert(root.Right, data)
	}

	return root
}

func PreOrder(root *adt.Node) {
	q := queue.NewQueue()
	q.Enqueue(root)

	for !q.IsEmpty() {
		curr, _ := q.Dequeue()

		if curr.Left != nil {
			q.Enqueue(curr.Left)
		}

		if curr.Right != nil {
			q.Enqueue(curr.Right)
		}
		fmt.Println(curr.Value)
	}
}
