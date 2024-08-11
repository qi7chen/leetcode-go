package leetcode

// https://leetcode.cn/problems/add-binary
func addBinary(a string, b string) string {
	var n = max(len(a), len(b)) + 1
	var sum = make([]byte, n)
	for i := 0; i < n; i++ {
		sum[i] = '0'
	}
	var i = len(a) - 1
	var j = len(b) - 1
	var k = len(sum) - 1
	for i >= 0 || j >= 0 {
		var left = i >= 0 && a[i] == '1'
		var right = j >= 0 && b[j] == '1'
		if left && right {
			sum[k-1] += 1
		} else {
			if left || right {
				sum[k] += 1
			}
		}
		if sum[k] >= '2' {
			sum[k-1] += 1
			sum[k] -= 2
		}
		i--
		j--
		k--
	}
	// remove prefix '0'
	for k+1 < n && sum[k] == '0' {
		k++
	}
	return btos(sum[k:])
}
