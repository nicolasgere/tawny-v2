package object

import (
	"go.uber.org/zap"
	"tawny/lib/configuration"
)

type DurableObject struct {
	Id            string
	logger        *zap.Logger
	configuration *configuration.Configuration
	MessageIn     chan *MessageIn
	ClientEvent   chan *ClientEvent
	clients       map[string]chan *MessageOut
}

func (self *DurableObject) Start() {
	self.clients = map[string]chan *MessageOut{}
	go func() {
		for {
			select {
			case event := <-self.ClientEvent:
				{
					switch event.Type {
					case CLIENT_DISCONNECTED:
						{
							close(self.clients[event.ClientId])
							delete(self.clients, event.ClientId)
						}
					case CLIENT_CONNECTED:
						{
							self.clients[event.ClientId] = *event.ChannelOut
						}
					}
				}
			case c := <-self.MessageIn:
				{
					for _, v := range self.clients {
						v <- &MessageOut{
							Data: c.Data,
						}
					}
				}
			}
		}
	}()
}
func (self *DurableObject) Publish(data []byte) {
	self.MessageIn <- &MessageIn{
		Data: data,
	}
}
func (self *DurableObject) ConnectClient(id string) chan *MessageOut {
	ch := make(chan *MessageOut, 100)
	self.ClientEvent <- &ClientEvent{
		ClientId:   id,
		Type:       CLIENT_CONNECTED,
		ChannelOut: &ch,
	}
	return ch
}a
func (self *DurableObject) DisconnectClient(id string) {
	self.ClientEvent <- &ClientEvent{
		ClientId: id,
		Type:     CLIENT_DISCONNECTED,
	}
}
func NewDurableObject(id string, logger *zap.Logger, configuration *configuration.Configuration) DurableObject {
	d := DurableObject{
		Id:            id,
		logger:        logger,
		configuration: configuration,
		MessageIn:     make(chan *MessageIn, 1000),
		ClientEvent:   make(chan *ClientEvent, 1000),
	}
	d.Start()
	return d
}
