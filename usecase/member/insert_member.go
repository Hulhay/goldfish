package member

import (
	"errors"

	"github.com/Hulhay/goldfish/shared"
	validation "github.com/go-ozzo/ozzo-validation"
)

type InsertMemberRequest struct {
	MemberNIK    string `json:"member_nik"`
	MemberName   string `json:"member_name"`
	MemberIsHead bool   `json:"member_is_head"`
	FamilyID     string `json:"family_id"`
	FamilyNIK    string `json:"family_nik"`
}

func (c *InsertMemberRequest) Validate() error {

	if err := validation.Validate(c.MemberNIK, validation.Required); err != nil {
		return errors.New("nik must be filled")
	}

	if !shared.IsNIKFormat(c.MemberNIK) {
		return errors.New("invalid nik")
	}

	if err := validation.Validate(c.MemberName, validation.Required); err != nil {
		return errors.New("name must be filled")
	}

	if !c.MemberIsHead {
		if err := validation.Validate(c.FamilyID, validation.Required); err != nil {
			return errors.New("family_id must be filled")
		}
		if c.FamilyNIK != `` {
			return errors.New("invalid request")
		}
	} else {
		if err := validation.Validate(c.FamilyNIK, validation.Required); err != nil {
			return errors.New("nik (keluarga) must be filled")
		}
		if c.FamilyID != `` {
			return errors.New("invalid request")
		}
	}

	return nil
}
