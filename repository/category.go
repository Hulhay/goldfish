package repository

import (
	"context"

	"github.com/Hulhay/goldfish/model"
	"gorm.io/gorm"
)

type categoryRepository struct {
	qry *gorm.DB
}

type CategoryRepository interface {
	InsertCategory(ctx context.Context, params *model.Category) error
	GetCategoryByCategoryName(ctx context.Context, categoryName string) (*model.Category, error)
	GetListCategory(ctx context.Context) ([]model.Category, error)
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		qry: db,
	}
}

func (r *categoryRepository) InsertCategory(ctx context.Context, params *model.Category) error {
	var category *model.Category

	if err := r.qry.Model(&category).Create(map[string]interface{}{
		"category_name": params.CategoryName,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *categoryRepository) GetCategoryByCategoryName(ctx context.Context, categoryName string) (*model.Category, error) {
	var category *model.Category

	if err := r.qry.Model(&category).Where("category_name = ?", categoryName).First(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) GetListCategory(ctx context.Context) ([]model.Category, error) {
	var res []model.Category

	db := r.qry.Model(&model.Category{})

	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
