package easy

import "testing"

func TestIsPalinding(t *testing.T) {
	cases := []struct{
		input int
		want bool
	}{
		{10, false},
	}

	for i, cs := range cases{
		if isPalindrome(cs.input) != cs.want{
			t.Errorf("case %d: input = %d, want = %v, actual = %v", i, cs.input, cs.want, !cs.want)
		}
	}
}