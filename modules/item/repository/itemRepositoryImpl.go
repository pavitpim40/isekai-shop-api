package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemEntity "github.com/Rayato159/isekai-shop-api/modules/item/entity"
	_itemException "github.com/Rayato159/isekai-shop-api/modules/item/exception"
)

type itemRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemRepository {
	return &itemRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *itemRepositoryImpl) FindItems(itemFilterDto *_itemEntity.ItemFilterDto) ([]*_itemEntity.Item, error) {
	query := r.db.Model(&_itemEntity.Item{})
	if itemFilterDto.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilterDto.Name+"%")
	}
	if itemFilterDto.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilterDto.Description+"%")
	}

	offset := int((itemFilterDto.Page - 1) * itemFilterDto.Size)
	size := int(itemFilterDto.Size)

	items := make([]*_itemEntity.Item, 0)

	if err := query.Offset(offset).Limit(size).Find(&items).Error; err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, &_itemException.ItemListingException{}
	}

	return items, nil
}

func (r *itemRepositoryImpl) CountItems(itemFilterDto *_itemEntity.ItemFilterDto) (int64, error) {
	query := r.db.Model(&_itemEntity.Item{}).Where("is_archive = ?", false)

	if itemFilterDto.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemFilterDto.Name+"%")
	}
	if itemFilterDto.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemFilterDto.Description+"%")
	}

	var count int64

	if err := query.Count(&count).Error; err != nil {
		r.logger.Error("Failed to count items", err.Error())
		return -1, &_itemException.CountItemsException{}
	}

	return count, nil
}

func (r *itemRepositoryImpl) InsertItem(item *_itemEntity.Item) (*_itemEntity.Item, error) {
	insertedItem := new(_itemEntity.Item)

	if err := r.db.Create(item).Scan(insertedItem).Error; err != nil {
		r.logger.Error("Failed to insert item", err.Error())
		return nil, &_itemException.InsertItemException{}
	}

	return insertedItem, nil
}

func (r *itemRepositoryImpl) UpdateItem(itemID uint64, item *_itemEntity.UpdateItemDto) (*_itemEntity.Item, error) {
	updatedItem := new(_itemEntity.Item)

	if err := r.db.Model(&_itemEntity.Item{}).Where(
		"id = ?", itemID,
	).Updates(
		item,
	).Scan(
		updatedItem,
	).Error; err != nil {
		r.logger.Error("Failed to update item", err.Error())
		return nil, &_itemException.UpdateItemException{}
	}

	return updatedItem, nil
}

func (r *itemRepositoryImpl) ArchiveItem(itemID uint64) error {
	if err := r.db.Table("items").Where(
		"id = ?", itemID,
	).Update(
		"is_archive", true,
	).Error; err != nil {
		r.logger.Error("Failed to archive item", err.Error())
		return &_itemException.ArchiveItemException{ItemID: itemID}
	}

	return nil
}
