package repository

import "github.com/hikaru7719/receipt-rest-api/domain/model"

type ReceiptRepository interface {
	FindOne(receiptID int) (*model.Receipt, error)
}
