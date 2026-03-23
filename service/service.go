package service

import (
	"flower-shop/pkg/logger"
	"flower-shop/storage"
)

type IServiceManager interface {
	User() userService
	Role() roleService
}

type Service struct {
	userService userService
	roleService roleService
	logger      logger.ILogger
}

func New(storage storage.IStorage, logger logger.ILogger) Service {
	services := Service{}
	services.userService = NewUserService(storage, logger)
	services.roleService = NewRoleService(storage, logger)
	services.logger = logger

	return services
}

func (n Service) User() userService {
	return n.userService
}

func (n Service) Role() roleService {
	return n.roleService
}
