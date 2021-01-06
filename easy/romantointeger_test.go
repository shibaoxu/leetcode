package easy

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	cases := []struct{
		input string
		want int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for i, c := range cases{
		if result := romanToInt(c.input); result != c.want{
			t.Errorf("case %d: input = %v want = %d actual = %d", i, c.input, c.want, result)
		}
	}
}