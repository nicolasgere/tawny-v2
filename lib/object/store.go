package object

import (
	"go.uber.org/zap"
)

func (self ServiceImpl) GetDurableObject(id string) (d DurableObject) {
	d, ok := self.durableObjectStore[id]
	if !ok {
		childLogger := self.logger.With(zap.String("id", id))
		d = NewDurableObject(id, childLogger, &self.configuration)
		self.mu.Lock()
		self.durableObjectStore[id] = d
		self.mu.Unlock()
	}
	return d
}
