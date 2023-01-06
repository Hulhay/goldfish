package repository

import (
	"context"

	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/usecase/wallet"
	"gorm.io/gorm"
)

type walletRepository struct {
	qry *gorm.DB
}

type WalletRepository interface {
	GetWallet(ctx context.Context) (*model.Wallet, error)
	UpdateWallet(ctx context.Context, params *wallet.UpdateWalletRequest) error
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{
		qry: db,
	}
}

func (r *walletRepository) GetWallet(ctx context.Context) (*model.Wallet, error) {
	var wallet *model.Wallet

	if err := r.qry.Model(&wallet).First(&wallet).Error; err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *walletRepository) UpdateWallet(ctx context.Context, params *wallet.UpdateWalletRequest) error {
	var wallet *model.Wallet

	if err := r.qry.Model(&wallet).Where("wallet_id = ?", 1).Updates(map[string]interface{}{
		"wallet_balance":    params.WalletBalance,
		"wallet_updated_at": params.WalletUpdatedAt,
	}).Error; err != nil {
		return err
	}
	return nil
}
