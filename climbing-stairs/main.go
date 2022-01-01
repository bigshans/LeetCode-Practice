package main

import "fmt"

func climbStairs(n int) int {
	a := []int{0, 1, 2, 3}
	for i := 4; i <= n; i += 1 {
		a = append(a, a[i-1]+a[i-2])
	}
	return a[n]
}

func main() {
	fmt.Println(climbStairs(45))
}
