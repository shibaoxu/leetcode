package easy

import (
	"shibaoxu.cn/leetcode/utils"
	"testing"
)

var cases = []struct{
	input []int
	target int
	result []int
}{
	{[]int{2,7,11,5}, 9, []int{0,1}},
	{[]int{3,2,4}, 6, []int{1,2}},
	{[]int{3,3}, 6, []int{0,1}},
}

func TestTwoSumLoop(t *testing.T) {
	for i, cs := range cases {
		if result:= twoSumLoop(cs.input, cs.target); !utils.CompareSlice(cs.result, result){
			t.Errorf("case %d: input = %v target = %d want = %v actusl = %v", i, cs.input, cs.target, cs.result, result)
		}
	}
}

func TestTwoSumHash(t *testing.T)  {
	for i, cs := range cases {
		if result:= twoSumHash(cs.input, cs.target); !utils.CompareSlice(cs.result, result){
			t.Errorf("case %d: input = %v target = %d want = %v actusl = %v", i, cs.input, cs.target, cs.result, result)
		}
	}
}