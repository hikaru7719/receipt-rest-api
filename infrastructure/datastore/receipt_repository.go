package datastore

import (
	"errors"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
	"github.com/jinzhu/gorm"
)

// ReceiptRepository - レシートに対応するリポジトリの構造体
type ReceiptRepository struct {
	db *gorm.DB
}

// NewReceiptRepository -　ReceiptRepositoryの生成
func NewReceiptRepository(db *gorm.DB) repository.ReceiptRepository {
	return &ReceiptRepository{db: db}
}

// FindOne - レシートの取得
func (r *ReceiptRepository) FindOne(receiptID int) (*model.Receipt, error) {
	receipt := new(model.Receipt)
	if r.db.First(receipt, receiptID).RecordNotFound() {
		return nil, errors.New("record not found")
	}

	return receipt, nil
}

// Create - レシートの生成
func (r *ReceiptRepository) Create(receipt *model.Receipt) (*model.Receipt, error) {
	if err := r.db.Create(receipt).Error; err != nil {
		return receipt, err
	}

	return receipt, nil
}
