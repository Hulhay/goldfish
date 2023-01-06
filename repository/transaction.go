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
	GetHistoryTransaction(ctx context.Context, params transaction.GetHistoryTransactionRequest) ([]model.Transaction, error)
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

func (r *transactionRepository) GetHistoryTransaction(ctx context.Context, params transaction.GetHistoryTransactionRequest) ([]model.Transaction, error) {
	var res []model.Transaction

	db := r.qry.Model(&model.Transaction{})

	if params.Category != `` {
		db = db.Where("trx_category = ?", params.Category)
	}

	if params.Type != `` {
		db = db.Where("trx_type = ?", params.Type)
	}

	if params.StartDate != `` {
		db = db.Where("trx_created_at > ?", params.StartDateTime)
	}

	if params.EndDate != `` {
		db = db.Where("trx_created_at < ?", params.EndDateTime)
	}

	db = db.Order("trx_created_at DESC")

	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
