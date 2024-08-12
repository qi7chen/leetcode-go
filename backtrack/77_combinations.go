package leetcode_backtrack

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合
// https://leetcode.cn/problems/combinations/
func combine(n int, k int) [][]int {
	var ctx = &combineContext{
		n: n,
		k: k,
	}
	backtrack(ctx, 1)
	return ctx.result
}

type combineContext struct {
	result [][]int
	path   []int
	n      int
	k      int
}

func backtrack(ctx *combineContext, start int) {
	if len(ctx.path) == ctx.k {
		var clone = append([]int{}, ctx.path...)
		ctx.result = append(ctx.result, clone)
		return
	}
	for i := start; i <= ctx.n; i++ {
		ctx.path = append(ctx.path, i)
		backtrack(ctx, i+1)
		ctx.path = ctx.path[:len(ctx.path)-1] // 回溯
	}
}
