package scaler

import (
	"strings"
	"testing"
)

func TestScaling(t *testing.T) {
	tests := []struct {
		name            string
		s               Scaler
		input           float64
		expected        float64
		expectedMessage string
	}{
		{
			name:     "Scale 0-10 onto 0-100",
			s:        New(0, 10, 0, 100),
			input:    5.0,
			expected: 50,
		},
		{
			name:     "Scale 0-10 onto 0-100",
			s:        New(-10, 10, 0, 100),
			input:    0.0,
			expected: 50,
		},
		{
			name:     "Shift the range",
			s:        New(1, 11, 0, 10),
			input:    11.0,
			expected: 10,
		},
		{
			name:     "Scale a small input to large output",
			s:        New(0, 1, -100, 100),
			input:    0,
			expected: -100,
		},
		{
			name:     "Scale a small input to large output",
			s:        New(0, 1, -100, 100),
			input:    0.5,
			expected: 0.0,
		},
		{
			name:            "Less than the input range",
			s:               New(0, 1, -100, 100),
			input:           -1,
			expected:        -100,
			expectedMessage: "scale: input -1 was not within range 0 to 1",
		},
		{
			name:            "Exceed the input range",
			s:               New(0, 1, -100, 100),
			input:           2,
			expected:        100,
			expectedMessage: "scale: input 2 was not within range 0 to 1",
		},
	}

	for _, test := range tests {
		actual, err := test.s.Scale(test.input)

		if actual != test.expected {
			t.Errorf("%s: for %v and input %v, expected %v, but got %v", test.name, test.s, test.input, test.expected, actual)
		}

		if err != nil {
			if test.expectedMessage == "" {
				t.Errorf("%s: unexpected error %v", test.name, err)
			}
			if !strings.HasPrefix(err.Error(), test.expectedMessage) {
				t.Errorf("%s: unexpected error to start with prefix '%v' but got '%v'", test.name, test.expectedMessage, err)
			}
			continue
		}

		inverse, err := test.s.Invert(actual)

		if inverse != test.input {
			t.Errorf("%s: for %v, with inverse input %v, expected %v, but got %v", test.name, test.s, actual, test.input, inverse)
		}

		if err != nil && test.expectedMessage == "" {
			t.Errorf("%s: unexpected error %v", test.name, err)
		}
		if err != nil && !strings.HasPrefix(err.Error(), test.expectedMessage) {
			t.Errorf("%s: unexpected error to start with prefix '%v' but got '%v'", test.name, test.expectedMessage, err)
		}
	}
}

func TestInvert(t *testing.T) {
	tests := []struct {
		name            string
		s               Scaler
		input           float64
		expected        float64
		expectedMessage string
	}{
		{
			name:     "Scale 0-10 onto 0-100",
			s:        New(0, 10, 0, 100),
			input:    50,
			expected: 5,
		},
		{
			name:            "Less than the input range",
			s:               New(0, 1, -100, 100),
			input:           -1000,
			expected:        0,
			expectedMessage: "invert: input -1000 was not within range -100 to 100",
		},
		{
			name:            "Exceed the input range",
			s:               New(0, 1, -100, 100),
			input:           1000,
			expected:        1,
			expectedMessage: "invert: input 1000 was not within range -100 to 100",
		},
	}

	for _, test := range tests {
		actual, err := test.s.Invert(test.input)

		if actual != test.expected {
			t.Errorf("%s: for %v and input %v, expected %v, but got %v", test.name, test.s, test.input, test.expected, actual)
		}

		if err != nil {
			if test.expectedMessage == "" {
				t.Errorf("%s: unexpected error %v", test.name, err)
			}
			if !strings.HasPrefix(err.Error(), test.expectedMessage) {
				t.Errorf("%s: unexpected error to start with prefix '%v' but got '%v'", test.name, test.expectedMessage, err)
			}
			continue
		}
	}
}
