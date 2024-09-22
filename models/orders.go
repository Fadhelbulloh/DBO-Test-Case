package models

import "time"

type Order struct {
	ID            uint      `json:"id"`
	OrderNumber   string    `json:"order_number"`
	CustomerID    uint      `json:"customer_id"`
	TotalAmount   float64   `json:"total_amount"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
