# gofin

Finance functions for the Go language

## Installation

```bash
go get github.com/aminnairi/gofin
```

## Usage

### GetFinalCapital

```go
func GetFinalCapital(
  year uint8,
  interestInPercentage float32,
  monthlyInvestedCapital float32
) float32
```

```go
package main

import (
  finance "github.com/aminnairi/gofin"
)

func main() {
  capital := finance.GetFinalCapital(5, 7, 100)
}
```