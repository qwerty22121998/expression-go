package tree

type StackNode struct {
	Val  *Node
	Prev *StackNode
}

type Stack struct {
	Tail *StackNode
}

func (s *Stack) Push(node *Node) {
	sNode := &StackNode{
		Val:  node,
		Prev: s.Tail,
	}
	s.Tail = sNode
}

func (s *Stack) Pop() *StackNode {
	now := s.Tail
	if now == nil {
		return nil
	}
	s.Tail = now.Prev
	return now
}
