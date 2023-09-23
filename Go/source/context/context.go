package main

import (
	"errors"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

// Context context的方法可以被多个goroutine同时调用
type Context interface {
	// Deadline 返回context的过期时间
	Deadline() (deadline time.Time, ok bool)
	// Done 返回context中的channel
	Done() <-chan struct{}
	// Err 返回错误
	Err() error
	// Value 返回context中key对应的value
	Value(key any) any
}

// Canceled context被cancel时,ctx.Err()会返回此信息
var Canceled = errors.New("context canceled")

// DeadlineExceeded context超时会报此错误
var DeadlineExceeded error = deadlineExceededError{}

type deadlineExceededError struct{}

func (deadlineExceededError) Error() string   { return "context deadline exceeded" }
func (deadlineExceededError) Timeout() bool   { return true }
func (deadlineExceededError) Temporary() bool { return true }

type emptyCtx struct{}

// Deadline 返回默认值,
func (emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

// Done nil channel,无法读到信号,永远阻塞
func (emptyCtx) Done() <-chan struct{} {
	return nil
}

// Err 空context没有错误, 此时调用err.Error()会panic
func (emptyCtx) Err() error {
	return nil
}

// Value 空context没有值
func (emptyCtx) Value(key any) any {
	return nil
}

type backgroundCtx struct {
	emptyCtx
}

type todoCtx struct {
	emptyCtx
}

// 返回这个context的类型: context.Background
func (backgroundCtx) String() string {
	return "context.Background"
}

// 返回这个context的类型: context.TODO
func (todoCtx) String() string {
	return "context.TODO"
}

// Background 返回一个非零的, 空的Context. 它从未被取消, 没有值, 也没有截止日期.
// 通常由主函数, 初始化和测试使用, 并作为传入请求的顶级上下文
func Background() Context {
	return backgroundCtx{}
}

func TODO() Context {
	return todoCtx{}
}

// CancelFunc 调用一个操作去停止, 可以被多个goroutine同时调用,第一次调用之后,后面的操作没有不会执行操作
type CancelFunc func()

// WithCancel 调用cancel()来触发信号, 然后接收到Done()信号
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := withCancel(parent)
	return c, func() {
		// 需要取消context时, 调用此方法
		c.cancel(true, Canceled, nil)
	}
}

// A CancelCauseFunc behaves like a [CancelFunc] but additionally sets the cancellation cause.
// This cause can be retrieved by calling [Cause] on the canceled Context or on
// any of its derived Contexts.
//
// If the context has already been canceled, CancelCauseFunc does not set the cause.
// For example, if childContext is derived from parentContext:
//   - if parentContext is canceled with cause1 before childContext is canceled with cause2,
//     then Cause(parentContext) == Cause(childContext) == cause1
//   - if childContext is canceled with cause2 before parentContext is canceled with cause1,
//     then Cause(parentContext) == cause1 and Cause(childContext) == cause2
type CancelCauseFunc func(cause error)

// WithCancelCause 对比WithCancel, 增加了一个cause
func WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc) {
	c := withCancel(parent)
	return c, func(cause error) { c.cancel(true, Canceled, cause) }
}

func withCancel(parent Context) *cancelCtx {
	// 校验父Context是否为空
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	c := &cancelCtx{}
	// 传播取消
	c.propagateCancel(parent, c)
	return c
}

// Cause 返回context被取消的原因, 如果还没有被取消, 返回nil
func Cause(c Context) error {
	if cc, ok := c.Value(&cancelCtxKey).(*cancelCtx); ok {
		cc.mu.Lock()
		defer cc.mu.Unlock()
		return cc.cause
	}

	return nil
}

// AfterFunc arranges to call f in its own goroutine after ctx is done
// (cancelled or timed out).
// If ctx is already done, AfterFunc calls f immediately in its own goroutine.
//
// Multiple calls to AfterFunc on a context operate independently;
// one does not replace another.
//
// Calling the returned stop function stops the association of ctx with f.
// It returns true if the call stopped f from being run.
// If stop returns false,
// either the context is done and f has been started in its own goroutine;
// or f was already stopped.
// The stop function does not wait for f to complete before returning.
// If the caller needs to know whether f is completed,
// it must coordinate with f explicitly.
//
// If ctx has a "AfterFunc(func()) func() bool" method,
// AfterFunc will use it to schedule the call.
func AfterFunc(ctx Context, f func()) (stop func() bool) {
	a := &afterFuncCtx{
		f: f,
	}
	a.cancelCtx.propagateCancel(ctx, a)
	return func() bool {
		stopped := false
		a.once.Do(func() {
			stopped = true
		})
		if stopped {
			a.cancel(true, Canceled, nil)
		}
		return stopped
	}
}

