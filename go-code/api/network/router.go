package network

type Router struct {
	*Network
}

func newRouter(n *Network) {
	s := &Router{n}

	n.Router(GET, "/send", s.send)
	n.Router(GET, "/send-with-tag", s.sendWithTag)
	n.Router(GET, "/send-with-child", s.sendWithChild)

	n.Router(GET, "/receive-from-other-host", s.sendWithOtherHost)
	n.Router(GET, "/receive-one-from-other-host", s.receiveOne)
	n.Router(GET, "/receive-two-from-other-host", s.receiveTwo)

	n.Router(GET, "/send-for-panic", s.sendForPanic)
	n.Router(GET, "/receive-for-error", s.receiveForError)

	n.Router(GET, "/send-for-baggage", nil)
	n.Router(GET, "/receive-for-baggage", nil)

}
