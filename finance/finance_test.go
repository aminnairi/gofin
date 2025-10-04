package finance

import (
	"math"
	"testing"
)

func TestGetCapitalAt(t *testing.T) {
	capital := GetFinalCapital(38, 10, 500)

	if capital != 2402688.5 {
		t.Errorf("Capital invested should be 2402688,5, %f computed", capital)
	}
}

func TestTotalCreditCost(t *testing.T) {
	expectedCost := 36886.0
	expectedRate := 25.0
	cost, rate := TotalCreditCost(150_000, 20*12, 1.75, 0.3)
	roundedRate := math.Round(rate * 100)
	roundedCost := math.Round(cost)

	if roundedCost != expectedCost {
		t.Errorf("Expected credit cost to be %f, %f received.", expectedCost, roundedCost)
	}

	if roundedRate != expectedRate {
		t.Errorf("Expected credit rate to be %f%%, %f%% received.", expectedRate, roundedRate)
	}
}
