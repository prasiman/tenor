package models

import (
	"errors"
	"time"

	"main/pkg/hash"
)

type User struct {
	ID            uint64    `gorm:"primarykey" json:"id"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	FullName      string    `json:"full_name"`
	LegalName     string    `json:"legal_name"`
	IDNumber      string    `json:"id_number"`
	CityOfBirth   string    `json:"city_of_birth"`
	DateOfBirth   time.Time `json:"date_of_birth"`
	Salary        float64   `json:"salary"`
	IDNumberPhoto string    `json:"id_number_photo"`
	SelfiePhoto   string    `json:"selfie_photo"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func Get() []User {
	var users []User
	db.Find(&users)
	// defer sqlDb.Close()

	return users
}

func UserCreate(email string, password string) (bool, error) {
	existingUser := GetUserByEmail(email)

	if existingUser != nil {
		return false, errors.New("user already exist")
	}

	hashedPassword, err := hash.HashString(password)

	if err != nil {
		return false, err
	}

	tx := db.Exec("INSERT INTO users(email,password) VALUES(?,?)", email, hashedPassword)
	// defer sqlDb.Close()

	if tx.Error != nil {
		return false, tx.Error
	}

	return true, nil
}

func UserAuthenticate(email string, password string) (bool, *User, error) {
	user := GetUserByEmail(email)

	if user == nil {
		return false, nil, errors.New("data not found")
	}

	if user.ValidatePassword(password) {
		return true, user, nil
	}

	return false, user, errors.New("wrong password")
}

func (user *User) ValidatePassword(password string) bool {
	return hash.CheckStringHash(password, user.Password)
}

func GetUserByEmail(email string) *User {
	tx := db.Raw("SELECT * FROM users WHERE email = ?", email)

	if tx.Error != nil {
		return nil
	}

	var user *User
	tx.Scan(&user)

	return user
}

// Check if a user exists in database by given ID
func GetUserById(id float64) (*User, error) {
	tx := db.Raw("SELECT * FROM users WHERE id = ?", id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var user *User
	tx.Scan(&user)
	// defer sqlDb.Close()

	return user, nil
}
