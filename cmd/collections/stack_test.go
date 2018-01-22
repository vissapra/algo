package collections

import (
	"testing"
	"fmt"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	for i:='A'; i<= 'Z'; i++{
		stack.Push(i)
	}
	if len(stack.data) != 26 {
		t.Errorf("Expected size of 26 characters (alphabets). Has: %d",len(stack.data))
	}
	fmt.Println(stack)
}
