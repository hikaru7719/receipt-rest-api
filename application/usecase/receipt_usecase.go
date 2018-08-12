package usecase

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
)

type ReceiptUsecase interface {
	GetReceipt(id int) (*model.Receipt, error)
	PostReceipt(name, kind, date, memo string) (*model.Receipt, error)
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

func (r *receiptUsecase) PostReceipt(name, kind, date, memo string) (*model.Receipt, error) {
	receipt, err := model.NewReceipt(name, kind, date, memo)
	if err != nil {
		return nil, err
	}

	createdReceipt, err := r.Create(receipt)

	if err != nil {
		return nil, err
	}

	return createdReceipt, nil
}
