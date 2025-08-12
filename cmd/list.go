package cmd

import (
	"fmt"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/repo"
	"github.com/smolyaninov/go-expense-tracker-cli/internal/service"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List as expenses",
	Run: func(cmd *cobra.Command, args []string) {
		repository := repo.NewJSONExpenseRepository("data/expense.json")
		expenseService := service.NewExpenseService(repository)

		expenses, err := expenseService.GetAllExpenses()
		if err != nil {
			log.Fatalf("Error getting expenses: %v", err)
		}

		if len(expenses) == 0 {
			fmt.Println("No expenses found")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tDate\tDescription\tAmount")
		for _, e := range expenses {
			fmt.Fprintf(w, "%d\t%s\t%s\t$%.2f\n", e.ID, e.Date.Format("2006-01-02"), e.Description, e.Amount)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
