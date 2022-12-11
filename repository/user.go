package repository

import (
	"context"

	"github.com/Hulhay/goldfish/model"
	"gorm.io/gorm"
)

type userRepository struct {
	qry *gorm.DB
}

type UserRepository interface {
	InsertUser(ctx context.Context, params *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateStatusLoginTrue(ctx context.Context, email string) error
	UpdateStatusLoginFalse(ctx context.Context, email string) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		qry: db,
	}
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User

	if err := r.qry.Model(&user).Where("user_email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) InsertUser(ctx context.Context, params *model.User) error {
	var user *model.User

	if err := r.qry.Model(&user).Create(map[string]interface{}{
		"user_name":     params.Name,
		"user_email":    params.Email,
		"user_password": params.Password,
		"user_role":     params.Role,
		"user_is_login": params.IsLogin,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateStatusLoginTrue(ctx context.Context, email string) error {
	var user *model.User

	if err := r.qry.Model(&user).Where("user_email = ?", email).Updates(map[string]interface{}{
		"user_is_login": true,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateStatusLoginFalse(ctx context.Context, email string) error {
	var user *model.User

	if err := r.qry.Model(&user).Where("user_email = ?", email).Updates(map[string]interface{}{
		"user_is_login": false,
	}).Error; err != nil {
		return err
	}

	return nil
}