type afterFuncer interface {
	AfterFunc(func()) func() bool
}

type afterFuncCtx struct {
	cancelCtx
	once sync.Once // either starts running f or stops f from running
	f    func()
}

func (a *afterFuncCtx) cancel(removeFromParent bool, err, cause error) {
	a.cancelCtx.cancel(false, err, cause)
	if removeFromParent {
		removeChild(a.Context, a)
	}
	a.once.Do(func() {
		go a.f()
	})
}

// A stopCtx is used as the parent context of a cancelCtx when
// an AfterFunc has been registered with the parent.
// It holds the stop function used to unregister the AfterFunc.
type stopCtx struct {
	Context
	stop func() bool
}

var goroutines atomic.Int32

// &cancelCtxKey 用来判定一个context是cancelCtx
var cancelCtxKey int

// parentCancelCtx 从父context中查找是否存在cancelCtx类型
func parentCancelCtx(parent Context) (*cancelCtx, bool) {
	done := parent.Done()
	// 已经取消, 或者为nil, 无法从父context中获取到*cancelCtx
	if done == closedChan || done == nil {
		return nil, false
	}
	// 判断父context是否是cancelCtx类型
	p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)
	if !ok {
		return nil, false
	}
	// 如果父context是cancelCtx类型, 取出来
	pdone, _ := p.done.Load().(chan struct{})
	// 父context的done和刚刚取出来的done不相等, 说明父context的done已经发生了变化
	if pdone != done {
		return nil, false
	}

	return p, true
}

// removeChild 移除context从它的父context中.
func removeChild(parent Context, child canceler) {
	if s, ok := parent.(stopCtx); ok {
		s.stop()
		return
	}
	p, ok := parentCancelCtx(parent)
	if !ok {
		return
	}
	p.mu.Lock()
	if p.children != nil {
		delete(p.children, child)
	}
	p.mu.Unlock()
}

// canceler 监听到done信号, *cancelCtx, *timerCtx, *afterFuncCtx
type canceler interface {
	cancel(removeFromParent bool, err, cause error)
	Done() <-chan struct{}
}

// closedChan context取消时, 使用这个值, 然后其他地方就可以用这个值判断context是否取消了
var closedChan = make(chan struct{})

func init() {
	close(closedChan)
}

// cancelCtx 可以被取消,还会取消所有实现了canceler的子项
type cancelCtx struct {
	// 父context,有且只有一个
	Context
	// 并发保护
	mu   sync.Mutex
	done atomic.Value // of chan struct{}, created lazily, closed by first cancel call
	// set类型.
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
	cause    error                 // set to non-nil by the first cancel call
}

func (c *cancelCtx) Value(key any) any {
	// cancelCtxKey是全局的一个不可导出变量
	// 会有一个地方对cancelCtxKey做操作(非用户,只能是此包里面的一个地方)
	// - 倘若key是特定值,返回Context本身
	if key == &cancelCtxKey {
		return c
	}
	return value(c.Context, key)
}

// Done 懒汉模式,原子读写+互斥锁, 双检测, 确保done只有一个
func (c *cancelCtx) Done() <-chan struct{} {
	d := c.done.Load()
	// done有值, 直接返回, 避免不要的锁操作
	if d != nil {
		return d.(chan struct{})
	}
	// 如果d == nil, 加锁保护创建一个(这里保护的是c.done)
	c.mu.Lock()
	defer c.mu.Unlock()
	d = c.done.Load()
	// 再次检测, 原因是加锁之前可能有别的goroutine已经创建完了
	if d == nil {
		d = make(chan struct{})
		c.done.Store(d)
	}

	return d.(chan struct{})
}

func (c *cancelCtx) Err() error {
	c.mu.Lock()
	err := c.err
	c.mu.Unlock()
	return err
}

