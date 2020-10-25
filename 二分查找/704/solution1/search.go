package main

import "fmt"

func search(nums []int, target int) int {
	left:=0
	right:=len(nums)-1

	for left<=right{
		mid:=left+(right-left)/2
		switch {
		case nums[mid]<target:
			left=mid+1
		case nums[mid]>target:
			right=mid-1
		case nums[mid]==target:
			return mid
		}
	}
	return -1
}

func main() {
	nums:=[]int{-1,0,3,5,9,12}
	fmt.Println(search(nums,9))
}