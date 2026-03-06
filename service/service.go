package service

import (
	"flower-shop/pkg/logger"
	"flower-shop/storage"
)

type IServiceManager interface {
	User() userService
}

type Service struct {
	userService userService
	logger      logger.ILogger
}

func New(storage storage.IStorage, logger logger.ILogger) Service {
	services := Service{}
	services.userService = NewUserService(storage, logger)
	services.logger = logger

	return services
}

func (n Service) User() userService {
	return n.userService
}
