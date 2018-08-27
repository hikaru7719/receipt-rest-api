package model

// Credit -　クレジットカードを表す構造体
type Credit struct {
	ID             int `gorm:"primary_key"`
	CardName       string
	OpeningDate    string
	ClosingDate    string
	WithdrawalDate string
	PaymentMonth   string
}

// NewCredit - クレジットカードの生成
func NewCredit(cardName, openingDate, closingDate, withdrawalDate string) (*Credit, error) {
	return &Credit{CardName: cardName, OpeningDate: openingDate, ClosingDate: closingDate, WithdrawalDate: withdrawalDate}, nil
}
