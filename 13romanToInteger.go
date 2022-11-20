package main

func romanToInt(s string) int {
	var sum int
	var remain string = s
	trans := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	for i := len(s) - 1; i > 0; {
		if _, ok := trans[remain[i-1:]]; ok {
			sum += trans[remain[i-1:]]
			remain = remain[:i-1]
			i -= 2
		} else {
			sum += trans[remain[i:]]
			remain = remain[:i]
			i -= 1
		}
	}
	sum += trans[remain]
	return sum
}

// func main() {
// 	fmt.Println(romanToInt("CMLII")) // 952
// 	fmt.Println(romanToInt("MCMXCIV")) // 1994
// 	fmt.Println(romanToInt("LVIII")) // 58

// }
