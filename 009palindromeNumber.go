package main

import (
	"fmt"
	"strconv"
)

func palindromeNumber(x int) bool {
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

func main() {
	fmt.Println(palindromeNumber(123))  // false
	fmt.Println(palindromeNumber(-121)) //false
	fmt.Println(palindromeNumber(121))  // true

}
