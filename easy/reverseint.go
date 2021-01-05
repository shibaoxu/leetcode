package easy


func reverseInteger(x int) int {
	const (
		INT_MAX = int(int32(^uint32((0)) >> 1))
		INT_MIN = ^INT_MAX
	)
	reversed := 0
	for x != 0{
		pop := x % 10
		if reversed > INT_MAX/10 || (reversed == INT_MAX/10 && pop > 7) {
			return 0
		}

		if reversed < INT_MIN/10 || (reversed == INT_MIN && pop < -8) {
			return 0
		}
		reversed = reversed*10 + pop
		x /= 10
	}
	return reversed
}
