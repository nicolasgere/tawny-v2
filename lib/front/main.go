package front

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pubsub "tawny/protos"
)

type PushServer struct {
	pubsub.UnimplementedPubSubServiceServer
}

func (s *PushServer) Subscribe(subscribeInput *pubsub.SubscribeInput, stream pubsub.PubSubService_SubscribeServer) error {
	return nil
}
func (s *PushServer) Publish(ctx context.Context, pushInput *pubsub.PushInput) (*empty.Empty, error) {
	return nil, nil
}
