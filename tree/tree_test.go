package tree

import (
	"fmt"
	"testing"
)

func TestTreeInsert(t *testing.T) {
	data := []int{}
	
	root := &node{value: 5}
	Insert(root, 3)
	Insert(root, 7)
	Insert(root, 6)
	fmt.Println(root.right.left)

}
