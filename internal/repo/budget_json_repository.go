package repo

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type BudgetRepository interface {
	Load() (map[string]float64, error)
	Save(map[string]float64) error
}

type JSONBudgetRepository struct {
	filename string
}

func NewJSONBudgetRepository(filename string) *JSONBudgetRepository {
	return &JSONBudgetRepository{filename: filename}
}

func (r *JSONBudgetRepository) Load() (map[string]float64, error) {
	if _, err := os.Stat(r.filename); os.IsNotExist(err) {
		return map[string]float64{}, nil
	}

	file, err := os.Open(r.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var budgets map[string]float64
	if err := json.NewDecoder(file).Decode(&budgets); err != nil {
		return nil, err
	}

	return budgets, nil
}

func (r *JSONBudgetRepository) Save(budgets map[string]float64) error {
	if err := os.MkdirAll(filepath.Dir(r.filename), 0755); err != nil {
		return err
	}

	file, err := json.MarshalIndent(budgets, "", " ")
	if err != nil {
		return err
	}

	tempFile := r.filename + ".tmp"
	if err := os.WriteFile(tempFile, file, 0644); err != nil {
		return err
	}

	return os.Rename(tempFile, r.filename)
}
