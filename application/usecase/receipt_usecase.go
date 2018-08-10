package usecase

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
)

type ReceiptUsecase interface {
	GetReceipt(id int) (*model.Receipt, error)
}

type receiptUsecase struct {
	repository.ReceiptRepository
}

func NewReceiptUsecase(r repository.ReceiptRepository) ReceiptUsecase {
	return &receiptUsecase{r}
}

func (r *receiptUsecase) GetReceipt(id int) (*model.Receipt, error) {
	receipt, err := r.FindOne(id)

	if err != nil {
		return nil, err
	}

	return receipt, nil
}
