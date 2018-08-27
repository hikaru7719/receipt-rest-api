package datastore

import (
	"errors"
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysqlのドライバー
)

// ReceiptRepository - レシートに対応するリポジトリの構造体
type ReceiptRepository struct {
}

// NewReceiptRepository -　ReceiptRepositoryの生成
func NewReceiptRepository() repository.ReceiptRepository {
	return &ReceiptRepository{}
}

// FindOne - レシートの取得
func (r *ReceiptRepository) FindOne(receiptID int) (*model.Receipt, error) {
	receipt := new(model.Receipt)
	if DB.First(receipt, receiptID).RecordNotFound() {
		return nil, errors.New("record not found")
	}

	return receipt, nil
}

// Create - レシートの生成
func (r *ReceiptRepository) Create(receipt *model.Receipt) (*model.Receipt, error) {
	err := DB.Create(receipt).Error
	return receipt, err
}

// Delete - レシートの削除
func (r *ReceiptRepository) Delete(receiptID int) error {
	err := DB.Delete(&model.Receipt{ID: receiptID}).Error
	return err
}
