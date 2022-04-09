package parser

import (
	"github.com/RicardoBeltranRamirez/parser/internal/stack"
)

var bracketsMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

// AreBracketsBalanced checks whether brackets are balanced or not
// 			() -> true
//			{} -> true
//			[] -> true
//			({[a]}) -> true
//			{[] -> false
// Any other character than (), [], {} will be ignored by the parser
func AreBracketsBalanced(val string) bool {
	stackDst := stack.RuneStack{}

	if val == "" {
		return false
	}

	for _, v := range val {

		if v == '(' || v == '[' || v == '{' {
			stackDst.Push(v)
			continue
		}

		if v == ')' || v == '}' || v == ']' {
			//At this point, if the stack is empty, then no opening brackets were found, so we can say the brackets
			//are not balanced
			if stackDst.IsEmpty() {
				return false
			}

			openBracket := stackDst.Pop()
			matchingBracket := bracketsMap[openBracket]

			if v != matchingBracket {
				return false
			}
		}

	}

	return stackDst.IsEmpty()
}
