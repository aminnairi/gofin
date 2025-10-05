package budget

import (
	"math"
	"time"
)

type Exception struct {
	Amount          float32
	Start           time.Time
	End             time.Time
	MonthOccurrence int8
}

type Expense struct {
	Amount          float32
	Start           time.Time
	End             time.Time
	MonthOccurrence int8
	Exceptions      []Exception
}

type Income struct {
	Amount          float32
	Start           time.Time
	End             time.Time
	MonthOccurrence int8
	Exceptions      []Exception
}

type Budget struct {
	Start    time.Time
	Expenses []Expense
	Incomes  []Income
}

func (budget Budget) Forecast(date time.Time) (forecast float32) {
	if budget.Start.After(date) {
		return 0
	}

	currentDate := budget.Start

	for currentDate.Before(date) {
		startMonth := int8(currentDate.Month())

		for _, expense := range budget.Expenses {
			exceptionFound := false

			for _, exception := range expense.Exceptions {
				if exception.MonthOccurrence == 0 {
					if exception.Start.Equal(currentDate) || exception.End.Equal(currentDate) {
						forecast -= exception.Amount
						exceptionFound = true
						break
					}

					continue
				}

				if (exception.Start.Equal(currentDate) || exception.Start.After(currentDate)) && (exception.End.Equal(currentDate) || exception.End.Before(currentDate)) {
					forecast -= exception.Amount
					exceptionFound = true
					break
				}
			}

			if exceptionFound {
				continue
			}

			if expense.End.Before(currentDate) {
				continue
			}

			if expense.MonthOccurrence < 0 {
				continue
			}

			expenseMonth := int8(expense.Start.Month())
			monthDelta := int8(math.Abs(float64(startMonth - expenseMonth)))

			if expense.MonthOccurrence != 0 && monthDelta%expense.MonthOccurrence != 0 {
				continue
			}

			if expense.MonthOccurrence == 0 && !expense.Start.Equal(currentDate) {
				continue
			}

			forecast -= expense.Amount
		}

		for _, income := range budget.Incomes {
			exceptionFound := false

			for _, exception := range income.Exceptions {
				if exception.MonthOccurrence == 0 {
					if exception.Start.Equal(currentDate) || exception.End.Equal(currentDate) {
						forecast -= exception.Amount
						exceptionFound = true
						break
					}

					continue
				}

				if (exception.Start.Equal(currentDate) || exception.Start.After(currentDate)) && (exception.End.Equal(currentDate) || exception.End.Before(currentDate)) {
					forecast -= exception.Amount
					exceptionFound = true
					break
				}
			}

			if exceptionFound {
				continue
			}

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

			forecast += income.Amount
		}

		currentDate = currentDate.AddDate(0, 1, 0)
	}

	return forecast
}
