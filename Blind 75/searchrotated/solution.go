package main

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] == target {
			return left
		}

		if nums[right] == target {
			return right
		}

		// Check left sorted
		if nums[left] < nums[mid] {
			// Check if the target is in the left half
			if target > nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Check if the target is in the right half
			if target > nums[mid] && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
