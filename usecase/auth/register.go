package auth

import (
	"errors"

	"github.com/Hulhay/goldfish/model"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterRequest struct {
	UserName     string `json:"user_name"`
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserRole     string `json:"user_role"`
}

func (c *RegisterRequest) Validate() error {

	if err := validation.Validate(c.UserName, validation.Required); err != nil {
		return errors.New("name must be filled")
	}

	if err := validation.Validate(c.UserEmail, validation.Required); err != nil {
		return errors.New("email must be filled")
	}

	if err := validation.Validate(c.UserRole, validation.Required); err != nil {
		return errors.New("role must be filled")
	}

	if err := validation.Validate(c.UserEmail, is.Email); err != nil {
		return errors.New("invalid email format")
	}

	if err := validation.Validate(c.UserPassword, validation.Required); err != nil {
		return errors.New("password must be filled")
	}

	if err := validation.Validate(c.UserPassword, validation.Length(6, 0)); err != nil {
		return errors.New("password minimal 6 character")
	}

	if !model.IsValidRole[c.UserRole] {
		return errors.New("invalid role")
	}

	return nil
}
