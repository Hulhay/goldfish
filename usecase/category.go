package usecase

import (
	"context"
	"errors"

	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/repository"
	"github.com/Hulhay/goldfish/shared"
	"github.com/Hulhay/goldfish/usecase/category"
)

type categoryUC struct {
	categoryRepo repository.CategoryRepository
}

type Category interface {
	InsertCategory(ctx context.Context, params category.InsertCategoryRequest) error
	GetListCategory(ctx context.Context) ([]model.Category, error)
}

func NewCategoryUC(cr repository.CategoryRepository) Category {
	return &categoryUC{
		categoryRepo: cr,
	}
}

func (u *categoryUC) InsertCategory(ctx context.Context, params category.InsertCategoryRequest) error {

	var (
		err           error
		categoryExist *model.Category
	)

	if err = params.Validate(); err != nil {
		return err
	}

	categoryExist, _ = u.categoryRepo.GetCategoryByCategoryName(ctx, params.CategoryName)

	if categoryExist != nil {
		return errors.New("category exist")
	}

	categoryValue := shared.CreateValue(params.CategoryName)

	req := &model.Category{
		CategoryName:  params.CategoryName,
		CategoryValue: categoryValue,
	}

	err = u.categoryRepo.InsertCategory(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *categoryUC) GetListCategory(ctx context.Context) ([]model.Category, error) {

	var (
		category []model.Category
		err      error
	)

	category, err = u.categoryRepo.GetListCategory(ctx)
	if err != nil {
		return nil, err
	}

	return category, nil
}
