package easy

import (
	"testing"
)


func TestReverseInteger(t *testing.T) {
	cases := []struct {
		input, output int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
	}
	for i, cs := range cases{
		if result := reverseInteger(cs.input); result != cs.output{
			t.Errorf("case %d: input = %d want = %d actual = %d", i, cs.input, cs.output, result)
		}
	}
}
