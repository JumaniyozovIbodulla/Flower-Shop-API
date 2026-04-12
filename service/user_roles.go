package service

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg/logger"
	"flower-shop/storage"
)

type userRolesService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewUserRolesService(storage storage.IStorage, logger logger.ILogger) userRolesService {
	return userRolesService{
		storage: storage,
		logger:  logger,
	}
}

func (u userRolesService) Create(ctx context.Context, userRole models.UserRole) error {

	err := u.storage.UserRoles().Create(ctx, userRole)
	if err != nil {
		u.logger.Error("failed to create a new user role: ", logger.Error(err))
		return err
	}
	return nil
}

func (u userRolesService) Delete(ctx context.Context, userRole models.UserRole) error {

	err := u.storage.UserRoles().Delete(ctx, userRole)
	if err != nil {
		u.logger.Error("failed to delete the  user role: ", logger.Error(err))
		return err
	}
	return nil
}

func (u userRolesService) Update(ctx context.Context, userRole models.UserRole) error {

	err := u.storage.UserRoles().Update(ctx, userRole)
	if err != nil {
		u.logger.Error("failed to update the  user role: ", logger.Error(err))
		return err
	}
	return nil
}
