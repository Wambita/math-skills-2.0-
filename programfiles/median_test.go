package mathskills

import "testing"

// testing for median function , compare  output to expected input
func TestCalcMedian(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{input: []float64{101, 242, 999, 98, 456}, expected: 242},
		{input: []float64{100, 200, 500, 300, 400,700}, expected: 300},
	}

	for _, test := range tests {
		output := CalcMedian(test.input)
		if output != test.expected {
			t.Errorf("calculate median(%v) = %v; Expected %v", test.input, output, test.expected)
		}
	}
}
