package collections

import "fmt"

type Stack struct {
	data []interface{}
}

//Creates a new Stack
func New() *Stack {
	return &Stack{data: make([]interface{}, 0)}
}

//Push to the top of the Stack
func (s *Stack) Push(val interface{}) {
	s.data = append(s.data, val)
}

//Pop from the top of the Stack
func (s *Stack) Pop() interface{} {
	if len(s.data) > 0 {
		val := s.data[len(s.data)-1]
		s.data = s.data[0:(len(s.data) - 1)]
		return val
	}
	return nil
}

func main() {
	stack := New()
	for i := 'A'; i <= 'Z'; i++ {
		stack.Push(i)
	}
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%c\n", stack.Pop())
	}

}
