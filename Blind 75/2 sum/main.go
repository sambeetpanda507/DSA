package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := []int{}
	hash := make(map[int]int)
	for i, n := range nums {
		req := target - n
		if val, ok := hash[req]; ok {
			result = append(result, []int{val, i}...)
			break
		} else {
			hash[n] = i
		}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
