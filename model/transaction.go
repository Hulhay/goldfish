package model

import "time"

const (
	DEBIT  = "DEBIT"
	CREDIT = "CREDIT"
)

var (
	IsValidTrxType = map[string]bool{
		DEBIT:  true,
		CREDIT: true,
	}
)

type Transaction struct {
	ID           int64     `json:"id" gorm:"column:trx_id;type:int;primary key;auto_increment"`
	TrxCategory  string    `json:"trx_category" gorm:"column:trx_category;type:varchar(255)"`
	TrxFamilyID  string    `json:"trx_family_id" gorm:"column:trx_family_id;type:varchar(255)"`
	TrxAmount    float64   `json:"trx_amount" gorm:"column:trx_amount;type:float"`
	TrxType      string    `json:"trx_type" gorm:"column:trx_type;type:varchar(255)"`
	TrxNote      string    `json:"trx_note" gorm:"column:trx_note;type:varchar(255)"`
	TrxCreatedAt time.Time `json:"trx_created_at" gorm:"column:trx_created_at;type:timestamp"`
}
