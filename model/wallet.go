package model

import "time"

type Wallet struct {
	ID              int64     `json:"id" gorm:"column:wallet_id;type:int;primary key;auto_increment"`
	WalletBalance   float64   `json:"wallet_balance" gorm:"column:wallet_balance;type:double"`
	WalletUpdatedAt time.Time `json:"wallet_updated_at" gorm:"column:wallet_updated_at;type:timestamp"`
}
