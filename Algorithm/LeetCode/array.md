- [binary-search](https://leetcode.cn/problems/binary-search/)
```go
package main
// 二分查找, 有序数组搜索目标值target, 存在返回下标, 否则返回-1
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right =  mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
```
- [search-insert-position](https://leetcode.cn/problems/search-insert-position/)
```go
package main
// 搜索插入位置, 在排序数组中, 目标值存在则返回索引, 不存在则返回插入的位置
func searchInsert(nums []int, target int) int {
    left, right := 0,len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
    // left那个位置的值在右, right在左.

	return left
}
```
- [find-first-and-last-position-of-element-in-sorted-array](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)

```go
package main
// 非递减顺序搜索target的开始和结束位置
func searchRange(nums []int, target int) []int {
	leftFn := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)>>1
			if target < nums[mid] {
				right = mid - 1
			} else if target > nums[mid] {
				left = mid + 1
			} else {
                // 第一个就是target或者左边不再有相同的元素
				if mid == 0 || nums[mid] != nums[mid-1] {
					return mid
				} else {
					right = mid - 1
				}
			}
		}
		return -1
	}
	rightFn := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)>>1
			if target > nums[mid] {
				left = mid + 1
			} else if target < nums[mid] {
				right = mid - 1
			} else {
                // 最后一个是target或者后面再没有target了
				if mid == len(nums)-1 || nums[mid] != nums[mid+1] {
					return mid
				} else {
					left = mid + 1
				}
			}
		}
		return -1
	}
	
	return []int{leftFn(nums, target), rightFn(nums, target)}
}
```

- [sqrtx](https://leetcode.cn/problems/sqrtx)

```go
package main
// 计算非负整数x的算数平方根
func mySqrt(x int) int {
	left, right := 0,x
	for left <= right {
		mid := left + (right-left) >> 1
		temp := mid * mid
		if temp == x {
			return mid
		} else if temp < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
    // 返回right, right是最靠近x的算术平方根的
	return right
}
```

- [valid-perfect-square](https://leetcode.cn/problems/valid-perfect-square)

```go
// 判断num是否是完全平方数
func isPerfectSquare(num int) bool {
	left, right := 0,num
	for left <= right {
		mid := left + (right-left) >> 1
		temp := mid *mid 
		if temp == num {
			return true
		} else if temp > num {
			right = mid -1
		} else {
			left = mid + 1
		}
	}

	return false
}
```