func (c *cancelCtx) propagateCancel(parent Context, child canceler) {
	// 将父context保存到自己的结构中
	c.Context = parent
	done := parent.Done()
	if done == nil {
		// done==nil,无法从监听ctx.Done()
		// 父context永远不会取消,没有必有传播取消这个特性了
		return
	}
	select {
	case <-done:
		// done不为nil之后, 监听父context的done信号
		// 可以从父context的done中接收到值,说明父context已经取消
		// 此时,child也要取消, 由于这个child就是当前这一层的, 所以就是取消自身, 不需要从父context中移除自身, 因为父context已经取消
		child.cancel(false, parent.Err(), Cause(parent))
		return
	default:
	}
	// 父context还没有取消
	// 如果父context也是一个cancelCtx类型
	if p, ok := parentCancelCtx(parent); ok {
		// parent is a *cancelCtx, or derives from one.
		p.mu.Lock()
		if p.err != nil {
			// 父context已经取消
			child.cancel(false, p.err, p.cause)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			// 父context没有取消, 添加到父context的set中
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
		return
	}

	if a, ok := parent.(afterFuncer); ok {
		// parent implements an AfterFunc method.
		c.mu.Lock()
		stop := a.AfterFunc(func() {
			child.cancel(false, parent.Err(), Cause(parent))
		})
		c.Context = stopCtx{
			Context: parent,
			stop:    stop,
		}
		c.mu.Unlock()
		return
	}

	goroutines.Add(1)
	go func() {
		select {
		// 父context终止,取消子context
		case <-parent.Done():
			child.cancel(false, parent.Err(), Cause(parent))
			// 子context取消,不影响父context,因为传播是单向性的
		case <-child.Done():
		}
	}()
}

type stringer interface {
	String() string
}

// contextName 返回context的名字
func contextName(c Context) string {
	if s, ok := c.(stringer); ok {
		return s.String()
	}
	return reflect.TypeOf(c).String()
}

func (c *cancelCtx) String() string {
	return contextName(c.Context) + ".WithCancel"
}

// cancel 关闭c.done, 取消每一个子context.
func (c *cancelCtx) cancel(removeFromParent bool, err, cause error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
	if cause == nil {
		cause = err
	}
	c.mu.Lock()
	// 检查自身是否已经被取消(ctx可能被并发访问), 所以cancel可以被多次调用
	if c.err != nil {
		c.mu.Unlock()
		return // 已经取消了
	}
	c.err = err
	c.cause = cause
	d, _ := c.done.Load().(chan struct{})
	if d == nil {
		// d==nil, 说明c.Done还没有被调用, 此时给done赋值
		c.done.Store(closedChan)
	} else {
		// 关闭通道, 此时<-c.Done()可以接收到值, 即使是零值, 但是就不会阻塞住了
		close(d)
	}
	for child := range c.children {
		// 传入false的目的是为了在取消子context时避免死锁, 因为这里的cancel是在父context中的cancel中进行的
		// 父context的锁已经被获取, 所以在取消子context时, 也已经持有了父context的锁
		// 传入false的目的是为了在取消子context时避免死锁情况的发生
		child.cancel(false, err, cause)
	}
	c.children = nil
	c.mu.Unlock()
	if removeFromParent {
		// 将自身从父context中移除
		removeChild(c.Context, c)
	}
}

func WithoutCancel(parent Context) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	return withoutCancelCtx{parent}
}

type withoutCancelCtx struct {
	c Context
}

func (withoutCancelCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (withoutCancelCtx) Done() <-chan struct{} {
	return nil
}

func (withoutCancelCtx) Err() error {
	return nil
}

func (c withoutCancelCtx) Value(key any) any {
	return value(c, key)
}

func (c withoutCancelCtx) String() string {
	return contextName(c.c) + ".WithoutCancel"
}

// WithDeadline returns a copy of the parent context with the deadline adjusted
// to be no later than d. If the parent's deadline is already earlier than d,
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
// [Context.Done] channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is
// closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this [Context] complete.
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	return WithDeadlineCause(parent, d, nil)
}

