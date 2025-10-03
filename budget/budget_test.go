package budget

import (
	"testing"
	"time"
)

func TestBudget(t *testing.T) {
	budget := Budget{
		start: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.Local),
		incomes: []Income{
			{
				amount:          2000,
				start:           time.Date(2025, time.April, 1, 0, 0, 0, 0, time.Local),
				monthOccurrence: 1,
			},
		},
		expenses: []Expense{
			{
				amount:          1000,
				start:           time.Date(2025, time.April, 1, 0, 0, 0, 0, time.Local),
				monthOccurrence: 1,
			},
		},
	}

	amount := budget.GetForecast(time.Date(2050, time.April, 1, 0, 0, 0, 0, time.Local))

	var expectedAmount float32 = 300_000

	if amount != expectedAmount {
		t.Errorf("Expected %f to be equal to %f", amount, expectedAmount)
	}
}
