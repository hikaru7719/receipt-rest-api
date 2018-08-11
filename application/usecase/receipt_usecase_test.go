package usecase

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"testing"
)

func TestReceiptUsecase_GetReceipt(t *testing.T) {
	usecaseTest := NewReceiptUsecase(&ReceiptRepositoryStub{})
	actual, err := usecaseTest.GetReceipt(1)

	if err != nil {
		t.Error("error")
	}

	if actual.Name != "test" {
		t.Error("error")
	}
}

type ReceiptRepositoryStub struct{}

func (r *ReceiptRepositoryStub) FindOne(receiptId int) (*model.Receipt, error) {
	receipt := &model.Receipt{Name: "test"}
	return receipt, nil
}

func (r *ReceiptRepositoryStub) Create(receipt *model.Receipt) (*model.Receipt, error) {
	return receipt, nil
}
