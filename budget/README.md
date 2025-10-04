# gofin/budget

## Usage

### Expense

```go
type Expense struct {
	amount          float32
	start           time.Time
	monthOccurrence int8
}
```

An amount that starts at a specific time, with an occurrence in month, and that will decrease the amount of a budget.

```go
package main

import (
  "fmt"
  "time"

  "github.com/aminnairi/gofin/budget"
)

func main() {
  expense := budget.Expense{
      {
				amount:          1000,
				start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				monthOccurrence: 1,
      }
  }

  fmt.Println(expense)
}
```

### Income

```go
type Income struct {
	amount          float32
	start           time.Time
	monthOccurrence int8
}
```

An amount that starts at a specific time, with an occurrence in month, and that will increase the amount of a budget.

```go
package main

import (
  "fmt"
  "time"

  "github.com/aminnairi/gofin/budget"
)

func main() {
  income := budget.Expense{
      {
				amount:          1000,
				start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				monthOccurrence: 1,
      }
  }

  fmt.Println(income)
}
```

### Budget

```go
type Budget struct {
	start    time.Time
	expenses []Expense
	incomes  []Income
}
```

A budget is a grouping of expenses and incomes, that starts at a specific date.

```go
package main

import (
  "fmt"
  "time"

  "github.com/aminnairi/gofin/budget"
)

func main() {
  budget := budget.Budget{
    start: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Local),
    incomes: []budget.Income{
      {
				amount:          2000,
				start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				monthOccurrence: 1,
      }
    },
    expenses: []budget.Expense{
      {
				amount:          1000,
				start:           time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
				monthOccurrence: 1,
      }
    },
  }
}
```

### Forecast

```go
func (budget Budget) Forecast(date time.Time) (amount float32)
```

Let's you predict your budget amount for a given date. If the date provided is anterior to the budget start date, then the amount will be `0`.

```go
package main

import (
  "fmt"

  "github.com/aminnairi/gofin/budget"
)

func main() {
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

	amount := budget.Forecast(time.Date(2050, time.April, 1, 0, 0, 0, 0, time.Local))

  fmt.Println(amount)
}
```