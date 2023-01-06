package transaction

import (
	"errors"
	"time"

	"github.com/Hulhay/goldfish/model"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateTransactionRequest struct {
	TrxCategory string  `json:"trx_category"`
	TrxFamilyID string  `json:"trx_family_id"`
	TrxAmount   float64 `json:"trx_amount"`
	TrxType     string  `json:"trx_type"`
	TrxNote     string  `json:"trx_note"`
}

type CreateTransaction struct {
	TrxCategory  string
	TrxFamilyID  string
	TrxAmount    float64
	TrxType      string
	TrxNote      string
	TrxCreatedAt time.Time
}

func (c *CreateTransactionRequest) Validate() error {

	if err := validation.Validate(c.TrxCategory, validation.Required); err != nil {
		return errors.New("category must be filled")
	}

	if err := validation.Validate(c.TrxAmount, validation.Required); err != nil {
		return errors.New("amount must be filled")
	}

	if err := validation.Validate(c.TrxType, validation.Required); err != nil {
		return errors.New("type must be filled")
	}

	if !model.IsValidTrxType[c.TrxType] {
		return errors.New("invalid type")
	}

	if c.TrxCategory == model.KAS_BULANAN {
		if err := validation.Validate(c.TrxFamilyID, validation.Required); err != nil {
			return errors.New("family_id must be filled")
		}
	}

	if c.TrxCategory == model.LAIN_LAIN {
		if err := validation.Validate(c.TrxNote, validation.Required); err != nil {
			return errors.New("note must be filled")
		}
	}

	return nil
}
