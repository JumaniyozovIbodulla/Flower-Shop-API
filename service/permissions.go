package service

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg/logger"
	"flower-shop/storage"
)

type permissionService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewPermissionService(storage storage.IStorage, logger logger.ILogger) permissionService {
	return permissionService{
		storage: storage,
		logger:  logger,
	}
}

func (p permissionService) Create(ctx context.Context, permission models.AddPermission) error {

	err := p.storage.Permissions().Create(ctx, permission)
	if err != nil {
		p.logger.Error("failed to create a new permission: ", logger.Error(err))
		return err
	}
	return nil
}

func (p permissionService) GetAll(ctx context.Context, req models.GetAllPermissionsRequest) (models.GetAllPermissionsResponse, error) {

	resp, err := p.storage.Permissions().GetAll(ctx, req)
	if err != nil {
		p.logger.Error("failed to get all permissions: ", logger.Error(err))
		return models.GetAllPermissionsResponse{}, err
	}
	return resp, nil
}

func (p permissionService) Delete(ctx context.Context, ID string) error {

	err := p.storage.Permissions().Delete(ctx, ID)
	if err != nil {
		p.logger.Error("failed to delete a role: ", logger.Error(err))
		return err
	}
	return nil
}

func (p permissionService) Update(ctx context.Context, req models.UpdatePermission) error {

	err := p.storage.Permissions().Update(ctx, req)
	if err != nil {
		p.logger.Error("failed to update an permission's info: ", logger.Error(err))
		return err
	}
	return nil
}
