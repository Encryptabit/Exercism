/**
 * The number of minutes it takes to prepare a single layer.
 */
const PREPARATION_MINUTES_PER_LAYER = 2;
export const EXPECTED_MINUTES_IN_OVEN = 40;

export function remainingMinutesInOven(actualMinutesInOven) {
    var remainingMinutes = EXPECTED_MINUTES_IN_OVEN - actualMinutesInOven;

    return  remainingMinutes;
}

export function preparationTimeInMinutes(numberOfLayers) {
  var totalPrepTime =  numberOfLayers * PREPARATION_MINUTES_PER_LAYER

  return totalPrepTime
}

export function totalTimeInMinutes(numberOfLayers, actualMinutesInOven) {
  // 4. Calculate the total working time in minutes
  // sum of preperationTimeInMinutes() and remainingMinutesInOven()

 var total = preparationTimeInMinutes(numberOfLayers) + actualMinutesInOven;

 return total;
}
