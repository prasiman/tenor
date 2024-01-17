package models

import (
	"errors"
	"time"
)

type Contract struct {
	ID         uint64    `gorm:"primarykey" json:"id"`
	UserID     uint64    `json:"user_id"`
	AssetName  string    `json:"asset_name"`
	ContractNo string    `json:"contract_no"`
	OTR        float64   `json:"otr"`
	AdminFee   float64   `json:"admin_fee"`
	TenorMonth uint      `json:"tenor_month"`
	Interest   float64   `json:"interest"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Get all contracts by user ID
func GetAllContracts(userId uint64) ([]Contract, error) {
	tx := db.Raw("SELECT * FROM contracts WHERE user_id = ?", userId)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var contracts []Contract
	tx.Scan(&contracts)
	// defer sqlDb.Close()

	return contracts, nil
}

// Get a contract by ID and user ID
func GetContractByID(id uint64, userId uint64) (*Contract, error) {
	tx := db.Raw("SELECT * FROM contracts WHERE id = ? AND user_id = ?", id, userId)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected < 1 {
		return nil, errors.New("data not found")
	}

	var contract *Contract
	tx.Scan(&contract)
	// defer sqlDb.Close()

	return contract, nil
}

func GetSumOfUsedCreditLimitByTenorMonth(userId uint64, tenorMonth uint) (float64, error) {
	var sum *struct {
		Total float64 `json:"total"`
	}

	tx := db.Raw("SELECT SUM(otr) as total FROM contracts WHERE user_id = ? AND tenor_month = ?", userId, tenorMonth)

	if tx.Error != nil {
		return 0, tx.Error
	}

	tx.Scan(&sum)

	return sum.Total, nil
}
