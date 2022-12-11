package model

type Member struct {
	ID           int64  `json:"id" gorm:"column:member_id;type:int;primary key;auto_increment"`
	MemberNIK    string `json:"member_nik" gorm:"column:member_nik;type:varchar(255)"`
	MemberName   string `json:"member_name" gorm:"column:member_name;type:varchar(255)"`
	MemberIsHead bool   `json:"member_is_head" gorm:"column:member_is_head;type:bool"`

	FamilyID string `json:"family_id" gorm:"column:family_id;type:varchar(255)"`
}
