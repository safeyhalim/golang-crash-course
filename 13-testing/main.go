package main

import "errors"

func main() {

}

func Add(num1, num2 float64) (float64, error) { // Returns the sum and a placeholder of an error
	result := num1 + num2
	return result, nil
}

func Divide(dividend, divisor float64) (float64, error) {
	var result float64
	// cannot divide by 0
	if divisor == 0 {
		return result, errors.New("CANNOT DIVIDE BY 0") // Standard library errors package (like exeption handling?)
	}

	result = dividend / divisor
	return result, nil

}
