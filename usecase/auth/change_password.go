package auth

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ChangePasswordRequest struct {
	UserID          int64
	UserEmail       string
	UserOldPassword string `json:"user_old_password"`
	UserNewPassword string `json:"user_new_password"`
}

func (c *ChangePasswordRequest) Validate() error {

	if err := validation.Validate(c.UserOldPassword, validation.Required); err != nil {
		return errors.New("old password must be filled")
	}

	if err := validation.Validate(c.UserNewPassword, validation.Required); err != nil {
		return errors.New("new password must be filled")
	}

	if err := validation.Validate(c.UserOldPassword, validation.Length(6, 0)); err != nil {
		return errors.New("password minimal 6 character")
	}

	if err := validation.Validate(c.UserNewPassword, validation.Length(6, 0)); err != nil {
		return errors.New("password minimal 6 character")
	}

	return nil
}
