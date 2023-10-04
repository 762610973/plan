```go
type slice struct {
	array unsafe.Pointer
	len int
	cap int
}
```
- 源码位置: `src/runtime/slice.go`
```go
// 运行时创建切片
func makeslice(et *_type, len, cap int) unsafe.Pointer {
	// 根据cap结合每个元素的大小, 计算出消耗的总容量
	mem, overflow := math.MulUintptr(et.Size_, uintptr(cap))
	// 下面的两次次判断主要是为了得知是cap-panic还是len-panic
	// 计算容量溢出, 超过允许的最大值, len < 0, len > cap
	if overflow || mem > maxAlloc || len < 0 || len > cap {
		// 倘若容量超限, len取赋值或len超过cap, 直接panic
		mem, overflow := math.MulUintptr(et.Size_, uintptr(len))
		if overflow || mem > maxAlloc || len < 0 {
			panicmakeslicelen()
		}
		panicmakeslicecap()
	}
	// 使用mallocgc进行内存分配以及切片初始化
	return mallocgc(mem, et, true)
}
// 容量计算
// mallocgc, 对新切片进行内存初始化
// memmove, 将老切片的内容拷贝到新切片中
// 返回扩容后的新切片

func growslice(oldPtr unsafe.Pointer, newLen, oldCap, num int, et *_type) slice {
	// newLen: 数组长度+被append进来的元素数量
	if newLen < 0 {
		panic(errorString("growslice: len out of range"))
	}
	if et.Size_ == 0 {
		// 如果元素大小为0, 无需分配直接返回即可
		return slice{unsafe.Pointer(&zerobase), newLen, newLen}
	}
	newcap := oldCap
	doublecap := newcap + newcap
	if newLen > doublecap {
		// 新的长度大于原容量的两倍, 直接取新长度作为数组扩容后的容量
		newcap = newLen
	} else {
		const threshold = 256
		// 老容量小于256, 扩容新容量为原来的两倍
		if oldCap < threshold {
			newcap = doublecap
		} else {
			for 0 < newcap && newcap < newLen {
				// 在原容量基础上, 新容量 = 原容量扩容5/4+192
				// 平滑过渡
				newcap += (newcap + 3*threshold) / 4
			}
			// 数值越界溢出, 取预期的新容量cap封顶
			if newcap <= 0 {
				newcap = newLen
			}
		}
	}
	p = mallocgc()
	memmove()
	
	return slice{p,newLen,newcap}
}
```

- 切片在编译期间的生成的类型只会包含切片中的元素类型, 切片中元素的类型都是在编译期间确定的.
- 编译器创建切片: `cmd/compile/internal/types.NewSlice`
```go
// 编译时创建切片
func NewSlice(elem *Type) *Type {
	// 查找缓存中是否已经存在一个切片类型
	if t := elem.cache.slice; t != nil {
		if t.Elem() != elem {
			base.Fatalf("elem mismatch")
		}
		if elem.HasShape() != t.HasShape() {
			base.Fatalf("Incorrect HasShape flag for cached slice type")
		}
		return t
	}

	t := newType(TSLICE)
	// extra字段是一个只包含切片内元素类型的结构. 编译器确定类型之后, 会将类型存储在extra字段帮助程序在运行时动态获取
	t.extra = Slice{Elem: elem}
	elem.cache.slice = t
	if elem.HasShape() {
		t.SetHasShape(true)
	}
	return t
}
```
- 编译是拷贝和运行时拷贝都会调用memmove