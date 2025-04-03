package memento

type MementoNode struct {
	memento *EditorMemento
	next    *MementoNode
}

type MementoStack struct {
	current *MementoNode
}

func (s *MementoStack) Push(memento *EditorMemento) {
	m := MementoNode{memento: memento, next: s.current}
	s.current = &m
}

func (s *MementoStack) Pop() *EditorMemento {
	if s.current == nil {
		return nil
	}
	current := s.current
	s.current = s.current.next
	return current.memento
}
