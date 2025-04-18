package _15_prop_tests_

import "testing"

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Description string
		Input       int
		Expected    string
	}{
		{
			Description: "1 should convert to I",
			Input:       1,
			Expected:    "I",
		},
		{
			Description: "2 should convert to II",
			Input:       2,
			Expected:    "II",
		},
		{
			Description: "3 should convert to III",
			Input:       3,
			Expected:    "III",
		},
		{
			Description: "4 gets converted to IV (can't repeat more than 3 times)",
			Input:       4,
			Expected:    "IV",
		},
		{
			Description: "5 gets converted to V",
			Input:       5,
			Expected:    "V",
		},
		{
			Description: "9 gets converted to IX",
			Input:       9,
			Expected:    "IX",
		},
	}
	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			want := c.Expected
			got := ConvertToRoman(c.Input)
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
