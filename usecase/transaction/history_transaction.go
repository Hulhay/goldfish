package transaction

import (
	"errors"
	"time"

	"github.com/Hulhay/goldfish/shared"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GetHistoryTransactionRequest struct {
	Category      string `query:"category"`
	Type          string `query:"type"`
	StartDate     string `query:"start_date"`
	EndDate       string `query:"end_date"`
	StartDateTime string
	EndDateTime   string
}

type GetHistoryTransactionResponse struct {
	TrxID        int
	TrxCategory  string
	TrxAmount    float64
	TrxType      string
	TrxCreatedAt time.Time
}

func (c *GetHistoryTransactionRequest) Validate() error {

	if err := validation.Validate(c.StartDate, validation.Required); err != nil {
		return errors.New("start_date must be filled")
	}

	startDate, err := time.Parse(shared.FormatDate, c.StartDate)
	if err != nil {
		return errors.New("invalid start_date format")
	}

	if err := validation.Validate(c.EndDate, validation.Required); err != nil {
		return errors.New("end_date must be filled")
	}

	endDate, err := time.Parse(shared.FormatDate, c.EndDate)
	if err != nil {
		return errors.New("invalid end_date format")
	}

	if endDate.Before(startDate) {
		return errors.New("end_date cannot less than start_date")
	}

	if (endDate.Sub(startDate).Hours() / 24) >= 31 {
		return errors.New("range date maximum 31 days")
	}

	c.StartDateTime = shared.StartDateString(c.StartDate)
	c.EndDateTime = shared.EndDateString(c.EndDate)

	return nil
}
