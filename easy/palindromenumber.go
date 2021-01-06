package easy

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}

	reversed := 0
	for {
		pop := x % 10
		reversed = reversed*10 + pop
		x /= 10

		if reversed == x || reversed == x*10+pop {
			return true
		}
		if reversed > x || x == 0 {
			break
		}
	}
	return false
}
