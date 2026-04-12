package service

import (
	"context"
	"flower-shop/api/models"
	"flower-shop/pkg/logger"
	"flower-shop/storage"
)

type flowerService struct {
	storage storage.IStorage
	logger  logger.ILogger
}


func NewFlowerService(storage storage.IStorage, logger logger.ILogger) flowerService {
	return flowerService{
		storage: storage,
		logger: logger,
	}
}

func (f flowerService) Create(ctx context.Context, flower models.AddFlower) error {
	err := f.storage.Flowers().Create(ctx, flower)
	if err != nil {
		f.logger.Error("failed to create a new flower: ", logger.Error(err))
		return err
	}
	return nil
}

func (f flowerService) Update(ctx context.Context, flower models.UpdateFlower) error {
	err := f.storage.Flowers().Update(ctx, flower)
	if err != nil {
		f.logger.Error("failed to update the flower: ", logger.Error(err))
		return err
	}
	return nil
}

func (f flowerService) Delete(ctx context.Context, ID string) error {
	err := f.storage.Flowers().Delete(ctx, ID)
	if err != nil {
		f.logger.Error("failed to delete the flower: ", logger.Error(err))
		return err
	}
	return nil
}

func (f flowerService) Get(ctx context.Context, ID string) (models.Flower, error) {
	resp, err := f.storage.Flowers().Get(ctx, ID)
	if err != nil {
		f.logger.Error("failed to get the flower: ", logger.Error(err))
		return models.Flower{}, err
	}
	return resp, nil
}

func (f flowerService) GetAll(ctx context.Context, req models.GetAllFlowersRequest) (models.GetAllFlowersResponse, error) {
	resp, err := f.storage.Flowers().GetAll(ctx, req)
	if err != nil {
		f.logger.Error("failed to get the flowers: ", logger.Error(err))
		return models.GetAllFlowersResponse{}, err
	}
	return resp, nil
}

