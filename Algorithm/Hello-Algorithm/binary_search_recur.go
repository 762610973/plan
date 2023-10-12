package Hello_Algorithm

/* 二分查找：问题 f(i, j) */
func dfsRecur(nums []int, target, i, j int) int {
	// 如果区间为空，代表没有目标元素，则返回 -1
	if i > j {
		return -1
	}
	//    计算索引中点
	m := i + ((j - i) >> 1)
	//判断中点与目标元素大小
	if nums[m] < target {
		// 小于则递归右半数组
		// 递归子问题 f(m+1, j)
		return dfsRecur(nums, target, m+1, j)
	} else if nums[m] > target {
		// 小于则递归左半数组
		// 递归子问题 f(i, m-1)
		return dfsRecur(nums, target, i, m-1)
	} else {
		// 找到目标元素，返回其索引
		return m
	}
}

/* 二分查找 */
func binarySearch(nums []int, target int) int {
	n := len(nums)
	return dfsRecur(nums, target, 0, n-1)
}
