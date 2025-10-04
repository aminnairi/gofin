package finance

import "math"

func GetFinalCapital(year uint8, interestInPercentage float32, monthlyInvestedCapital float32) (capital float32) {
	for range year {
		var yearlyInvestedAmount float32 = monthlyInvestedCapital * 12
		var interestMultiplier float32 = 1 + (interestInPercentage / 100)

		capital = (capital + yearlyInvestedAmount) * interestMultiplier
	}

	return capital
}

func TotalCreditCost(
	loanAmount float64,
	durationInMonths int,
	annualRateInPercentage float64,
	annualInsuranceFeeInPercentage float64,
) (cost float64, rate float64) {
	monthlyRate := annualRateInPercentage / 100 / 12
	duration := float64(durationInMonths)
	powerTerm := math.Pow(1+monthlyRate, -duration)
	denominator := 1 - powerTerm

	if monthlyRate == 0 || denominator == 0 {
		cost = 0
	} else {
		monthlyAnnuity := loanAmount * (monthlyRate / denominator)
		totalRepaid := monthlyAnnuity * duration
		cost = totalRepaid - loanAmount
	}

	insuranceFee := loanAmount * (annualInsuranceFeeInPercentage / 100) * (duration / 12)
	cost += insuranceFee
	rate = cost / loanAmount

	return cost, rate
}
