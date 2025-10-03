package finance

import "testing"

func TestGetCapitalAt(t *testing.T) {
	capital := GetFinalCapital(38, 10, 500)

	if capital != 2402688.5 {
		t.Errorf("Capital invested should be 2402688,5, %f computed", capital)
	}
}
