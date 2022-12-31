package member

import (
	"errors"

	"github.com/Hulhay/goldfish/shared"
)

type GetMemberRequest struct {
	MemberNIK  string `query:"member_nik"`
	MemberName string `query:"member_name"`
	FamilyNIK  string `query:"family_nik"`
	IsHead     bool   `query:"is_head"`
}

type MemberListResponse struct {
	MemberNIK    string `json:"member_nik"`
	MemberName   string `json:"member_name"`
	FamilyNIK    string `json:"family_nik"`
	FamilyID     string `json:"family_id"`
	MemberIsHead bool   `json:"member_is_head"`
}

func (c *GetMemberRequest) Validate() error {

	if c.MemberNIK != `` && !shared.IsNIKFormat(c.MemberNIK) {
		return errors.New("invalid nik")
	}

	if c.FamilyNIK != `` && !shared.IsNIKFormat(c.FamilyNIK) {
		return errors.New("invalid nik")
	}

	return nil
}
