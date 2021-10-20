package front

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"tawny/lib/configuration"
	"tawny/lib/object"
	pubsub "tawny/protos"

	"github.com/golang/protobuf/ptypes/empty"
)

type PushServer struct {
	pubsub.UnimplementedPubSubServiceServer
	Logger        *zap.Logger
	Configuration *configuration.Configuration
	ObjectService object.Service
}

var c = make(chan *pubsub.Message, 1000)

func (s *PushServer) Subscribe(subscribeInput *pubsub.SubscribeInput, stream pubsub.PubSubService_SubscribeServer) error {
	id := subscribeInput.Id
	topic, ok := GetTopicFromMetadata(stream.Context())
	if !ok {
		return errors.New("topic metadata require")
	}

	durableObject := s.ObjectService.GetDurableObject(topic)
	outputChannel := durableObject.ConnectClient(id)

Loop:
	for {
		select {
		case <-stream.Context().Done():
			{
				durableObject.DisconnectClient(id)
				break Loop
			}
		case m := <-outputChannel:
			err := stream.Send(&pubsub.Message{
				Data: m.Data,
			})
			if err != nil {
				s.Logger.Error("error in broadcasting message", zap.Error(err))
				break Loop
			}
		}
	}
	return nil
}
func (s *PushServer) Publish(ctx context.Context, pushInput *pubsub.PushInput) (*empty.Empty, error) {
	topic, ok := GetTopicFromMetadata(ctx)
	if !ok {
		return nil, errors.New("topic metadata require")
	}
	o := s.ObjectService.GetDurableObject(topic)
	o.Publish(pushInput.Data)
	return &empty.Empty{}, nil
}

//if PORT != "9090" {
//	fmt.Println("PROXYING REQUEST")
//	var conn *grpc.ClientConn
//	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("did not connect: %s", err)
//	}
//	fmt.Println("TATATA")
//
//	c := pubsub.NewPubSubServiceClient(conn)
//	fmt.Println("TOTOTOTO")
//	res, err := c.Subscribe(metadata.NewOutgoingContext(stream.Context(), m), subscribeInput)
//	if err != nil {
//		return err
//	}
//	fmt.Println("GO GO GO")
//
//	for {
//		p := pubsub.Message{}
//		err := res.RecvMsg(&p)
//		fmt.Println("GO GO GO GOO GOGO")
//		fmt.Println(err)
//		if err != nil {
//			break
//		}
//		err = stream.Send(&p)
//		if err != nil {
//			fmt.Println(err)
//			break
//		}
//	}
//	defer conn.Close()
//	return nil
//}
