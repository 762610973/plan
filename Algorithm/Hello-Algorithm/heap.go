package Hello_Algorithm

import (
	"container/heap"
	"fmt"
)

// Go 语言中可以通过实现 heap.Interface 来构建整数大顶堆
// 实现 heap.Interface 需要同时实现 sort.Interface
type intHeap []int

// Push heap.Interface 的方法，实现推入元素到堆
func (h *intHeap) Push(x any) {
	// Push 和 Pop 使用 pointer receiver 作为参数
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.(int))
}

// Pop heap.Interface 的方法，实现弹出堆顶元素
func (h *intHeap) Pop() any {
	// 待出堆元素存放在最后
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

// Len sort.Interface 的方法
func (h *intHeap) Len() int {
	return len(*h)
}

// Less sort.Interface 的方法
func (h *intHeap) Less(i, j int) bool {
	// 如果实现小顶堆，则需要调整为小于号
	return (*h)[i] > (*h)[j]
}

// Swap sort.Interface 的方法
func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Top 获取堆顶元素
func (h *intHeap) Top() any {
	return (*h)[0]
}

/* 获取左子节点索引 */
func (h *intHeap) left(i int) int {
	return 2*i + 1
}

/* 获取右子节点索引 */
func (h *intHeap) right(i int) int {
	return 2*i + 2
}

/* 获取父节点索引 */
func (h *intHeap) parent(i int) int {
	// 向下整除
	return (i - 1) / 2
}

// 访问堆顶元素
func (h *intHeap) peek() any {
	return (*h)[0]
}

/* 元素入堆 */
func (h *intHeap) push(val any) {
	// 添加节点
	*h = append(*h, val.(int))
	// 从底至顶堆化
	h.siftUp(len(*h) - 1)
}

/* 从节点 i 开始，从底至顶堆化 */
func (h *intHeap) siftUp(i int) {
	for {
		// 获取节点 i 的父节点
		p := h.parent(i)
		// 当“越过根节点”或“节点无须修复”时，结束堆化
		if p < 0 || (*h)[i] <= (*h)[p] {
			break
		}
		// 交换两节点
		h.Swap(i, p)
		// 循环向上堆化
		i = p
	}
}

/* 元素出堆 */
func (h *intHeap) pop() any {
	// 判空处理
	if (*h).Len() == 0 {
		fmt.Println("error")
		return nil
	}
	// 交换根节点与最右叶节点（即交换首元素与尾元素）
	h.Swap(0, h.Len()-1)
	// 删除节点
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	// 从顶至底堆化
	h.siftDown(0)

	// 返回堆顶元素
	return val
}

/* 从节点 i 开始，从顶至底堆化 */
func (h *intHeap) siftDown(i int) {
	for {
		// 判断节点 i, l, r 中值最大的节点，记为 m
		l, r, m := h.left(i), h.right(i), i
		if l < h.Len() && (*h)[l] > (*h)[m] {
			m = l
		}
		if r < h.Len() && (*h)[r] > (*h)[m] {
			m = r
		}
		// 若节点 i 最大或索引 l, r 越界，则无须继续堆化，跳出
		if m == i {
			break
		}
		// 交换两节点
		h.Swap(i, m)
		// 循环向下堆化
		i = m
	}
}

func testHeap() {
	/* 初始化堆 */
	// 初始化大顶堆
	maxHeap := &intHeap{}
	heap.Init(maxHeap)
	/* 元素入堆 */
	// 调用 heap.Interface 的方法，来添加元素
	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)
	heap.Push(maxHeap, 4)
	heap.Push(maxHeap, 5)

	/* 获取堆顶元素 */
	top := maxHeap.Top()
	fmt.Printf("堆顶元素为 %d\n", top)

	/* 堆顶元素出堆 */
	// 调用 heap.Interface 的方法，来移除元素
	heap.Pop(maxHeap) // 5
	heap.Pop(maxHeap) // 4
	heap.Pop(maxHeap) // 3
	heap.Pop(maxHeap) // 2
	heap.Pop(maxHeap) // 1

	/* 获取堆大小 */
	size := len(*maxHeap)
	fmt.Printf("堆元素数量为 %d\n", size)

	/* 判断堆是否为空 */
	isEmpty := len(*maxHeap) == 0
	fmt.Printf("堆是否为空 %t\n", isEmpty)
}
