package easy

func twoSumLoop(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++{
		for j:= i + 1; j <len(nums); j++{
			if nums[i] + nums[j] == target{
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumHash(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, v := range nums{
		complement := target - v
		if index, ok := cache[complement]; ok{
			return []int{index, i}
		}else{
			cache[v] = i
		}
	}
	return nil
}