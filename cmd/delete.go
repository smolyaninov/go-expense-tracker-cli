package cmd

import (
	"fmt"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/repo"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/service"

	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense by ID",
	Run: func(cmd *cobra.Command, args []string) {
		repository := repo.NewJSONExpenseRepository("data/expense.json")
		expenseService := service.NewExpenseService(repository)

		err := expenseService.DeleteExpense(deleteID)
		if err != nil {
			fmt.Printf("Error deleting expense: %v", err)
			return
		}

		fmt.Println("Expense deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "ID")
	deleteCmd.MarkFlagRequired("id")
}
