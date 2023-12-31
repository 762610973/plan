# array

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
package main
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

- [remove-element](https://leetcode.cn/problems/remove-element/)
```go
package main
// 原地移除所有等于val的元素, 返回移除后数组长度
func removeElement(nums []int, val int) int {
    // slow记录不相等的元素
    slow, fast := 0,0
    for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast ++
    }

	return slow
}
```

- [remove-duplicates-from-sorted-array](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)
```go
package main
// 原地删除有序数组中的重复项, 使得每个元素出现一次, 保持相对顺序, 返回总和
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	slow, fast := 0,1
	for fast < length {
        // 不相等就把slow+1, 然后赋值
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	
	return slow+1
}
```

- [move-zeroes](https://leetcode.cn/problems/move-zeroes/)
```go
package main
// 将所有的0移动到数组末尾, 保持非零元素的相对顺序 
func moveZeroes(nums []int)  {
	length := len(nums)
	slow, fast := 0,0
	for fast < length {
        // fast寻找不为0的元素, 和slow交换
		if nums[fast] != 0 {
			nums[fast],nums[slow] = nums[slow],nums[fast]
			slow++
		}
		fast++
	}
}
```

- [backspace-string-compare](https://leetcode.cn/problems/backspace-string-compare/)
```go
package main
func backspaceCompare(s string, t string) bool {
	return help1(s) == help1(t)
   	// return help2(s) == help2(t)
}

func help1(str string) string {
	b := []byte(str)
	var res []byte
	for i := 0;i < len(b); i++ {
		if b[i] != '#' {
			res = append(res, b[i])
		} else {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		}
	}
	return string(res)
}
func help2(str string) string {
	b := []byte(str)
	for i := 0;i < len(b); i++ {
		if b[i] == '#' {
			if i != 0 {
				b = append(b[:i-1], b[i+1:]...)
				i-=2
			} else {
				b = b[1:]
				i--
			}
		}
	}
	return string(b)
}
```

- [squares-of-a-sorted-array](https://leetcode.cn/problems/squares-of-a-sorted-array/)
```go
package main
func sortedSquares(nums []int) []int {
	length := len(nums)
	res := make([]int, length, length)
	index := length-1
	left, right := 0,length-1
    // 这里一定是<=
	for left <= right {
		l := nums[left]*nums[left]
		r := nums[right]*nums[right]
		if l > r {
			res[index] = l
			left++
      // 不能用else if处理相等情况, 因为left<=right, 会多一个, 所以直接else即可, 此时l==r, 也是多处理一次
		} else {
			res[index] = r
			right--
		} 
		index--
	}

	return res
}
```

- [minimum-size-subarray-sum](https://leetcode.cn/problems/minimum-size-subarray-sum/)
```go
package main
func minSubArrayLen(target int, nums []int) int {
	res := 1 << 31
	sum := 0
	left := 0
	for idx, val := range nums {
		sum += val
		for sum >= target {
			res = min(res, idx-left+1)
			sum -= nums[left]
			left++
		}
	}
	// 如果res没有改变, 说明数组的总和小于target
	if res == 1 << 31 {
		return 0
	}

	return res
}
```

- 最小的因数集合
```go
package main
// 求最小的因数集合
func main() {
	var fn func(int)
	m := make(map[int]struct{})
	fn = func(num int) {
		// 从2到num遍历,用num整除i
		for i := 2; i < num; i++ {
			// 能整除, 说明是
			if num%i == 0 {
				if !is(i) {
					m[i] = struct{}{}
				}
			}
		}
	}
	fn(100)
}
func is(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return true
		}
	}
	return false
}
```

- [minimum-window-substring](https://leetcode.cn/problems/minimum-window-substring/)
```go
package main
// s覆盖t的最小子串
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	tMap := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	resL, resR := -1, -1
	left := 0
	minLen := len(s)+1
	sMap := make(map[byte]int, len(s))
	contain := func() bool {
		for k,v := range tMap {
			if sMap[k] < v {
				return false
			}
		}

		return true
	}

	for i := 0; i < len(s); i++ {
		if tMap[s[i]] > 0 {
			sMap[s[i]] ++
		}
		for contain() && i-left + 1 >= len(t){
			if i-left+1 < minLen {
				minLen = i-left+1
				resL, resR = left, i+1
			}
			sMap[s[left]]--
			left++
		}
	}
	if resL == -1 {
		return ""
	}
	
	return s[resL:resR]
}
func minWindow(s string, t string) string {
	tMap := make(map[byte]int)
	sMap := make(map[byte]int)
	// 统计t
	for i := range t {
		tMap[t[i]]++
	}
	equalNum := 0
	start, resLen := 0, math.MaxInt32
	left, right := 0,0
	for right < len(s) {
		nowVal := s[right]
		right++
		// t中存在
		if tMap[nowVal] > 0 {
			sMap[nowVal]++
			if sMap[nowVal] == tMap[nowVal] {
				equalNum++
			}
		}
		for equalNum == len(tMap) {
			if right - left < resLen {
				resLen = right - left
				start = left
			}
			d := s[left]
			left++
			// 存在最左边那个
			if tMap[d] > 0 {
				if sMap[d] == tMap[d] {
					equalNum--
				}
				sMap[d]--
			}
		}
	}
	if resLen == math.MaxInt32 {
		return ""
	}
	return s[start:start+resLen]
}
```

- [spiral-matrix-ii](https://leetcode.cn/problems/spiral-matrix-ii/)
```go
package main
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int,n)
	}
	total := n * n
	left, right := 0, n-1
	top, low := 0, n-1
	index := 1
	for index <= total {
		for i := left; i <= right; i++ {
			res[top][i] = index
			index++
		}
		top++
		for i := top; i <= low; i++ {
			res[i][right] = index
			index++
		}
		right--
		for i := right; i >= left; i-- {
			res[low][i] = index
			index++
		}
		low--
		for i := low; i >= top; i-- {
			res[i][left] = index
			index++
		}
		left++
	}

	return res
}
```