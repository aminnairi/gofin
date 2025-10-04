package budget

import (
	"fmt"
	"math"
	"time"
)

type Expense struct {
	amount          float32
	start           time.Time
	end             time.Time
	monthOccurrence int8
}

type Income struct {
	amount          float32
	start           time.Time
	end             time.Time
	monthOccurrence int8
}

type Budget struct {
	start    time.Time
	expenses []Expense
	incomes  []Income
}

func (budget Budget) Forecast(date time.Time) (amount float32) {
	if budget.start.After(date) {
		return 0
	}

	currentDate := budget.start

	for currentDate.Before(date) {
		startMonth := int8(currentDate.Month())

		for _, expense := range budget.expenses {
			if expense.start.Before(budget.start) {
				continue
			}

			if expense.end.Before(currentDate) {
				continue
			}

			if expense.monthOccurrence < 0 {
				continue
			}

			expenseMonth := int8(expense.start.Month())
			monthDelta := int8(math.Abs(float64(startMonth - expenseMonth)))

			if expense.monthOccurrence != 0 && monthDelta%expense.monthOccurrence != 0 {
				fmt.Println("Dépense qui ne correspond pas a la date courant")
				continue
			}

			fmt.Println("Dépense correspondante")

			if expense.monthOccurrence == 0 && !expense.start.Equal(currentDate) {
				continue
			}

			amount -= expense.amount
		}

		for _, income := range budget.incomes {
			if income.start.Before(budget.start) {
				continue
			}

			if income.end.Before(currentDate) {
				continue
			}

			if income.monthOccurrence < 0 {
				continue
			}

			incomeMonth := int8(income.start.Month())
			monthDelta := int8(math.Abs(float64(startMonth - incomeMonth)))

			if income.monthOccurrence != 0 && monthDelta%income.monthOccurrence != 0 {
				continue
			}

			if income.monthOccurrence == 0 && !income.start.Equal(currentDate) {
				continue
			}

			amount += income.amount
		}

		currentDate = currentDate.AddDate(0, 1, 0)
	}

	return amount
}
