package actor

import "github.com/google/uuid"

type ActorDisplay struct {
	Icon string
	ID   uuid.UUID
	Name string
}

type ActorFactory struct {
	RangeActor map[uuid.UUID]*ActorDisplay
}

func NewActorFactory() *ActorFactory {
	return &ActorFactory{
		RangeActor: make(map[uuid.UUID]*ActorDisplay),
	}
}

func (f *ActorFactory) GetActor(id uuid.UUID) *ActorDisplay {
	actor, isExists := f.RangeActor[id]
	if isExists {
		return actor
	}
	return nil
}

func (f *ActorFactory) AddActor(actor *ActorDisplay) {
	f.RangeActor[actor.ID] = actor
}
