package collections

func ReverseStack(stack *Stack){
	if  stack.Len() == 0 {
		return
	}
	val := stack.Pop()
	ReverseStack(stack)
	insertAtBottom(stack, val)
}

func insertAtBottom(stack *Stack, val interface{}) {
	if stack.Len() == 0 {
		stack.Push(val)
		return
	}
	pop := stack.Pop()
	insertAtBottom(stack,val)
	stack.Push(pop)
}
