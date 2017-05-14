package scaler

import (
	"fmt"
)

// Scaler scales an input to the output.
// e.g. for an input of 0-10 and output of 0-100, an
// input of 5 would produce an output of 50.
type Scaler struct {
	InputMinimum  int
	InputMaximum  int
	OutputMinimum int
	OutputMaximum int
}

// New creates a utility to scale input values to output values.
func New(inputMinimum, inputMaximum int, outputMinimum, outputMaximum int) Scaler {
	return Scaler{
		InputMinimum:  inputMinimum,
		InputMaximum:  inputMaximum,
		OutputMinimum: outputMinimum,
		OutputMaximum: outputMaximum,
	}
}

// Scale calculates the output value for the input.
func (s Scaler) Scale(in float64) (float64, error) {
	if in < float64(s.InputMinimum) {
		return float64(s.OutputMinimum), fmt.Errorf("scale: input %v was not within range %v to %v", in, s.InputMinimum, s.InputMaximum)
	}

	if in > float64(s.InputMaximum) {
		return float64(s.OutputMaximum), fmt.Errorf("scale: input %v was not within range %v to %v", in, s.InputMinimum, s.InputMaximum)
	}

	outputRange := float64(s.OutputMaximum - s.OutputMinimum)
	inputRange := float64(s.InputMaximum - s.InputMinimum)

	return ((outputRange / inputRange) * (in - float64(s.InputMinimum))) + float64(s.OutputMinimum), nil
}

// Invert calculates the input value from the output.
func (s Scaler) Invert(in float64) (float64, error) {
	if in < float64(s.OutputMinimum) {
		return float64(s.InputMinimum), fmt.Errorf("invert: input %v was not within range %v to %v", in, s.OutputMinimum, s.OutputMaximum)
	}

	if in > float64(s.OutputMaximum) {
		return float64(s.InputMaximum), fmt.Errorf("invert: input %v was not within range %v to %v", in, s.OutputMinimum, s.OutputMaximum)
	}

	outputRange := float64(s.InputMaximum - s.InputMinimum)
	inputRange := float64(s.OutputMaximum - s.OutputMinimum)

	return ((outputRange / inputRange) * (in - float64(s.OutputMinimum))) + float64(s.InputMinimum), nil
}
