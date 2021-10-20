package object

type EventType int32

const (
	CLIENT_DISCONNECTED EventType = 0
	CLIENT_CONNECTED    EventType = 1
)

type MessageIn struct {
	Data []byte
}

type MessageOut struct {
	Data []byte
}

type ClientEvent struct {
	ClientId   string
	ChannelOut *chan *MessageOut
	Type       EventType
}
