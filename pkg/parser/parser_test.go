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
			"([)]",
			false,
		},
		{
			"{[}",
			false,
		},
		{
			"[{a}(a)}]",
			false,
		},
		{
			"[{{a}(a)]}",
			false,
		},
		{
			"a[]{}",
			false,
		},
	}

	for _, testCase := range cases {
		result := AreBracketsBalanced(testCase.input)

		if result != testCase.expected {
			require.FailNow(t, fmt.Sprintf("input:%s, expected:%v, got:%v", testCase.input, testCase.expected, result))
		}
	}
}
