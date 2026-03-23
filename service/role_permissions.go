package service

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg/logger"
	"flower-shop/storage"
)

type rolePermissionsService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewRolePermissionsService(storage storage.IStorage, logger logger.ILogger) rolePermissionsService {
	return rolePermissionsService{
		storage: storage,
		logger:  logger,
	}
}

func (r rolePermissionsService) Create(ctx context.Context, rolePermission models.RolePermission) error {

	err := r.storage.RolePermissions().Create(ctx, rolePermission)
	if err != nil {
		r.logger.Error("failed to create a new role permission: ", logger.Error(err))
		return err
	}
	return nil
}

func (p rolePermissionsService) Delete(ctx context.Context, rolePermission models.RolePermission) error {

	err := p.storage.RolePermissions().Delete(ctx, rolePermission)
	if err != nil {
		p.logger.Error("failed to delete a permission of the role: ", logger.Error(err))
		return err
	}
	return nil
}

