package flyweight

import (
	"design-pattern-practice/structural-pattern/flyweight/actor"
	"design-pattern-practice/structural-pattern/flyweight/message"

	"github.com/google/uuid"
)

func main() {
	actors := actor.NewActorFactory()
	idSon1 := uuid.New()
	idSon2 := uuid.New()
	idSon3 := uuid.New()
	actors.AddActor(&actor.ActorDisplay{
		ID:   idSon1,
		Name: "son 1",
		Icon: "icon son 1",
	})
	actors.AddActor(&actor.ActorDisplay{
		ID:   idSon2,
		Name: "son 2",
		Icon: "icon son 2",
	})
	actors.AddActor(&actor.ActorDisplay{
		ID:   idSon3,
		Name: "son 3",
		Icon: "icon son 3",
	})

	var messages []message.Message
	messages = append(messages, &message.TextMessage{Actor: actors.RangeActor[idSon1], Content: "this is content message 1"})
	messages = append(messages, &message.ImageMessage{Actor: actors.RangeActor[idSon2], Base64: "this is base 64 image"})
	messages = append(messages, &message.TextMessage{Actor: actors.RangeActor[idSon3], Content: "this is content message 2"})

	for _, message := range messages {
		message.DisplayMessage()
	}

}
