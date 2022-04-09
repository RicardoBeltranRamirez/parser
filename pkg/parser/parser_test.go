package parser

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBracketsAreBalanced(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{
			"{a}",
			true,
		},
		{
			"([])",
			true,
		},
		{
			"([]{}())",
			true,
		},
		{
			"(){}()",
			true,
		},
		{
			"[{a}(a)]",
			true,
		},
		{
			"{[(a)(a)]}",
			true,
		},
		{
			"",
			false,
		},
		{
			"{[(a)(a)]}b",
			true,
		},
		{
			"b{[()()]}",
			true,
		},
	}

	for _, testCase := range cases {
		result := AreBracketsBalanced(testCase.input)
		require.Equal(t, result, testCase.expected, fmt.Sprintf("input:%s, expected:%v, got:%v", testCase.input, testCase.expected, result))
	}
}
