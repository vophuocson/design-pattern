package state

import "fmt"

type TcpState interface {
	Open()
	Close()
	Acknowledge()
}

type tcpEstablished struct{}

func (s *tcpEstablished) Open() {
	fmt.Println("tcpEstablished Open()")
}

func (s *tcpEstablished) Close() {
	fmt.Println("tcpEstablished Close()")
}

func (s *tcpEstablished) Acknowledge() {
	fmt.Println("tcpEstablished Acknowledge()")
}

type tcpListen struct{}

func (s *tcpListen) Open() {
	fmt.Println("tcpListen Open()")
}

func (s *tcpListen) Close() {
	fmt.Println("tcpListen Close()")
}

func (s *tcpListen) Acknowledge() {
	fmt.Println("tcpListen Acknowledge()")
}

type tcpClose struct{}

func (s *tcpClose) Open() {
	fmt.Println("tcpClose Open()")
}

func (s *tcpClose) Close() {
	fmt.Println("tcpClose Close()")
}

func (s *tcpClose) Acknowledge() {
	fmt.Println("tcpClose Acknowledge()")
}

func CreateTPCState(state string) TcpState {
	switch state {
	case "established":
		return &tcpEstablished{}
	case "close":
		return &tcpClose{}
	default:
		return &tcpListen{}
	}
}
