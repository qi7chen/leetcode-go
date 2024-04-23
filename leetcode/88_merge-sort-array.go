package leetcode

// https://leetcode.cn/problems/merge-sorted-array
func merge(x []int, m int, y []int, n int) {
	var i, j int
	for i < m && j < n {
		if x[i] < y[j] {
			insertAt(x, i+1, y[j])
			i++
		} else if x[i] > y[j] {
			insertAt(x, i, y[j])
			j++
		} else {
			i++
			j++
		}
	}
}

// insertAt 把`v`插入到第`i`个位置
func insertAt(s []int, i int, v int) {

}
