package member

import (
	"errors"

	"github.com/Hulhay/goldfish/shared"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GetMemberDetailRequest struct {
	MemberNIK string `query:"member_nik"`
}

type MemberDetailResponse struct {
	MemberID     int    `json:"member_id"`
	MemberNIK    string `json:"member_nik"`
	MemberName   string `json:"member_name"`
	FamilyNIK    string `json:"family_nik"`
	FamilyID     string `json:"family_id"`
	MemberIsHead bool   `json:"member_is_head"`
}

func (c *GetMemberDetailRequest) Validate() error {

	if err := validation.Validate(c.MemberNIK, validation.Required); err != nil {
		return errors.New("nik must be filled")
	}

	if !shared.IsNIKFormat(c.MemberNIK) {
		return errors.New("invalid nik")
	}

	return nil
}
