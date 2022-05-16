package expenses

import "errors"

//import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {

	matches := make([]Record,0)

	for _ , v := range in {
		if(predicate(v)) {
			matches = append(matches,v)
		}
	}

	return matches
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func (r Record) bool {
		return (p.From <= r.Day && p.To >= r.Day)
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func (r Record) bool {
		return c == r.Category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var runningTotal float64

	filterByPeriod := ByDaysPeriod(p)

	filteredResults := Filter(in, filterByPeriod)

	for _ , v := range filteredResults {
		runningTotal += v.Amount
	}

	return runningTotal
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var runningTotal float64
	var err error

	filterByDays := Filter(in, ByDaysPeriod(p))

	finalFiltered := Filter(filterByDays,ByCategory(c))

	if(len(finalFiltered) > 0) {
		for _ , v := range finalFiltered {
			runningTotal += v.Amount
		}
	} else if len(filterByDays) == 0 {
		// Do nothing to use 0 values
	} else {
		err = errors.New("no results for filter")
	}

	return runningTotal, err
}
