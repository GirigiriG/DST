package tree

type node struct {
	value int
	left  *node
	right *node
}

func Insert(root *node, data int) *node {
	if root == nil {
		return &node{value: data}
	}

	if data < root.value {
		root.left = Insert(root.left, data)
	}

	if data > root.value {
		root.right = Insert(root.right, data)
	}

	return root
}
