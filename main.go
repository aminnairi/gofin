package main

import (
	"fmt"
	"time"

	"github.com/aminnairi/gofin/budget"
)

func main() {
	fmt.Println("Coming soon, being able to create your budget from your terminal!")

	budget := budget.Budget{
		Start: time.Date(2025, time.September, 1, 0, 0, 0, 0, time.Local),
		Incomes: []budget.Income{
			{
				// Salaire
				Amount:          3430,
				Start:           time.Date(2025, time.September, 1, 0, 0, 0, 0, time.Local),
				End:             time.Date(2025, time.December, 31, 0, 0, 0, 0, time.Local),
				MonthOccurrence: 1,
			},
		},
		Expenses: []budget.Expense{},
	}

	forecast := budget.Forecast(time.Date(2025, time.December, 31, 0, 0, 0, 0, time.Local))

	fmt.Println("Budget forecast:", forecast)
}
