package leetcode_array

// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//在「杨辉三角」中，每个数是它左上方和右上方的数的和
// https://leetcode.cn/problems/pascals-triangle

// 1 <= numRows <= 30
// a, 每一排的第一个数和最后一个数都是 1，即 c[i][0]=c[i][i]=1
// b, 其余数字等于左上方的数，加上正上方的数，即 c[i][j]=c[i−1][j−1]+c[i−1][j]
func generate(numRows int) [][]int {
	var c = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		c[i] = make([]int, i+1)
		c[i][0] = 1
		c[i][i] = 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
	return c
}

// 1 <= numRows <= 30
func generate2(numRows int) [][]int {
	var rows = make([][]int, 0, numRows)
	rows = append(rows, []int{1}) // first row
	for i := 1; i < numRows; i++ {
		var prev = rows[i-1]
		var row = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j >= 1 && j <= len(prev) {
				row[j] = prev[j-1]
			}
			if j >= 0 && j < len(prev) {
				row[j] += prev[j]
			}
		}
		rows = append(rows, row)
	}
	return rows
}
