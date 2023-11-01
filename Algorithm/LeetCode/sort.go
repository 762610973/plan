package main

import "fmt"

func partition(nums []int, left, right int) int {
	i, j := left+1, right
	for i < j {
		if nums[i] > nums[left] {
			// 当前元素大于基准数, 放到最后, 同时j--
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	// 如果基准数小于要拆分数, i--, 便于下面交换元素
	if nums[left] <= nums[i] {
		i--
	}
	nums[left], nums[i] = nums[i], nums[left]

	return i
}

func quickSort(nums []int, left, right int) {
	if left < right {
		l := partition(nums, left, right)
		quickSort(nums, left, l-1)
		quickSort(nums, l+1, right)
	}
}

func main() {
	rawNums := []int{1, 7, 4, 8, 3, 5, 9}
	quickSort(rawNums, 0, len(rawNums)-1)
	fmt.Println(rawNums)
}
