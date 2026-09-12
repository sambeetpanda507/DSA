package main

import "math"

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	ans := math.MaxInt
	for left <= right {
		var mid int = left + (right-left)/2
		if nums[mid] < ans {
			ans = nums[mid]
		}

		if nums[left] <= nums[mid] {
			if nums[left] > nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return ans
}
