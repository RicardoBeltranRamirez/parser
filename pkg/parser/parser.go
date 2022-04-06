package parser

import (
	"github.com/RicardoBeltranRamirez/parser/internal/stack"
)

func AreBracketsBalanced(val string) bool {
	stackDst := stack.Stack{}

	for _, v := range val {

		if v == '(' || v == '[' || v == '{' {
			stackDst.Push(v)
			continue
		}

		//At this point, if the stack is empty, then no opening brackets were found, so we can say the brackets
		//are not balanced
		if stackDst.IsEmpty() {
			return false
		}

		switch v {
		case ')':
			openBracket := stackDst.Pop()
			if openBracket == '[' || openBracket == '{' {
				return false
			}
			break

		case '}':
			openBracket := stackDst.Pop()
			if openBracket == '(' || openBracket == '[' {
				return false
			}
			break

		case ']':
			openBracket := stackDst.Pop()
			if openBracket == '(' || openBracket == '{' {
				return false
			}
			break
		}

	}

	return stackDst.IsEmpty()
}
