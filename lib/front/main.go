package front

import (
	"context"
	"fmt"
	pubsub "tawny/protos"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-uuid"
)

type PushServer struct {
	pubsub.UnimplementedPubSubServiceServer
}

var c = make(chan *pubsub.Message, 1000)

var mapStore = map[string]DurableObject{}

func (s *PushServer) Subscribe(subscribeInput *pubsub.SubscribeInput, stream pubsub.PubSubService_SubscribeServer) error {
	id, _ := uuid.GenerateUUID()
	store, ok := mapStore[subscribeInput.Topic]
	if ok {

	} else {
		store = NewStore(subscribeInput.Topic)
		mapStore[subscribeInput.Topic] = store
	}
	c := make(chan *MessageOut, 1000)
	store.NewClient(id, c)

Loop:
	for {
		select {
		case <-stream.Context().Done():
			{
				fmt.Println("OKOK OK I AM DONE")
				break
			}
		case m := <-c:
			err := stream.Send(&pubsub.Message{
				Data: m.Data,
			})
			if err != nil {
				fmt.Println(err)
				break Loop
			}
		}
	}
	return nil
}
func (s *PushServer) Publish(ctx context.Context, pushInput *pubsub.PushInput) (*empty.Empty, error) {
	store, ok := mapStore[pushInput.Topic]
	if ok {

	} else {
		store = NewStore(pushInput.Topic)
		mapStore[pushInput.Topic] = store
	}
	store.MessageIn <- &MessageIn{
		Data: pushInput.Data,
	}
	return &empty.Empty{}, nil
}

type MessageIn struct {
	Data []byte
}

type MessageOut struct {
	Data []byte
}

type ClientEvent struct {
	Id string
	Type string
}

type DurableObject struct {
	Id        string
	MessageIn chan *MessageIn
	ClientEvent chan *MessageIn
	clients   map[string]chan *MessageOut
}

func (self *DurableObject) Query(name string, params map[string]string) (result interface{}, err error) {
	return
}
func (self *DurableObject) NewClient(name string, channelOut chan *MessageOut) {
	self.clients[name] = channelOut
	return
}

func (self *DurableObject) Start() {
	self.clients = map[string]chan *MessageOut{}
	go func() {
		for {
			select {
			case
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

func NewStore(id string) DurableObject {
	d := DurableObject{
		Id:        id,
		MessageIn: make(chan *MessageIn, 1000),
	}
	d.Start()
	return d
}
