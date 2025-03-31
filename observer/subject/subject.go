package subject

import "design-pattern-practice/observer/observer"

type Subject interface {
	Register(observer.Observer)
	UnRegister(observer.Observer)
	NotifyAll()
	ChangeContent(content string)
}

type New struct {
	observers []observer.Observer
	content   string
}

func (s *New) Register(o observer.Observer) {
	s.observers = append(s.observers, o)
}
func (s *New) UnRegister(o observer.Observer) {
	for i, observer := range s.observers {
		if observer.GetID() == o.GetID() {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *New) NotifyAll() {
	for _, ob := range s.observers {
		ob.Update(s.content)
	}
}
func (s *New) ChangeContent(content string) {
	s.content = content
}

func NewSubject() Subject {
	return &New{}
}
