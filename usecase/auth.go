package usecase

import (
	"context"
	"errors"

	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/repository"
	"github.com/Hulhay/goldfish/shared"
	"github.com/Hulhay/goldfish/usecase/auth"
)

type authUC struct {
	userRepo repository.UserRepository
}

type Auth interface {
	Register(ctx context.Context, params auth.RegisterRequest) error
	Login(ctx context.Context, params auth.LoginRequest) (*model.User, error)
	Logout(ctx context.Context, email string) error
}

func NewAuthUC(ur repository.UserRepository) Auth {
	return &authUC{
		userRepo: ur,
	}
}

func (u *authUC) Register(ctx context.Context, params auth.RegisterRequest) error {
	var (
		encryptedPassword string
		err               error
		user              *model.User
	)

	if err = params.Validate(); err != nil {
		return err
	}

	user, _ = u.userRepo.GetUserByEmail(ctx, params.UserEmail)

	if user != nil {
		return errors.New("email is used")
	}

	encryptedPassword, err = shared.EncryptPassword(params.UserPassword)
	if err != nil {
		return err
	}

	req := &model.User{
		UserName:     params.UserName,
		UserUsername: params.UserUsername,
		UserEmail:    params.UserEmail,
		UserPassword: encryptedPassword,
		UserRole:     params.UserRole,
		UserIsLogin:  false,
	}

	err = u.userRepo.InsertUser(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *authUC) Login(ctx context.Context, params auth.LoginRequest) (*model.User, error) {

	var (
		err  error
		user *model.User
	)

	if err = params.Validate(); err != nil {
		return nil, err
	}

	user, err = u.userRepo.GetUserByEmail(ctx, params.Email)
	if err != nil || user == nil {
		return nil, errors.New("email not found")
	}

	err = shared.CheckPassword(params.Password, user.UserPassword)
	if err != nil {
		return nil, errors.New("wrong password")
	}

	err = u.userRepo.UpdateStatusLoginTrue(ctx, params.Email)
	if err != nil {
		return nil, err
	}

	user, _ = u.userRepo.GetUserByEmail(ctx, params.Email)

	return user, nil
}

func (u *authUC) Logout(ctx context.Context, email string) error {

	var (
		err  error
		user *model.User
	)

	user, _ = u.userRepo.GetUserByEmail(ctx, email)
	if !user.UserIsLogin {
		return errors.New("you don't have permission. you need to login first")
	}

	err = u.userRepo.UpdateStatusLoginFalse(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
