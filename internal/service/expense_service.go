package service

import (
	"github.com/smolyaninov/go-expense-tracker-cli/internal/domain"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/repo"
)

type ExpenseService struct {
	repo repo.ExpenseRepository
}

func NewExpenseService(repo repo.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo}
}

func (s *ExpenseService) AddExpense(description string, amount float64) (domain.Expense, error) {
	expenses, err := s.repo.Load()
	if err != nil {
		return domain.Expense{}, err
	}

	newID := 1
	if len(expenses) > 0 {
		newID = expenses[len(expenses)-1].ID + 1
	}

	expense, err := domain.NewExpense(newID, description, amount)
	if err != nil {
		return domain.Expense{}, err
	}

	expenses = append(expenses, *expense)

	if err := s.repo.Save(expenses); err != nil {
		return domain.Expense{}, err
	}

	return *expense, nil
}
