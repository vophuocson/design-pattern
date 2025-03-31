package state

import tcpcontext "design-pattern-practice/state/tcp-context"

func mai() {
	tcp := tcpcontext.CreateTPC()
	tcp.SetState("established")
	tcp.Open()
	tcp.SetState("close")
	tcp.Open()
}
