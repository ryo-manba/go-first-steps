package main

import "fmt"

func fibonacci(n int) (ans []int) {
	if n < 2 {
		return make([]int, 0)
	}
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			ans = append(ans, 1)
		} else {
			ans = append(ans, ans[i-1] + ans[i-2])
		}
	}
	return
}

func main() {
	ans := fibonacci(10)
	fmt.Println(ans)
}
