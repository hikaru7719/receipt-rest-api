package form

type ReceiptForm struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
	Date string `json:"date"`
	Memo string `json:"memo"`
}
