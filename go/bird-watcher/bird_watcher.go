package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var total int

	for i := 0; i < len(birdsPerDay); i++ {
		total += birdsPerDay[i]
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	lowerBound := 7 * (week - 1)
	upperBound := 7 * week
	lengthOfBirdsInWeek := len(birdsPerDay)
	var runningTotal int

	if (upperBound >= lengthOfBirdsInWeek) {
		upperBound = lengthOfBirdsInWeek
	}

	for i := lowerBound; i < upperBound; i++ {
		runningTotal += birdsPerDay[i]
	}

	return runningTotal
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i % 2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}

	return birdsPerDay
}
