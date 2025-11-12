package allyourbase

import (
	"reflect"
	"testing"
)

func TestConvertToBase(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := ConvertToBase(tc.inputBase, tc.inputDigits, tc.outputBase)
			if tc.expectedError != "" {
				if err == nil {
					t.Errorf("ConvertToBase(%d, %#v, %d) expected error: %q", tc.inputBase, tc.inputDigits, tc.outputBase, tc.expectedError)
				} else if tc.expectedError != err.Error() {
					t.Errorf("ConvertToBase(%d, %#v, %d)\nexpected error: %q\ngot: %q", tc.inputBase, tc.inputDigits, tc.outputBase, tc.expectedError, err.Error())
				}
			} else if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("ConvertToBase(%d, %#v, %d) = %#v, want:%#v", tc.inputBase, tc.inputDigits, tc.outputBase, actual, tc.expected)
			}
		})
	}
}

func TestConvertToBase16(t *testing.T) {
	tc := struct {
		description   string
		inputBase     int
		inputDigits   []int
		outputBase    int
		expected      []int
		expectedError string
	}{
		description:   "negative digit",
		inputBase:     2,
		inputDigits:   []int{1, -1, 1, 0, 1, 0},
		outputBase:    10,
		expected:      []int(nil),
		expectedError: "all digits must satisfy 0 <= d < input base",
	}

	actual, err := ConvertToBase(tc.inputBase, tc.inputDigits, tc.outputBase)
	if tc.expectedError != "" {
		if err == nil {
			t.Errorf("ConvertToBase(%d, %#v, %d) expected error: %q", tc.inputBase, tc.inputDigits, tc.outputBase, tc.expectedError)
		} else if tc.expectedError != err.Error() {
			t.Errorf("ConvertToBase(%d, %#v, %d)\nexpected error: %q\ngot: %q", tc.inputBase, tc.inputDigits, tc.outputBase, tc.expectedError, err.Error())
		}
	} else if !reflect.DeepEqual(tc.expected, actual) {
		t.Errorf("ConvertToBase(%d, %#v, %d) = %#v, want:%#v", tc.inputBase, tc.inputDigits, tc.outputBase, actual, tc.expected)
	}

}

func TestConvertToBase17(t *testing.T) {
	tc := struct {
		description   string
		inputBase     int
		inputDigits   []int
		outputBase    int
		expected      []int
		expectedError string
	}{
		description:   "invalid positive digit",
		inputBase:     2,
		inputDigits:   []int{1, 2, 1, 0, 1, 0},
		outputBase:    10,
		expected:      []int(nil),
		expectedError: "all digits must satisfy 0 <= d < input base",
	}

	actual, err := ConvertToBase(tc.inputBase, tc.inputDigits, tc.outputBase)
	if tc.expectedError != "" {
		if err == nil {
			t.Errorf("ConvertToBase(%d, %#v, %d) expected error: %q", tc.inputBase, tc.inputDigits, tc.outputBase, tc.expectedError)
		} else if tc.expectedError != err.Error() {
			t.Errorf("ConvertToBase(%d, %#v, %d)\nexpected error: %q\ngot: %q", tc.inputBase, tc.inputDigits, tc.outputBase, tc.expectedError, err.Error())
		}
	} else if !reflect.DeepEqual(tc.expected, actual) {
		t.Errorf("ConvertToBase(%d, %#v, %d) = %#v, want:%#v", tc.inputBase, tc.inputDigits, tc.outputBase, actual, tc.expected)
	}

}
