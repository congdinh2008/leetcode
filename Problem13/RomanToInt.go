package main

func main(){
	s:="MCMXCIV"
	println(romanToInt(s))
}

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1,
		"IV": 4,
		"V": 5,
		"IX": 9,
		"X": 10,
		"XL": 40,
		"L": 50,
		"XC": 90,
		"C": 100,
		"CD": 400,
		"D": 500,
		"CM": 900,
		"M": 1000,
	}

	result := 0
	
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			if value, ok := romanMap[s[i:i+2]]; ok {
				result += value
				i++
				continue
			}
		}
		result += romanMap[s[i:i+1]]
	}
	
	return result
}