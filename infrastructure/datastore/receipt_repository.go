package datastore

import (
	"errors"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
	"github.com/jinzhu/gorm"
)

type ReceiptRepository struct {
	db *gorm.DB
}

func NewReceiptRepository(db *gorm.DB) repository.ReceiptRepository {
	return &ReceiptRepository{db: db}
}

func (r *ReceiptRepository) FindOne(receiptId int) (*model.Receipt, error) {
	receipt := new(model.Receipt)
	if r.db.First(receipt, receiptId).RecordNotFound() {
		return nil, errors.New("record not found")
	}

	return receipt, nil
}

func (r *ReceiptRepository) Create(receipt *model.Receipt) (*model.Receipt, error) {
	if err := r.db.Create(receipt).Error; err != nil {
		return receipt, err
	}

	return receipt, nil
}
