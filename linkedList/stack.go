package main

import "fmt"

// 通过链表实现栈

type StackNode struct {
	Value interface{}
	Next  *StackNode
}

type Stack struct {
	Size     uint
	RootNode *StackNode
}

func NewStack() *Stack {
	return &Stack{Size: 0, RootNode: &StackNode{Value: 0, Next: nil}}
}

func (s *Stack) push(v interface{}) {
	if s.RootNode == nil { // 第一次
		s.RootNode.Next = &StackNode{Next: nil, Value: v}
		return
	}
	next := s.RootNode.Next
	s.RootNode.Next = &StackNode{Next: next, Value: v}
}

func (s *Stack) pop() interface{} {
	var node *StackNode = s.RootNode.Next
	if node == nil {
		return nil
	}
	s.RootNode.Next = node.Next
	return node.Value
}

func bracketMatching(v []rune, symbol []rune) bool {
	stack := NewStack()
	for _, r := range v {
		if r == symbol[0] {
			stack.push(symbol[0])
		}
		if r == symbol[1] {
			if stack.pop() == nil {
				return false
			}
		}
	}
	return stack.pop() == nil
}

func main() {
	stack := NewStack()
	stack.push("0")
	stack.push("1")
	stack.push("2")
	stack.push("3")
	stack.push("4")
	stack.push(5)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())

	//testStr := "(上海)(YNN)"
	//testStr2 := "(上海(YNN)"
	testStr3 := "(上海(闵行区)(吴泾镇))"
	//testStr4 := `Hello World`
	runes := []rune(testStr3)
	symbol := []rune("()")
	//symbol := []rune(`""`)
	fmt.Println(bracketMatching(runes, symbol))
}
