package stack

//RuneStack represents a stack data structure
type RuneStack struct {
	nodes []rune
}

func (s *RuneStack) Push(val rune) {
	s.nodes = append(s.nodes, val)
}

func (s *RuneStack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *RuneStack) Len() int {
	return len(s.nodes)
}
func (s *RuneStack) Pop() rune {
	element := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return element
}
