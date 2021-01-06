package easy

func romanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	result := 0
	for i:=0; i< len(s); i++{
		if i < len(s) - 1{
			r := s[i:i+2]
			if v, ok := m[r]; ok {
				result += v
				i++
			}else{
				result += m[string(s[i])]
			}
		}else{
			result += m[string(s[i])]
		}

	}
	return result
}
