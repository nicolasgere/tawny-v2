package front

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func GetTopicFromMetadata(ctx context.Context) (topic string, ok bool) {
	var m metadata.MD
	m, ok = metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}
	t := m.Get("topic")
	if len(t) != 1 {
		ok = false
		return
	}
	topic = t[0]
	ok = true
	return
}
