- [binary-search](https://leetcode.cn/problems/binary-search/)
```go
package main

func search(nums []int, target int) int {
    left,right := 0,len(nums)-1
    for left <= right {
        mid := left+(right-left)/2
        if target == nums[mid] {
            return mid
        } else if target > nums[mid] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
	
    return -1
}
```
- [search-insert-position](https://leetcode.cn/problems/search-insert-position/)
```go
package main
func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right-left) / 2
		if nums[mid] == target {
			return mid
        } else if nums[mid] < target {
			left = mid + 1
        } else {
			right = mid - 1
        }
    }
	
	return left
}
```