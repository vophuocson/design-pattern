package observer

import (
	"design-pattern-practice/observer/observer"
	"design-pattern-practice/observer/subject"
)

func mani() {
	user1 := observer.NewObserver()
	user2 := observer.NewObserver()
	user3 := observer.NewObserver()
	user4 := observer.NewObserver()

	sub := subject.NewSubject()
	sub.Register(user1)
	sub.Register(user2)
	sub.Register(user3)
	sub.Register(user4)

	sub.ChangeContent("this is the content that is changed")
	sub.NotifyAll()

}
