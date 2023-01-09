package usecase

import (
	"context"
	"errors"

	"github.com/Hulhay/goldfish/repository"
	"github.com/Hulhay/goldfish/usecase/profile"
)

type profileUC struct {
	userRepo repository.UserRepository
}

type Profile interface {
	GetProfile(ctx context.Context, email string) (*profile.ProfileResponse, error)
}

func NewProfileUC(ur repository.UserRepository) Profile {
	return &profileUC{
		userRepo: ur,
	}
}

func (u *profileUC) GetProfile(ctx context.Context, email string) (*profile.ProfileResponse, error) {

	user, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, errors.New("email not found")
	}

	res := &profile.ProfileResponse{
		ID:           user.ID,
		UserName:     user.UserName,
		UserUsername: user.UserUsername,
		UserEmail:    user.UserEmail,
		UserRole:     user.UserRole,
	}

	return res, nil
}
