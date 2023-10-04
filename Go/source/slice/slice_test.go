package slice

import (
	"fmt"
	"testing"
)

func Test_slice1(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[9:]
	s[9] = 100
	s = nil
	fmt.Println(s1, len(s1), cap(s1))
}

// 容量翻倍
func Test_slice2(t *testing.T) {
	s := make([]int, 10)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

// 由于初始化不同, 容量没有翻倍
func Test_slice3(t *testing.T) {
	s := make([]int, 0, 10)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

// 容量没有变化, 因为预先分配了多余的容量
func Test_slice4(t *testing.T) {
	s := make([]int, 10, 11)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

// 容量为4
func Test_slice5(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:9]
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

// 截取操作以s[8]作为内存空间的起点,  截取所得新切片s1的长度和容量依赖于原切片s的长度和容量
// 并在此基础上减去头部8个未使用的单位
func Test_slice6(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	// 长度为2, 容量为4
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

// 数组访问越界, s1发生扩容, 对原来的数组没有影响
func Test_slice7(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1 = append(s1, []int{10, 11, 12}...)
	_ = s[10]
	// ...
}

// changeSlice不会影响原来切片的长度和容量
func Test_slice8(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice(s1)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func changeSlice(s1 []int) {
	s1 = append(s1, 10)
}

// 越界发生panic
func Test_slice9(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}
	s = append(s[:2], s[3:]...)
	// 此时s的长度已经发生改变
	t.Logf("s: %v, len: %d, cap: %d", s, len(s), cap(s))
	_ = s[4]
}

func Test_slice10(t *testing.T) {
	s := []int{0, 1, 2}
	s = append(s, 3, 4, 5, 6)
	// 长度为3+4=7
	// 容量: newLen(7) > 6(oldLen+oldLen), cap为newLen(7), 但是要考虑内存分配问题, 所以真实容量为8
	// runtime.roundupsize函数会将待申请的内存向上取整, 取整时会使用runtime.class_to_size数组, 可以提高内存的分配效率并较少碎片
	t.Logf("s: %v, len: %d, cap: %d", s, len(s), cap(s))
}
