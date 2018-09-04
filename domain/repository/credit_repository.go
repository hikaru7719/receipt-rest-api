package repository

import "github.com/hikaru7719/receipt-rest-api/domain/model"

// CreditRepository - クレジットカードに対応するリポジトリのインターフェース
type CreditRepository interface {
	Create(credit *model.Credit) (*model.Credit, error)
	FindAllByUserID(userID int) ([]*model.Credit, error)
}
