package models

import "testing"

func TestValidPhoneNumber(t *testing.T) {
	var tests = []struct {
		input          string
		expectedOutput bool
	}{
		{"9883443344", true},
		{"09883443344", true},
		{"919883443344", true},
		{"0919883443344", true},
		{"+919883443344", true},
		{"+91-9883443344", true},
		{"0091-9883443344", true},
		{"+91 -9883443344", true},
		{"+91 - 9883443344", true},
		{"0091 - 9883443344", true},
		{"23", false},
		{"", false},
	}
	for _, e := range tests {
		res := isValidMobileNumber(e.input)
		if res != e.expectedOutput {
			t.Errorf("isValidMobileNumber(%v) = %v, expected %v",
				e.input, res, e.expectedOutput)
		}
	}
}

func TestValidUserName(t *testing.T) {
	var tests = []struct {
		input          string
		expectedOutput bool
	}{
		{"Mr. ABC XYZ", true},
		{"MR. ABC", true},
		{"ABC", true},
		{"abc", true},
		{"Abc", true},
		{"abc xyz pqr", true},
		{"ABC123", false},
		{"ABC_XYZ", false},
		{"ABC-XYZ", false},
		{"ABC@XYZ", false},
		{"123ABC", false},
		{"", false},
	}
	for _, e := range tests {
		res := isValidUserName(e.input)
		if res != e.expectedOutput {
			t.Errorf("isValidUserName(%v) = %v, expected %v",
				e.input, res, e.expectedOutput)
		}
	}
}
