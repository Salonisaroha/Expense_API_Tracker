package models

import "time"

type Expense struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
