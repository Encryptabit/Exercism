package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0 : return 3.213
	case balance >= 0 && balance < 1000 : return .5
	case balance >= 1000 && balance < 5000 : return 1.621
	case balance >= 5000 : return 2.475 
	}

	return 0
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	panic("Please implement the YearsBeforeDesiredBalance function")
}