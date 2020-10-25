package main

func main() {

}

func fib(n int) int {
	if n==0{
		return 0
	}
	if n==2|| n==1{
		return 1
	}
	prev:=1
	cur:=1
	for i:=3;i<=n;i++{
		sum:=prev+cur
		prev=cur
		cur=sum
	}
	return cur
}