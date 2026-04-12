package service

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg/logger"
	"flower-shop/storage"
)

type roleService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewRoleService(storage storage.IStorage, logger logger.ILogger) roleService {
	return roleService{
		storage: storage,
		logger:  logger,
	}
}

func (r roleService) Create(ctx context.Context, role models.AddRole) error {
	err := r.storage.Roles().Create(ctx, role)
	if err != nil {
		r.logger.Error("failed to create a new role: ", logger.Error(err))
		return err
	}
	return nil
}

func (r roleService) GetAll(ctx context.Context, req models.GetAllRolesRequest) (models.GetAllRolesResponse, error) {

	resp, err := r.storage.Roles().GetAll(ctx, req)
	if err != nil {
		r.logger.Error("failed to get all roles: ", logger.Error(err))
		return models.GetAllRolesResponse{}, err
	}
	return resp, nil
}

func (r roleService) Delete(ctx context.Context, ID string) error {

	err := r.storage.Roles().Delete(ctx, ID)
	if err != nil {
		r.logger.Error("failed to delete a role: ", logger.Error(err))
		return err
	}
	return nil
}

func (r roleService) Update(ctx context.Context, req models.UpdateRole) error {

	err := r.storage.Roles().Update(ctx, req)
	if err != nil {
		r.logger.Error("failed to update an role's info: ", logger.Error(err))
		return err
	}
	return nil
}
