# `sync.Map`
- 读操作不会加锁. 写操作先将当前map复制一份(dirty map), 然后在dirty map中进行修改, 最后替换会原来的map
- readOnly用于保存在读取时的快照, 保证读取操作的一致性.
- 写入dirty map时, 发现一些键不在m中, 将amended设置为true
## 结构
```go
type Map struct {
	mu Mutex
	read atomic.Pointer[readOnly]
	dirty map[any]*entry
	misses int
}
// readOnly用于提供并发安全的只读访问模式.
type readOnly struct {
	m       map[any]*entry
	amended bool // 如果dirty map中包含一些不在m中的值, 则为true, 否则为false.
}
// any类型的value
type entry struct {
	p atomic.Pointer[any]
}

// 用于标记已从dirty map中删除的item
var expunged = new(any)
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
		read = m.loadReadOnly()
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
		// 如果key存在, 尝试交换一下(更新或插入)
		if v, ok := e.trySwap(&value); ok {
			if v == nil {
				return nil, false
			}
			return *v, true
		}
	}
	m.mu.Lock()
	read = m.loadReadOnly()
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			// The entry was previously expunged, which implies that there is a
			// non-nil dirty map and this entry is not in it.
			m.dirty[key] = e
		}
		if v := e.swapLocked(&value); v != nil {
			loaded = true
			previous = *v
		}
	} else if e, ok := m.dirty[key]; ok {
		if v := e.swapLocked(&value); v != nil {
			loaded = true
			previous = *v
		}
	} else {
		if !read.amended {
			// We're adding the first new key to the dirty map.
			// Make sure it is allocated and mark the read-only map as incomplete.
			m.dirtyLocked()
			m.read.Store(&readOnly{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry(value)
	}
	m.mu.Unlock()
	return previous, loaded
}
```