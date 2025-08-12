package repo

import "github.com/smolyaninov/go-expense-tracker-cli/internal/domain"

type ExpenseRepository interface {
	Load() ([]domain.Expense, error)
	Save([]domain.Expense) error
}

type jsonExpenseRepo struct {
	inner *JSONRepository[[]domain.Expense]
}

func NewJSONExpenseRepository(filename string) ExpenseRepository {
	return &jsonExpenseRepo{
		inner: NewJSONRepository[[]domain.Expense](filename),
	}
}

func (r *jsonExpenseRepo) Load() ([]domain.Expense, error) {
	return r.inner.Load()
}

func (r *jsonExpenseRepo) Save(data []domain.Expense) error {
	return r.inner.Save(data)
}
