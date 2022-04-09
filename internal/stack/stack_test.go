package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRuneStack_Push(t *testing.T) {

	stacktDst := RuneStack{}

	stacktDst.Push('a')
	stacktDst.Push('b')

	require.Equal(t, false, stacktDst.IsEmpty())
	require.Equal(t, 2, stacktDst.Len())
}

func TestRuneStack_Pop(t *testing.T) {
	stacktDst := RuneStack{}

	stacktDst.Push('a')
	stacktDst.Push('b')

	popedValue := stacktDst.Pop()

	require.Equal(t, 'b', popedValue)
	require.Equal(t, false, stacktDst.IsEmpty())
	require.Equal(t, 1, stacktDst.Len())
}
