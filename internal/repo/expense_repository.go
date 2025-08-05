package repo

import "github.com/smolyaninov/go-expense-tracker-cli/internal/domain"

type ExpenseRepository interface {
	Load() ([]domain.Expense, error)
	Save([]domain.Expense) error
}
