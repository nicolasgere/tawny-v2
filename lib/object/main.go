package object

import (
	"go.uber.org/zap"
	"sync"
	"tawny/lib/configuration"
)

type Service interface {
	GetDurableObject(id string) (d DurableObject)
}

type ServiceImpl struct {
	logger             *zap.Logger
	configuration      configuration.Configuration
	mu                 sync.Mutex
	durableObjectStore map[string]DurableObject
}

func NewObjectService(logger *zap.Logger, configuration configuration.Configuration) Service {
	return ServiceImpl{
		logger:        logger,
		configuration: configuration,
	}
}
