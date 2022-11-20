package main

import "fmt"

func isValid(s string) bool {
	var bType1, bType2, bType3 bool = true, true, true
	for _, v := range s {
		switch string(v) {
		case "(":
			bType1 = !bType1
		case ")":
			bType1 = !bType1
		case "[":
			bType2 = !bType2
		case "]":
			bType2 = !bType2
		case "{":
			bType3 = !bType3
		case "}":
			bType3 = !bType3
		}
		if bType1 && bType2 && bType3 {
			return true
		}
	}
	return false
}

func main() {
	s1 := "()()[][]{}{}"
	s2 := "("
	s3 := "([)]"
	result1 := isValid(s1)
	result2 := isValid(s2)
	result3 := isValid(s3)
	fmt.Println(result1, result2, result3)
}
