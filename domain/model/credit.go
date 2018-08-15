package model

// Credit -　クレジットカードを表す構造体
type Credit struct {
	ID             int
	CardName       string
	OpenigDate     string
	ClosingDate    string
	WithdrawalDate string
}


// NewCredit - クレジットカードの生成
func NewCredit(cardName, openingDate, closingDate, withdrawalDate string) (*Credit, error) {
	return &Credit{CardName: cardName, OpenigDate: openingDate, ClosingDate: closingDate, WithdrawalDate: withdrawalDate},nil
}
