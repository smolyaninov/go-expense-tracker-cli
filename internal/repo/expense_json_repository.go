package repo

import (
	"encoding/json"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/domain"
	"os"
	"path/filepath"
)

type ExpenseRepository interface {
	Load() ([]domain.Expense, error)
	Save([]domain.Expense) error
}

type JSONExpenseRepository struct {
	filename string
}

func NewJSONExpenseRepository(filename string) *JSONExpenseRepository {
	return &JSONExpenseRepository{filename: filename}
}

func (r *JSONExpenseRepository) Load() ([]domain.Expense, error) {
	if _, err := os.Stat(r.filename); os.IsNotExist(err) {
		return []domain.Expense{}, nil
	}

	file, err := os.Open(r.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var expenses []domain.Expense
	if err := json.NewDecoder(file).Decode(&expenses); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *JSONExpenseRepository) Save(expenses []domain.Expense) error {
	if err := os.MkdirAll(filepath.Dir(r.filename), 0755); err != nil {
		return err
	}

	file, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		return err
	}

	tempFile := r.filename + ".tmp"
	if err := os.WriteFile(tempFile, file, 0644); err != nil {
		return err
	}

	return os.Rename(tempFile, r.filename)
}
