package budget

import (
	"testing"
	"time"
)

func TestBudgetForecast(t *testing.T) {
	budget := Budget{
		Start: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
		Incomes: []Income{
			{
				Amount:          2000,
				Start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				End:             time.Date(2050, time.December, 1, 0, 0, 0, 0, time.Local),
				MonthOccurrence: 1,
			},
		},
		Expenses: []Expense{
			{
				Amount:          1000,
				Start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				End:             time.Date(2050, time.December, 1, 0, 0, 0, 0, time.Local),
				MonthOccurrence: 1,
			},
			{
				Amount:          1000,
				Start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				End:             time.Date(2050, time.December, 1, 0, 0, 0, 0, time.Local),
				MonthOccurrence: 2,
			},
			{
				Amount:          1000,
				Start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				End:             time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				MonthOccurrence: 1,
			},
		},
	}

	amount := budget.Forecast(time.Date(2050, time.January, 1, 0, 0, 0, 0, time.Local))

	var expectedAmount float32 = 149_000

	if amount != expectedAmount {
		t.Errorf("Expected %f to be equal to %f", amount, expectedAmount)
	}
}
