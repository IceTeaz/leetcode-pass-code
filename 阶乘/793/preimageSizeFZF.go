package main

import "math"

func preimageSizeFZF(K int) int {
	return rightBound(K)-leftBound(K)+1
}


func getZeroesCount(n int) int {
	div:=5
	res:=0
	for ;div<=n;{
		res+=n/div
		div*=5
	}
	return res
}

func leftBound(target int) int {
	left :=0
	right :=math.MaxInt64

	for ; left < right;{
		mid:= left +(right-left)/2
		zeroesCount:=getZeroesCount(mid)
		switch {
		case zeroesCount < target:
			left =mid+1
		case zeroesCount > target:
			right=mid
		case zeroesCount == target:
			right=mid
		}
	}
	return left
}

func rightBound(target int) int {
	left :=0
	right :=math.MaxInt64

	for ; left < right;{
		mid:= left +(right-left)/2
		zeroesCount:=getZeroesCount(mid)

		switch {
		case zeroesCount < target:
			left =mid+1
		case zeroesCount > target:
			right=mid
		case zeroesCount == target:
			left =mid+1
		}
	}
	return right-1
}