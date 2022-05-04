package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int{
	var multiplier int

	if averagePrepTime == 0 {
		multiplier = 2
	} else {
		multiplier = averagePrepTime
	}

	return len(layers) * multiplier
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
	var runningSauceTotal float64
	var runningNoodleTotal int

	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			runningSauceTotal +=  0.2
		} else if layers[i] == "noodles" {
			runningNoodleTotal += 50
		}
	}

	return runningNoodleTotal, runningSauceTotal
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList = append(myList[0:len(myList)-1],friendsList[len(friendsList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	copyOfOriginal := make([]float64, len(quantities))

	copy(copyOfOriginal, quantities)

	// Halve quantities
	for j := 0; j < len(copyOfOriginal); j++ {
		copyOfOriginal[j] = (copyOfOriginal[j] * 0.5) * float64(portions)
	}

	return copyOfOriginal
}