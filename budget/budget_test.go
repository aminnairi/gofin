package budget

import (
	"testing"
	"time"
)

func TestBudgetForecast(t *testing.T) {
	budget := Budget{
		start: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
		incomes: []Income{
			{
				amount:          2000,
				start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				end:             time.Date(2050, time.December, 1, 0, 0, 0, 0, time.Local),
				monthOccurrence: 1,
			},
		},
		expenses: []Expense{
			{
				amount:          1000,
				start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				end:             time.Date(2050, time.December, 1, 0, 0, 0, 0, time.Local),
				monthOccurrence: 1,
			},
			{
				amount:          1000,
				start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				end:             time.Date(2050, time.December, 1, 0, 0, 0, 0, time.Local),
				monthOccurrence: 2,
			},
			{
				amount:          1000,
				start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				end:             time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				monthOccurrence: 1,
			},
		},
	}

	amount := budget.Forecast(time.Date(2050, time.January, 1, 0, 0, 0, 0, time.Local))

	var expectedAmount float32 = 149_000

	if amount != expectedAmount {
		t.Errorf("Expected %f to be equal to %f", amount, expectedAmount)
	}
}
