package cmd

import (
	"fmt"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/repo"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/service"
	"log"

	"github.com/spf13/cobra"
)

var (
	description string
	amount      float64
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		repository := repo.NewJSONExpenseRepository("data/expense.json")
		expenseService := service.NewExpenseService(repository)

		expense, err := expenseService.AddExpense(description, amount)
		if err != nil {
			log.Fatalf("Error adding expense: %v", err)
		}

		fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Amount")

	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
