package leetcode_array

func binarySearch(x []int, cmp func(int) bool) int {
	var left = 0
	var right = len(x)
	for left < right {
		h := int(uint(left+right) >> 1) // avoid overflow
		if !cmp(h) {
			left = h + 1
		} else {
			right = h
		}
	}
	return left
}

func binaryFind(x []int, cmp func(int) int) (int, bool) {
	var left = 0
	var right = len(x)
	for left < right {
		h := int(uint(left+right) >> 1) // avoid overflow
		if cmp(h) > 0 {
			left = h + 1
		} else {
			right = h
		}
	}
	return left, left < len(x) && cmp(left) == 0
}

// leftBound 寻找左侧边界，`x`中有多少个数小于`target`
func leftBound(x []int, target int) int {
	var left = 0
	var right = len(x)
	for left < right {
		h := int(uint(left+right) >> 1) // avoid overflow
		if x[h] < target {
			left = h + 1
		} else {
			right = h // x[h] >= target
		}
	}
	return left
}

// rightBound = 寻找右侧边界,`x`中有多少个数小于等于`target`
func rightBound(x []int, target int) int {
	var left = 0
	var right = len(x)
	for left < right {
		h := int(uint(left+right) >> 1) // avoid overflow
		if x[h] <= target {
			left = h + 1
		} else {
			right = h // x[h] > target
		}
	}
	return max(0, left-1)
}
