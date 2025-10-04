package budget

import (
	"math"
	"time"
)

type Expense struct {
	amount          float32
	start           time.Time
	monthOccurrence int8
}

type Income struct {
	amount          float32
	start           time.Time
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

	start := budget.start

	for start.Before(date) {
		startMonth := int8(start.Month())

		for _, expense := range budget.expenses {
			expenseMonth := int8(expense.start.Month())
			monthDelta := int8(math.Abs(float64(startMonth - expenseMonth)))

			if monthDelta%expense.monthOccurrence == 0 {
				amount -= expense.amount
			}
		}

		for _, income := range budget.incomes {
			incomeMonth := int8(income.start.Month())
			monthDelta := int8(math.Abs(float64(startMonth - incomeMonth)))

			if monthDelta%income.monthOccurrence == 0 {
				amount += income.amount
			}
		}

		start = start.AddDate(0, 1, 0)
	}

	return amount
}
