package domain

import (
	"fmt"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

func NewExpense(id int, description string, amount float64) (*Expense, error) {
	if description == "" {
		return nil, fmt.Errorf("description cannot be empty")
	}

	if amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}

	return &Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		Date:        time.Now().UTC(),
	}, nil
}
