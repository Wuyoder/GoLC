package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	for i, j := 0, len(strconv.Itoa(x))-1; i < j; i, j = i+1, j-1 {
		if strconv.Itoa(x)[i] != strconv.Itoa(x)[j] {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(isPalindrome(123))  // false
// 	fmt.Println(isPalindrome(-121)) //false
// 	fmt.Println(isPalindrome(121))  // true

// }
