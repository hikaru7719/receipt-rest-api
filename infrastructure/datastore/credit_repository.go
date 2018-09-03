package datastore

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
)

// CreditRepository - クレジットカードに関するリポジトリ
type CreditRepository struct {
}

// NewCreditRepository - CreditRepositoryの生成
func NewCreditRepository() repository.CreditRepository {
	return &CreditRepository{}
}

// Create - クレジットカードの生成
func (r *CreditRepository) Create(credit *model.Credit) (*model.Credit, error) {
	err := DB.Create(credit).Error
	return credit, err
}

// FindAllByUserID - 該当ユーザのクレジットカード全件取得
func (r *CreditRepository) FindAllByUserID(userID int) ([]*model.Credit, error) {
	creditSlice := make([]*model.Credit, 15)
	err := DB.Select(creditSlice, "user_id=?", userID).Error
	return creditSlice, err
}
