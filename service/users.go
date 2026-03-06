package service

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg/logger"
	"flower-shop/storage"
)

type userService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewUserService(storage storage.IStorage, logger logger.ILogger) userService {
	return userService{
		storage: storage,
		logger:  logger,
	}
}

func (u userService) Create(ctx context.Context, user models.AddUser) error {

	err := u.storage.Users().Create(ctx, user)
	if err != nil {
		u.logger.Error("failed to create a new user: ", logger.Error(err))
		return err
	}
	return nil
}

func (u userService) GetAll(ctx context.Context, req models.GetAllUsersRequest) (models.GetAllUsersResponse, error) {

	resp, err := u.storage.Users().GetAll(ctx, req)
	if err != nil {
		u.logger.Error("failed to get all users: ", logger.Error(err))
		return models.GetAllUsersResponse{}, err
	}
	return resp, nil
}

func (u userService) Delete(ctx context.Context, ID int64) error {

	err := u.storage.Users().Delete(ctx, ID)
	if err != nil {
		u.logger.Error("failed to delete an user: ", logger.Error(err))
		return err
	}
	return nil
}
