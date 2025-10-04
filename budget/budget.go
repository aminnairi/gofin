package budget

import (
	"fmt"
	"math"
	"time"
)

type Expense struct {
	Amount          float32
	Start           time.Time
	End             time.Time
	MonthOccurrence int8
}

type Income struct {
	Amount          float32
	Start           time.Time
	End             time.Time
	MonthOccurrence int8
}

type Budget struct {
	Start    time.Time
	Expenses []Expense
	Incomes  []Income
}

func (budget Budget) Forecast(date time.Time) (amount float32) {
	if budget.Start.After(date) {
		return 0
	}

	currentDate := budget.Start

	for currentDate.Before(date) {
		startMonth := int8(currentDate.Month())

		for _, expense := range budget.Expenses {
			if expense.End.Before(currentDate) {
				continue
			}

			if expense.MonthOccurrence < 0 {
				continue
			}

			expenseMonth := int8(expense.Start.Month())
			monthDelta := int8(math.Abs(float64(startMonth - expenseMonth)))

			if expense.MonthOccurrence != 0 && monthDelta%expense.MonthOccurrence != 0 {
				fmt.Println("Dépense qui ne correspond pas a la date courant")
				continue
			}

			fmt.Println("Dépense correspondante")

			if expense.MonthOccurrence == 0 && !expense.Start.Equal(currentDate) {
				continue
			}

			amount -= expense.Amount
		}

		for _, income := range budget.Incomes {
			if income.End.Before(currentDate) {
				continue
			}

			if income.MonthOccurrence < 0 {
				continue
			}

			incomeMonth := int8(income.Start.Month())
			monthDelta := int8(math.Abs(float64(startMonth - incomeMonth)))

			if income.MonthOccurrence != 0 && monthDelta%income.MonthOccurrence != 0 {
				continue
			}

			if income.MonthOccurrence == 0 && !income.Start.Equal(currentDate) {
				continue
			}

			amount += income.Amount
		}

		currentDate = currentDate.AddDate(0, 1, 0)
	}

	return amount
}
