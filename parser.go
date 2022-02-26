package parser

import (
	"github.com/RicardoBeltranRamirez/parser/stack"
)

func AreBracketsBalanced(val string) bool {
	stackDst := stack.Stack{}

	for _, v := range val {

		if v == '(' || v == '[' || v == '{' {
			stackDst.Push(v)
			continue
		}

		//At this point, if the stack is empty, then no opening bracket were found, so we can say the brackets
		//are not balanced
		if stackDst.IsEmpty() {
			return false
		}

		switch v {
		case ')':
			check := stackDst.Pop()
			if check == '[' || check == '{' {
				return false
			}
			break

		case '}':
			check := stackDst.Pop()
			if check == '(' || check == '[' {
				return false
			}
			break

		case ']':
			check := stackDst.Pop()
			if check == '(' || check == '{' {
				return false
			}
			break
		}

	}

	return stackDst.IsEmpty()
}
