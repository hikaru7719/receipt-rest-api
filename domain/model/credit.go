package model

// Credit -　クレジットカードを表す構造体
type Credit struct {
	ID                int    `json:"id" gorm:"primary_key"`
	UserID            int    `json:"user_id"`
	CardName          string `json:"card_name"`
	FinishDate        int    `json:"finish_date"`
	WithdrawalDate    int `json:"withdrawal_date"`
	LaterPaymentMonth int `json:"later_payment_month"`
}

// NewCredit - クレジットカードの生成
func NewCredit(userID int, cardName string, finishDate, withdrawalDate, laterPaymentMonth int) (*Credit, error) {
	return &Credit{UserID: userID, CardName: cardName,FinishDate:finishDate, WithdrawalDate: withdrawalDate, LaterPaymentMonth:laterPaymentMonth}, nil
}
