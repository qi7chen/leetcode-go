package leetcode

// 斐波那契数
// https://leetcode.cn/problems/fibonacci-number/
func fib(n int) int {
	var a = 0
	var b = 1
	var sum int
	for i := 0; i < n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return a
}
