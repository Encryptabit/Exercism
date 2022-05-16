package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string { return "silly nephew, there cannot be " + fmt.Sprint(e.cows) + " cows"}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()

	if (err == ErrScaleMalfunction && fodderAmount > 0) {
		fodderAmount = (fodderAmount * 2) / float64(cows) 

		return fodderAmount, nil

	}

	if ((fodderAmount < 0 && err == ErrScaleMalfunction) || fodderAmount < 0 && err == nil) {
		return 0, errors.New("negative fodder")
	}

	if (fodderAmount < 0) {
		return 0, err
	}

	if (cows == 0) {
		return 0, errors.New("division by zero")
	}

	if (cows < 0) {
		return 0, &SillyNephewError{cows}
	}

	
	if err != nil  {
		return 0, err
	}
	
	fodderAmount = fodderAmount / float64(cows)


	return fodderAmount,err
}
