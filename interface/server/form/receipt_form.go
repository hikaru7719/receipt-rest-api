package form

// ReceiptForm - JSONの一時保管先
type ReceiptForm struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Kind     string `json:"kind"`
	Date     string `json:"date"`
	Memo     string `json:"memo"`
	CreditID int    `json:"credit_id"`
}
