package model

type Family struct {
	ID                 int64  `json:"id" gorm:"column:id;type:int;primary key;auto_increment"`
	FamilyID           string `json:"family_id" gorm:"column:family_id;type:varchar(255)"`
	FamilyNIK          string `json:"family_nik" gorm:"column:family_nik;type:varchar(255)"`
	FamilyMemberHeadID int64  `json:"family_member_head_id" gorm:"column:family_member_head_id;type:int"`
}
