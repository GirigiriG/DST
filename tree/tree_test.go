package tree

import (
	"testing"

	"github.com/GirigiriG/DST/adt"
)

func TestTreeInsert(t *testing.T) {
	data := []int{5, 3, 7, 6}
	root := &adt.Node{Value: 5}

	for _, v := range data[1:] {
		Insert(root, v)
	}
}

func TestPreOrder(t *testing.T) {
	data := []int{5, 3, 4, 7, 6, 9}
	root := &adt.Node{Value: 5}

	for _, v := range data[1:] {
		Insert(root, v)
	}

	PreOrder(root)
}
