package stack

//Stack represents a stack data structure
type Stack struct {
	nodes []rune
}

func (s *Stack) Push(val rune) {
	s.nodes = append(s.nodes, val)
}

func (s *Stack) IsEmpty() bool {
	return len(s.nodes) == 0
}

func (s *Stack) Pop() rune {
	element := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return element
}
