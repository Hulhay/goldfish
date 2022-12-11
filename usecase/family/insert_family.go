package family

import (
	"errors"

	"github.com/Hulhay/goldfish/shared"
)

type InsertFamilyRequest struct {
	FamilyID           string `json:"family_id"`
	FamilyNIK          string `json:"family_nik"`
	FamilyMemberHeadID int64  `json:"family_member_head_id"`
}

func (c *InsertFamilyRequest) Validate() error {

	if !shared.IsNIKFormat(c.FamilyNIK) {
		return errors.New("invalid nik")
	}

	return nil
}
