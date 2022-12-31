package category

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type InsertCategoryRequest struct {
	CategoryName string `json:"category_name"`
}

func (c *InsertCategoryRequest) Validate() error {

	if err := validation.Validate(c.CategoryName, validation.Required); err != nil {
		return errors.New("category_name must be filled")
	}

	if len(c.CategoryName) < 3 {
		return errors.New("category_name at least has 3 character")
	}

	return nil
}
