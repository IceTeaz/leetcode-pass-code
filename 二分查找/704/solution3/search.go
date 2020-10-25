package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] < target:
			left = mid
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] == target:
			return mid
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
}
