package models

type Order struct {
	ID          int     `json:"id"`
	Customer    string  `json:"customer"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}
