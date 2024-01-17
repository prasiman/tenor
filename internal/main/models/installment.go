package models

import (
	"time"
)

type Installment struct {
	ID           uint64    `gorm:"primarykey" json:"id"`
	ContractID   uint64    `json:"contract_id"`
	PaymentDue   float64   `json:"payment_due"`
	DueDate      time.Time `json:"due_date"`
	Status       string    `json:"status"`
	TotalLateFee float64   `json:"total_late_fee"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
