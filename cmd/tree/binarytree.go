package tree

import (
	"strconv"
)

type BinaryTree struct {
	val   int
	left  *BinaryTree
	right *BinaryTree
}

func NewBinaryTree(val int) *BinaryTree {
	return &BinaryTree{val: val}
}
func (b BinaryTree) Insert(val int, tree *BinaryTree) *BinaryTree {
	if tree == nil {
		return NewBinaryTree(val)
	}

	if val < tree.val {
		tree.left = b.Insert(val, tree.left)
	} else {

		tree.right = b.Insert(val, tree.right)
	}
	return tree
}

func (b BinaryTree) String() string {
	str := ""
	if b == (BinaryTree{}) {
		return str
	}
	return print(b, str)
}

func print(b BinaryTree, str string) string {
	if b != (BinaryTree{}) {
		if b.left != nil {
			str = print(*b.left, str)
		}
		str = str + strconv.Itoa(b.val)
		if b.right != nil {
			str = print(*b.right, str)
		}
	}
	return str
}
