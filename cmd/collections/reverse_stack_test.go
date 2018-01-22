package collections

import (
	"testing"
	"fmt"
)

func TestReverseStack(t *testing.T) {
	stack := &Stack{}
	stack.Push("A")
	stack.Push("B")
	stack.Push("C")
	stack.Push("D")
	fmt.Println(stack)
	ReverseStack(stack)
	fmt.Print(stack)
}
