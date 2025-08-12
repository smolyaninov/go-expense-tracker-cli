package cmd

import (
	"fmt"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/repo"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/service"
	"github.com/spf13/cobra"
	"log"
)

var summaryMonth int

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show total expenses (optional by month)",
	Run: func(cmd *cobra.Command, args []string) {
		repository := repo.NewJSONExpenseRepository("data/expense.json")
		expenseService := service.NewExpenseService(repository)

		if summaryMonth != 0 && (summaryMonth < 1 || summaryMonth > 12) {
			log.Fatalf("Month must be between 1 and 12")
		}

		if summaryMonth > 0 {
			total, err := expenseService.GetTotalAmountByMonth(summaryMonth)
			if err != nil {
				log.Fatalf("Error getting monthly summary: %v", err)
			}
			fmt.Printf("Total expenses for month %d: $%.2f\n", summaryMonth, total)
		} else {
			total, err := expenseService.GetTotalAmount()
			if err != nil {
				log.Fatalf("Error getting total summary: %v", err)
			}
			fmt.Printf("Total expenses: $%.2f\n", total)
		}
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().IntVarP(&summaryMonth, "month", "m", 0, "Filter summary by month (1-12)")
}
