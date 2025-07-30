package services

import (
	"gin-app/dto"
	"gin-app/models"
	"gin-app/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(id uint, userId uint) (*models.Item, error)
	Create(input dto.CreateItemInput, userId uint) (*models.Item, error)
	Update(id uint, input dto.UpdateItemInput, userId uint) (*models.Item, error)
	Delete(id uint, userId uint) error
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(id uint, userId uint) (*models.Item, error) {
	return s.repository.FindById(id, userId)
}

func (s *ItemService) Create(input dto.CreateItemInput, userId uint) (*models.Item, error) {
	return s.repository.Create(input, userId)
}

func (s *ItemService) Update(id uint, input dto.UpdateItemInput, userId uint) (*models.Item, error) {
	return s.repository.Update(id, input, userId)
}

func (s *ItemService) Delete(id uint, userId uint) error {
	return s.repository.Delete(id, userId)
}
