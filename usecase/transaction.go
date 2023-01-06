package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Hulhay/goldfish/model"
	"github.com/Hulhay/goldfish/repository"
	"github.com/Hulhay/goldfish/usecase/transaction"
	"github.com/Hulhay/goldfish/usecase/wallet"
)

type transactionUC struct {
	transactionRepo repository.TransactionRepository
	categoryRepo    repository.CategoryRepository
	familyRepo      repository.FamilyRepository
	walletRepo      repository.WalletRepository
	timeRepo        repository.TimeRepository
}

type Transaction interface {
	CreateTransaction(ctx context.Context, params transaction.CreateTransactionRequest) error
	GetHistoryTransaction(ctx context.Context, params transaction.GetHistoryTransactionRequest) ([]transaction.GetHistoryTransactionResponse, error)
}

func NewTransactionUC(
	tr repository.TransactionRepository,
	fr repository.FamilyRepository,
	wr repository.WalletRepository,
	cr repository.CategoryRepository,
	timeRepo repository.TimeRepository,
) Transaction {
	return &transactionUC{
		transactionRepo: tr,
		familyRepo:      fr,
		walletRepo:      wr,
		categoryRepo:    cr,
		timeRepo:        timeRepo,
	}
}

func (u *transactionUC) CreateTransaction(ctx context.Context, params transaction.CreateTransactionRequest) error {

	var (
		err           error
		balance       float64
		currentWallet *model.Wallet
		categoryExist *model.Category
	)

	if err = params.Validate(); err != nil {
		return err
	}

	categoryExist, err = u.categoryRepo.GetCategoryByCategoryValue(ctx, params.TrxCategory)
	if err != nil || categoryExist == nil {
		return errors.New("invalid category")
	}

	categoryValue := categoryExist.CategoryValue

	currentWallet, err = u.walletRepo.GetWallet(ctx)
	if err != nil || currentWallet == nil {
		return errors.New("error while get wallet")
	}

	if params.TrxFamilyID != `` {
		familyExist, err := u.familyRepo.GetFamilyByFamilyID(ctx, params.TrxFamilyID)
		if err != nil || familyExist == nil {
			return errors.New("family not found")
		}
	}

	now := u.timeRepo.Now(time.Now())
	err = u.transactionRepo.CreateTransaction(ctx, &transaction.CreateTransaction{
		TrxCategory:  categoryValue,
		TrxFamilyID:  params.TrxFamilyID,
		TrxAmount:    params.TrxAmount,
		TrxType:      params.TrxType,
		TrxNote:      params.TrxNote,
		TrxCreatedAt: now,
	})
	if err != nil {
		return err
	}

	switch params.TrxType {
	case model.DEBIT:
		balance = currentWallet.WalletBalance + params.TrxAmount
	case model.CREDIT:
		balance = currentWallet.WalletBalance - params.TrxAmount
	}

	err = u.walletRepo.UpdateWallet(ctx, &wallet.UpdateWalletRequest{
		WalletBalance:   balance,
		WalletUpdatedAt: now,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *transactionUC) GetHistoryTransaction(ctx context.Context, params transaction.GetHistoryTransactionRequest) ([]transaction.GetHistoryTransactionResponse, error) {

	var (
		err             error
		transactionData []model.Transaction
		category        *model.Category
		res             []transaction.GetHistoryTransactionResponse
	)

	if err = params.Validate(); err != nil {
		return nil, err
	}

	transactionData, err = u.transactionRepo.GetHistoryTransaction(ctx, params)
	if err != nil {
		return nil, err
	}
	if transactionData == nil {
		return res, nil
	}

	for _, trx := range transactionData {

		category, err = u.categoryRepo.GetCategoryByCategoryValue(ctx, trx.TrxCategory)
		if err != nil {
			return nil, err
		}

		res = append(res, transaction.GetHistoryTransactionResponse{
			TrxID:        int(trx.ID),
			TrxCategory:  category.CategoryName,
			TrxAmount:    trx.TrxAmount,
			TrxType:      trx.TrxType,
			TrxCreatedAt: trx.TrxCreatedAt,
		})
	}

	return res, nil
}
