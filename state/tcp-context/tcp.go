package tcpcontext

import (
	"design-pattern-practice/state/state"
)

type TCP interface {
	SetState(s string)
	Open()
	Acknowledge()
	Close()
}

type tcpContext struct {
	stateInternal string
	st            state.TcpState
}

func (t *tcpContext) Open() {
	t.st.Open()
}

func (t *tcpContext) Close() {
	t.st.Close()
}

func (t *tcpContext) Acknowledge() {
	t.st.Acknowledge()
}

func (t *tcpContext) SetState(s string) {
	if t.stateInternal != s {
		t.stateInternal = s
		t.st = state.CreateTPCState(s)
	}
}

func CreateTPC() TCP {
	return &tcpContext{}
}
