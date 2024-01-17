package models

type CreditLimit struct {
	// ID            uint64    `gorm:"primarykey" json:"id"`
	// UserID        uint64    `json:"user_id"`
	TenorMonth    uint    `json:"tenor_month"`
	CurrentAmount float64 `json:"current_amount"`
	MaxAmount     float64 `json:"max_amount"`
	// CreatedAt     time.Time `json:"created_at"`
	// UpdatedAt     time.Time `json:"updated_at"`
}

// Get credit limit by user ID
func GetCreditLimit(userId uint64) ([]CreditLimit, error) {
	tx := db.Raw("SELECT tenor_month, max_amount FROM credit_limits WHERE user_id = ? ORDER BY tenor_month ASC", userId)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var limits []CreditLimit
	tx.Scan(&limits)
	// defer sqlDb.Close()

	for key, value := range limits {
		total, _ := GetSumOfUsedCreditLimitByTenorMonth(userId, value.TenorMonth)

		limits[key].CurrentAmount = value.MaxAmount - total
	}

	return limits, nil
}
