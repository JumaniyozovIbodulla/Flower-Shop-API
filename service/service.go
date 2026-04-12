package service

import (
	"flower-shop/pkg/logger"
	"flower-shop/storage"
)

type IServiceManager interface {
	User() userService
	Role() roleService
	Permission() permissionService
	RolePermissions() rolePermissionsService
	UserRoles() userRolesService
	Flowers() flowerService
}

type Service struct {
	userService            userService
	roleService            roleService
	permissionService      permissionService
	rolePermissionsService rolePermissionsService
	userRolesService       userRolesService
	flowerService          flowerService
	logger                 logger.ILogger
}

func New(storage storage.IStorage, logger logger.ILogger) Service {
	services := Service{}
	services.userService = NewUserService(storage, logger)
	services.roleService = NewRoleService(storage, logger)
	services.permissionService = NewPermissionService(storage, logger)
	services.rolePermissionsService = NewRolePermissionsService(storage, logger)
	services.userRolesService = NewUserRolesService(storage, logger)
	services.flowerService = NewFlowerService(storage, logger)
	services.logger = logger

	return services
}

func (n Service) User() userService {
	return n.userService
}

func (n Service) Role() roleService {
	return n.roleService
}

func (p Service) Permission() permissionService {
	return p.permissionService
}

func (s Service) RolePermissions() rolePermissionsService {
	return s.rolePermissionsService
}

func (s Service) UserRoles() userRolesService {
	return s.userRolesService
}

func (s Service) Flowers() flowerService {
	return s.flowerService
}
