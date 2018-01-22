package tree

import (
	"testing"
	"fmt"
)

func TestBinaryTree_Insert(t *testing.T) {
	bTree := NewBinaryTree(5)
	bTree.Insert(6,bTree)
	bTree.Insert(1,bTree)
	bTree.Insert(7,bTree)
	bTree.Insert(2,bTree)
	bTree.Insert(3,bTree)
	fmt.Println(bTree)
}
