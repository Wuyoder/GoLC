package main

import "fmt"

var nums = []int{2, 7, 11, 15}
var target = 9

func twoSum(nums []int, target int) []int {
	x := []int{}
	var remain int
	for i := 0; i < len(nums); i++ {
		remain = target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == remain {
				x = append(x, i, j)
			}
		}
	}
	return x
}

func main() {
	fmt.Println(twoSum(nums, target)) // [0 ,1]

}
