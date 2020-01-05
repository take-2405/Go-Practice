package calculate

import (
	"errors"
)

func add(a, b, sum int) error {
	//var err error
	c := a + b
	if sum != c {
		return errors.New("Calculation is wrong. !!")
	}
	return nil
}
