package math

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("invalid input f=[%f]", f)
	}
	var high float64
	var low float64
	high = f
	low = 0.0
	guess := low + (high-low)/2
	for math.Abs(f-guess*guess) > 0.000001 {
		fmt.Printf("high=[%10.9f], guess=[%10.9f], guess^2=[%10.9f]\t", high, guess, guess*guess)
		fmt.Printf("high-guess*guess=[%10.9f]\n", math.Abs(high-guess))
		guess = low + (high-low)/2
		if guess*guess > f {
			high = guess
		} else {
			low = guess
		}
	}
	return guess, nil
}
