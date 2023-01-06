package wallet

import "time"

type UpdateWalletRequest struct {
	WalletBalance   float64   `json:"wallet_balance"`
	WalletUpdatedAt time.Time `json:"wallet_updated_at"`
}
