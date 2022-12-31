package member

import (
	"errors"

	"github.com/Hulhay/goldfish/shared"
	validation "github.com/go-ozzo/ozzo-validation"
)

type EditMemberRequest struct {
	MemberID   int
	MemberNIK  string `json:"member_nik"`
	MemberName string `json:"member_name"`
}

func (c *EditMemberRequest) Validate() error {

	if err := validation.Validate(c.MemberNIK, validation.Required); err != nil {
		return errors.New("nik must be filled")
	}

	if !shared.IsNIKFormat(c.MemberNIK) {
		return errors.New("invalid nik")
	}

	if err := validation.Validate(c.MemberName, validation.Required); err != nil {
		return errors.New("name must be filled")
	}

	return nil
}
