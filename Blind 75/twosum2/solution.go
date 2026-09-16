package main

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	result := []int{}
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			result = append(result, []int{left + 1, right + 1}...)
			break
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return result
}
