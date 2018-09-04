package usecase

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/domain/repository"
)

// CreditUsecase - クレジットに関するユースケース
type CreditUsecase interface {
	PostCredit(userID int, cardName string, finishDate, withdrawalDate, laterPaymentMonth int) (*model.Credit, error)
	GetCreditList(userID int) ([]*model.Credit, error)
}

type creditUsecase struct {
	repository.CreditRepository
}

// NewCreditUsecase - creditUsecaseの生成
func NewCreditUsecase(r repository.CreditRepository) CreditUsecase {
	return &creditUsecase{r}
}

// PostCredit - クレジットカードの登録
func (c *creditUsecase) PostCredit(userID int, cardName string, finishDate, withdrawalDate, laterPaymentMonth int) (*model.Credit, error) {
	credit, err := model.NewCredit(userID, cardName, finishDate, withdrawalDate, laterPaymentMonth)
	if err != nil {
		return nil, err
	}
	createdCredit, err := c.Create(credit)
	if err != nil {
		return nil, err
	}
	return createdCredit, nil
}

// GetCreditList - クレジットカードのリストの取得
func (c *creditUsecase) GetCreditList(userID int) ([]*model.Credit, error) {
	creditList, err := c.FindAllByUserID(userID)
	if err != nil {
		return nil, err
	}
	return creditList, nil
}
