package main

import "fmt"

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// IV IX
// XL XC
// CD CM

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

func main() {
	fmt.Println(romanToInt("CMLII"))
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("LVIII"))

}
