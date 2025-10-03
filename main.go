package finance

func GetFinalCapital(year uint8, interestInPercentage float32, monthlyInvestedCapital float32) (capital float32) {
	for range year {
		var yearlyInvestedAmount float32 = monthlyInvestedCapital * 12
		var interestMultiplier float32 = 1 + (interestInPercentage / 100)

		capital = (capital + yearlyInvestedAmount) * interestMultiplier
	}

	return capital
}
