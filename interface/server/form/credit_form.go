package form

// CreditForm - JSONの一時保管先
type CreditForm struct {
	UserID            int    `json:"user_id"`
	CardName          string `json:"card_name"`
	FinishDate        int    `json:"finish_date"`
	WithdrawalDate    int    `json:"withdrawal_date"`
	LaterPaymentMonth int    `json:"later_payment_month"`
}
