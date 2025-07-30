package repositories

import (
	"errors"
	"gin-app/dto"
	"gin-app/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(id uint, userId uint) (*models.Item, error)
	Create(input dto.CreateItemInput, userId uint) (*models.Item, error)
	Update(id uint, input dto.UpdateItemInput, userId uint) (*models.Item, error)
	Delete(id uint, userId uint) error
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(id uint, userId uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == id && item.UserID == userId {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemMemoryRepository) Create(input dto.CreateItemInput, userId uint) (*models.Item, error) {
	newItem := models.Item{
		// ID:          uint(len(r.items) + 1),
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		UserID:      userId,
	}
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(id uint, input dto.UpdateItemInput, userId uint) (*models.Item, error) {
	for i, item := range r.items {
		if item.ID == id && item.UserID == userId {
			// ポインタの値を安全に取得
			name := item.Name
			if input.Name != nil {
				name = *input.Name
			}

			price := item.Price
			if input.Price != nil {
				price = *input.Price
			}

			description := item.Description
			if input.Description != nil {
				description = *input.Description
			}

			soldOut := item.SoldOut
			if input.SoldOut != nil {
				soldOut = *input.SoldOut
			}

			r.items[i] = models.Item{
				// ID:          id,
				Name:        name,
				Price:       price,
				Description: description,
				SoldOut:     soldOut,
				UserID:      userId,
			}
			return &r.items[i], nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemMemoryRepository) Delete(id uint, userId uint) error {
	for i, item := range r.items {
		if item.ID == id && item.UserID == userId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(input dto.CreateItemInput, userId uint) (*models.Item, error) {
	newItem := models.Item{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		UserID:      userId,
	}
	if err := r.db.Create(&newItem).Error; err != nil {
		return nil, err
	}
	return &newItem, nil
}

func (r *ItemRepository) Update(id uint, input dto.UpdateItemInput, userId uint) (*models.Item, error) {
	var item models.Item
	if err := r.db.First(&item, "id = ? AND user_id = ?", id, userId).Error; err != nil {
		return nil, err
	}
	if input.Name != nil {
		item.Name = *input.Name
	}
	if input.Price != nil {
		item.Price = *input.Price
	}
	if input.Description != nil {
		item.Description = *input.Description
	}
	if input.SoldOut != nil {
		item.SoldOut = *input.SoldOut
	}
	if err := r.db.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemRepository) Delete(id uint, userId uint) error {
	result := r.db.Delete(&models.Item{}, "id = ? AND user_id = ?", id, userId)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("item not found")
	}
	return nil
}

func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return &items, nil
}

func (r *ItemRepository) FindById(id uint, userId uint) (*models.Item, error) {
	var item models.Item
	if err := r.db.First(&item, "id = ? AND user_id = ?", id, userId).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
