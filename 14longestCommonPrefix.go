package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	lens := len(strs[0])
	var lcp string
	for i := 1; i <= lens; i++ {
		count := 0
		for _, v := range strs {
			if v[:i] == strs[0][:i] {
				count += 1
			}
		}
		if count == len(strs) {
			lcp = strs[0][:i]
		} else {
			return lcp
		}
	}
	return lcp
}

func main() {
	strs := []string{"a"}
	fmt.Println(longestCommonPrefix(strs)) // fl
}
