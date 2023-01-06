package repository

import (
	"context"

	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/usecase/transaction"
	"gorm.io/gorm"
)

type transactionRepository struct {
	qry *gorm.DB
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, params *transaction.CreateTransaction) error
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		qry: db,
	}
}

func (r *transactionRepository) CreateTransaction(ctx context.Context, params *transaction.CreateTransaction) error {
	var trx *model.Transaction

	if err := r.qry.Model(&trx).Create(map[string]interface{}{
		"trx_category":   params.TrxCategory,
		"trx_family_id":  params.TrxFamilyID,
		"trx_amount":     params.TrxAmount,
		"trx_type":       params.TrxType,
		"trx_note":       params.TrxNote,
		"trx_created_at": params.TrxCreatedAt,
	}).Error; err != nil {
		return err
	}

	return nil
}
