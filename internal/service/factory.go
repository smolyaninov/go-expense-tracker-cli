package service

import "github.com/smolyaninov/go-expense-tracker-cli/internal/repo"

func NewDefaultExpenseService() *ExpenseService {
	return NewExpenseService(repo.NewJSONExpenseRepository("data/expense.json"))
}

func NewDefaultBudgetService() *BudgetService {
	return NewBudgetService(repo.NewJSONBudgetRepository("data/budget.json"))
}