// WithDeadlineCause behaves like [WithDeadline] but also sets the cause of the
// returned Context when the deadline is exceeded. The returned [CancelFunc] does
// not set the cause.
func WithDeadlineCause(parent Context, d time.Time, cause error) (Context, CancelFunc) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if cur, ok := parent.Deadline(); ok && cur.Before(d) {
		// The current deadline is already sooner than the new one.
		return WithCancel(parent)
	}
	c := &timerCtx{
		deadline: d,
	}
	c.cancelCtx.propagateCancel(parent, c)
	dur := time.Until(d)
	if dur <= 0 {
		c.cancel(true, DeadlineExceeded, cause) // deadline has already passed
		return c, func() { c.cancel(false, Canceled, nil) }
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err == nil {
		c.timer = time.AfterFunc(dur, func() {
			c.cancel(true, DeadlineExceeded, cause)
		})
	}
	return c, func() { c.cancel(true, Canceled, nil) }
}

// A timerCtx carries a timer and a deadline. It embeds a cancelCtx to
// implement Done and Err. It implements cancel by stopping its timer then
// delegating to cancelCtx.cancel.
type timerCtx struct {
	// 基于cancelCtx
	cancelCtx
	// 用于在过期时刻中止context
	timer *time.Timer
	// context的过期时间
	deadline time.Time
}

func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, true
}

func (c *timerCtx) String() string {
	return contextName(c.cancelCtx.Context) + ".WithDeadline(" +
		c.deadline.String() + " [" +
		time.Until(c.deadline).String() + "])"
}

// 复用继承的 cancelCtx的cancel能力, 进行cancel处理
// 判断是否需要手动从parent的children set中移除
// 停止time.Timer
func (c *timerCtx) cancel(removeFromParent bool, err, cause error) {
	c.cancelCtx.cancel(false, err, cause)
	if removeFromParent {
		// Remove this timerCtx from its parent cancelCtx's children.
		removeChild(c.cancelCtx.Context, c)
	}
	c.mu.Lock()
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	c.mu.Unlock()
}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this [Context] complete:
//
//	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
//		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
//		defer cancel()  // releases resources if slowOperation completes before timeout elapses
//		return slowOperation(ctx)
//	}
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

// WithTimeoutCause behaves like [WithTimeout] but also sets the cause of the
// returned Context when the timeout expires. The returned [CancelFunc] does
// not set the cause.
func WithTimeoutCause(parent Context, timeout time.Duration, cause error) (Context, CancelFunc) {
	return WithDeadlineCause(parent, time.Now().Add(timeout), cause)
}

// WithValue only for request-scoped data that transits processes and APIs
// 每调用以此, 就生成一层
func WithValue(parent Context, key, val any) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	// key必须是可比较的非nil值
	if key == nil {
		panic("nil key")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return &valueCtx{parent, key, val}
}

// valueCtx 无须实现Context接口, 它组合/继承了Context
type valueCtx struct {
	Context
	key, val any
}

// stringify tries a bit to stringify v, without using fmt, since we don't
// want context depending on the unicode tables. This is only used by
// *valueCtx.String().
func stringify(v any) string {
	switch s := v.(type) {
	case stringer:
		return s.String()
	case string:
		return s
	}
	return "<not Stringer>"
}

func (c *valueCtx) String() string {
	return contextName(c.Context) + ".WithValue(type " +
		reflect.TypeOf(c.key).String() +
		", val " + stringify(c.val) + ")"
}

func (c *valueCtx) Value(key any) any {
	// 从第一层查找值
	if c.key == key {
		return c.val
	}
	// c.Context是当前Context的父Context
	return value(c.Context, key)
}

// value 查找key对应的value, 找到了就返回, 如果发现context是其他类型就返回context, 而不是value
func value(c Context, key any) any {
	// c = ctx.Context: 取出父context,向上寻找
	for {
		switch ctx := c.(type) {
		case *valueCtx:
			if key == ctx.key {
				return ctx.val
			}
			c = ctx.Context
		case *cancelCtx:
			if key == &cancelCtxKey {
				return c
			}
			c = ctx.Context
		case withoutCancelCtx:
			if key == &cancelCtxKey {
				// This implements Cause(ctx) == nil
				// when ctx is created using WithoutCancel.
				return nil
			}
			c = ctx.c
		case *timerCtx:
			if key == &cancelCtxKey {
				return &ctx.cancelCtx
			}
			c = ctx.Context
		case backgroundCtx, todoCtx:
			return nil
		default:
			// 如果输入的context.Context不是上述任何类型, 那么就调用其Value方法, 传入输入的key, 返回其结果
			// 用于自己实现Context接口的类型
			return c.Value(key)
		}
	}
}
