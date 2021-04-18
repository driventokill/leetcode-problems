// Package _509_fibonacci_number https://leetcode-cn.com/problems/fibonacci-number/
package _509_fibonacci_number

func fib(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 || n == 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

// Remember previous result by map
func fib1(n int) int {
	if n < 0 {
		return -1
	}

	m := make(map[int]int)
	return fibMemo(m, n)
}

func fibMemo(memo map[int]int, n int) int {
	if n == 0 || n ==1 {
		return n
	}
	// query from fibMemo
	if memo[n] != 0 {
		return memo[n]
	} else {
		memo[n] = fibMemo(memo, n-1) + fibMemo(memo, n-2)
		return memo[n]
	}
}

// Remember only last 2 value, bottom up
func fib2(n int) int {
	if n < 1 {
		return 0
	}

	if n == 0 || n ==1 {
		return n
	}

	prev, curr := 1, 1
	for i := 3; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}

	return curr
}