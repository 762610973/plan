# `sync.Map`
- 读操作不会加锁. 写操作先将当前map复制一份(dirty map), 然后在dirty map中进行修改, 最后替换会原来的map
- readOnly用于保存在读取时的快照, 保证读取操作的一致性.
- 写入dirty map时, 发现一些键不在m中, 将amended设置为true
## 优化点
1. 空间换时间, 通过冗余的两个数据结构(只读的read, 可写的dirty), 减少加锁对性能的影响. 对只读字段的操作不需要枷锁.
2. 优先从read读取, 更新, 删除, 因为对read字段的读取不需要锁
3. 动态调整, miss次数多了之后, 将dirty数据提升为read, 避免总是从dirty中加锁读取.
4. double-checking, 加锁之后还要再检查read字段, 确定真的不存在才操作dirty字段
5. 延迟删除, 删除一个键值总是打标机, 只有在提升dirty字段为read字段的时候才清理删除的数据(只迁移没有被标记为删除得kv).
## 结构
```go
type Map struct {
	mu Mutex
	read atomic.Pointer[readOnly]
	dirty map[any]*entry
	// miss次数多了之后, 将dirty数据提升为read, 避免总是从dirty中加锁读取.
	misses int
}
// readOnly用于提供并发安全的只读访问模式.
type readOnly struct {
	m       map[any]*entry
	amended bool // 如果dirty map中包含一些不在m中的值(比如新增一条数据)  , 则为true, 否则为false.
}
// any类型的value
type entry struct {
	p atomic.Pointer[any]
}

// 用于标记已从dirty map中删除的item, 如果一个item被删除, 标记为expunged
var expunged = new(any)
func (m *Map) dirtyLocked() {
	if m.dirty != nil {	// dirty已经存在, 不需要创建
		return
	}
	// 如果dirty为nil, 从read字段中复制一个出来dirty对象
	read := m.loadReadOnly()
	m.dirty = make(map[any]*entry, len(read.m))
	for k, e := range read.m {
		if !e.tryExpungeLocked() {	// 将unpunged的键值对复制到dirty中.
			m.dirty[k] = e
		}
	}
}
```
## 方法
```go
func newEntry(i any) *entry {
	e := &entry{}
	e.p.Store(&i)
	return e
}
// 以原子方式加载, get or create
func (m *Map) loadReadOnly() readOnly {
	if p := m.read.Load(); p != nil {
		return *p
	}
	return readOnly{}
}
func (e *entry) load() (value any, ok bool) {
	p := e.p.Load()
	if p == nil || p == expunged {
		return nil, false
	}
	return *p, true
}
func (e *entry) delete() (value any, ok bool) {
	for {
		p := e.p.Load()
		// 如果p是nil或者已经被删除, 返回false
		if p == nil || p == expunged {
			return nil, false
		}
		// cas操作, 将p设置为nil
		// 如果在并发情况下发生竞争交换失败, 重新加载并再次尝试.
		if e.p.CompareAndSwap(p, nil) {
			return *p, true
		}
	}
}
// trySwap
func (e *entry) trySwap(i *any) (*any, bool) {
	for {
		p := e.p.Load()
		// 如果被删除, 返回nil和false
		if p == expunged {
			return nil, false
		}
		// 将p设置为i
		if e.p.CompareAndSwap(p, i) {
			return p, true
		}
	}
}
func (m *Map) missLocked() {
	m.misses++
	// miss次数小于dirty长度, 直接返回
	if m.misses < len(m.dirty) {
		return
	}
	// 如果miss次数已经大于dirty长度了, 将dirty变为read
	m.read.Store(&readOnly{m: m.dirty})
	// 将dirty设置为nil
	m.dirty = nil
	// miss清零
	m.misses = 0
}
```
```go
// Load 从read中查, 从dirty map中查
func (m *Map) Load(key any) (value any, ok bool) {
	read := m.loadReadOnly()	// 获取只读的map
	e, ok := read.m[key]
	// 如果key不存在, 并且amended为true, 表示dirty map中有未被记录的值.
	if !ok && read.amended {
		m.mu.Lock()
		// 重新加载只读map防止在等待锁的过程中dirty map被提升为read map
		read = m.loadReadOnly()	// double check
		e, ok = read.m[key]
		// 尝试在dirty map中查找
		if !ok && read.amended {
			e, ok = m.dirty[key]
			// 无论是否从dirty map中找到, 都记录一次
			m.missLocked()
		}
		m.mu.Unlock()
	}
	// 不存在返回nil和false
	if !ok {
		return nil, false
	}
	return e.load()
}
// LoadAndDelete 加载然后删除, 如果不存在返回false
func (m *Map) LoadAndDelete(key any) (value any, loaded bool) {
	read := m.loadReadOnly()
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		read = m.loadReadOnly()
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			// 无论从dirty中是否查到, 都从dirty中删除
			delete(m.dirty, key)
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if ok {
		// 如果存在, 调用delete, 其实是将e(key对应的value entry)设置为nil
		return e.delete()
	}
	return nil, false
}
func (m *Map) Swap(key, value any) (previous any, loaded bool) {
	read := m.loadReadOnly()
	if e, ok := read.m[key]; ok {
		// 如果key存在, 说明是更新
		if v, ok := e.trySwap(&value); ok {
			if v == nil {
				return nil, false
			}
			return *v, true
		}
	}
	// read中不存在, 或者cas更新失败, 加锁访问dirty
	m.mu.Lock()
	read = m.loadReadOnly()
	if e, ok := read.m[key]; ok {	// 双检查, 看看read是否已经存在了
		if e.unexpungeLocked() {
			// 此项目已经被删除了, 通过将它的值设置nil, 标记为unexpunged
			m.dirty[key] = e
		}
		// 更新
		if v := e.swapLocked(&value); v != nil {
			loaded = true
			previous = *v
		}
		// read中没有, 看dirty中有没有, 有的话直接更新
	} else if e, ok := m.dirty[key]; ok {
		if v := e.swapLocked(&value); v != nil {
			loaded = true
			previous = *v
		}
		// read和dirty中都没有, 这是一个新的key, 要插入, 插入时对dirtymap操作
	} else {
		if !read.amended {
			// 需要创建dirty对象, 并且标记read的amended为true
			// 说明有元素它不包含而dirty包含
			m.dirtyLocked()
			// 对于一个新增的数据, 设置amended为true
			m.read.Store(&readOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)	// 将新值添加到dirty对象中
	}
	m.mu.Unlock()
	return previous, loaded
}
```