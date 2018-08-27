package usecase

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
)

// ReceiptUsecase - レシートに関するユースケースを表すインターフェース
type ReceiptUsecase interface {
	GetReceipt(id int) (*model.Receipt, error)
	PostReceipt(userID int, name string, price int, kind, date, memo string, creditID int) (*model.Receipt, error)
	DeleteReceipt(receiptID int) error
}

type receiptUsecase struct {
	repository.ReceiptRepository
}

// NewReceiptUsecase - ReceiptUsecaseの生成
func NewReceiptUsecase(r repository.ReceiptRepository) ReceiptUsecase {
	return &receiptUsecase{r}
}

// GetReceipt - レシートの取得
func (r *receiptUsecase) GetReceipt(id int) (*model.Receipt, error) {
	receipt, err := r.FindOne(id)
	return receipt, err
}

// PostReceipt - レシートの新規登録
func (r *receiptUsecase) PostReceipt(userID int, name string, price int, kind, date, memo string, creditID int) (*model.Receipt, error) {
	receipt, err := model.NewReceipt(userID, name, price, kind, date, memo, creditID)
	if err != nil {
		return nil, err
	}

	createdReceipt, err := r.Create(receipt)

	if err != nil {
		return nil, err
	}

	return createdReceipt, nil
}

// DeleteReceipt - レシートの削除
func (r *receiptUsecase) DeleteReceipt(receiptID int) error {
	err := r.Delete(receiptID)
	return err
}
