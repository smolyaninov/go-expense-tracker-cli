package cmd

import (
	"fmt"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/repo"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/service"
	"github.com/spf13/cobra"
	"log"
)

var (
	updateID          int
	updateDescription string
	updateAmount      float64
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an expense by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if updateDescription == "" && updateAmount == 0 {
			log.Fatalf("You must provide at least one field to update: --description or --amount")
		}

		repository := repo.NewJSONExpenseRepository("data/expense.json")
		expenseService := service.NewExpenseService(repository)

		err := expenseService.UpdateExpense(updateID, updateDescription, updateAmount)
		if err != nil {
			log.Fatalf("Error updating expense: %v", err)
		}

		fmt.Println("Expense updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "ID of the expense to update")
	updateCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "New description")
	updateCmd.Flags().Float64VarP(&updateAmount, "amount", "a", 0, "New amount")

	updateCmd.MarkFlagRequired("id")
}
