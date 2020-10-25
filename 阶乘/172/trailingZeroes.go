package _72

func trailingZeroes(n int) int {
	div:=5
	res:=0
	for ;div<=n;{
		res+=n/div
		div*=5
	}
	return res
}