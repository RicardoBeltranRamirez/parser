package parser

import "testing"

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
			t.Fatalf("input:%s, expected:%v, got:%v", testCase.input, testCase.expected, result)
		}
	}
}
