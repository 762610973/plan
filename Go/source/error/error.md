# error

> error是一个接口,定义如下

 ```go
type error interface {
	Error() string
}
 ```

## `errors.go`

> new一个error

```go
package errors

// New 方法接受一个字符串类型的错误信息,返回一个实现error接口的结构体/类
func New(text string) error {
    return &errorString{text}
}
// errorsString类型实现了error接口
type errorString struct {
	s string
}
func (e *errorString) Error() string {
	return e.s
}
```

## `join.go`

> 拼接多个err

```go
func Join(errs ...error) error {
	n := 0
	// 遍历errs, 计算不为nil的err数量
	for _, err := range errs {
		if err != nil {
			n++
		}
	}
	if n == 0 {
		return nil
	}
	e := &joinError{
		errs: make([]error, 0, n),
	}
	// 依次append进去
	for _, err := range errs {
		if err != nil {
			e.errs = append(e.errs, err)
		}
	}
	return e
}
// joinError 包含一组error
type joinError struct {
	errs []error
}
// 实现error接口,
func (e *joinError) Error() string {
	var b []byte
	for i, err := range e.errs {
		//  从第二个元素开始,加入换行符分割
		if i > 0 {
			b = append(b, '\n')
		}
		b = append(b, err.Error()...)
	}
	return string(b)
}

//  返回这一组error
func (e *joinError) Unwrap() []error {
	return e.errs
}

```

## `wrap.go`

```go
package errors

// Unwrap 用于得到一个错误的底层错误,这个错误必须实现[Unwrap() error]方法
func Unwrap(err error) error {
	// 将err断言成一个接口类型
	/*
		type xxx interface {
		    Unwrap() error
		}
	*/
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return nil
	}
	//调用这个类型的Unwrap方法
	return u.Unwrap()
}

// Is 判断err树中是否有一个err和target相等,对其子级进行深度优先遍历,有一个匹配就返回
func Is(err, target error) bool {
	if target == nil {
		// target == nil, 判断err是否为nil
		return err == target
	}
	// 利用反射判断target类型是否具有可比性 
	isComparable := reflectlite.TypeOf(target).Comparable()
	// 递归调用Unwrap拆包,返回下一层的err去判断是否相等
	for {
		// 如果可以比较, 并且两个err相等,返回true
		if isComparable && err == target {
			return true
		}
		// 类型断言成一个接口
		if x, ok := err.(interface {
			Is(error) bool
		}); ok && x.Is(target) {
			return true
		}
		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return false
			}
			// 有些error的Unwrap返回值是[]error类型 
		case interface{ Unwrap() []error }:
			for _, err = range x.Unwrap() {
				if Is(err, target) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}

// As 在err树中查找与target匹配的第一个错误, 如果找到,则将target设置为该错误值并返回true. 否则返回false
func As(err error, target any) bool {
	if err == nil {
		return false
	}
	if target == nil {
		// 要求target不能为nil
		panic("errors: target cannot be nil")
	}
	val := reflectlite.ValueOf(target)
	typ := val.Type()
	if typ.Kind() != reflectlite.Ptr || val.IsNil() {
		// 要求target必须为指针类型,并且不能为空
		panic("errors: target must be a non-nil pointer")
	}
	targetType := typ.Elem()
	if targetType.Kind() != reflectlite.Interface && !targetType.Implements(errorType) {
		// 要求target是一个接口类型或者实现了errorType
		panic("errors: *target must be interface or implement error")
	}
	// 深度优先搜搜, 使用Unwrap, 递归调用As
	for {
		// 判断是否可以将err分配给targetType 
		if reflectlite.TypeOf(err).AssignableTo(targetType) {
			// 如果可以分配,Set值
			val.Elem().Set(reflectlite.ValueOf(err))
			return true
		}
		// 如果err实现了As(any) bool 接口, 调用err本身的As方法
		if x, ok := err.(interface{ As(any) bool }); ok && x.As(target) {
			return true
		}
		// 通过Unwrap()获取错误
		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return false
			}
		case interface{ Unwrap() []error }:
			for _, err := range x.Unwrap() {
				if As(err, target) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}

var errorType = reflectlite.TypeOf((*error)(nil)).Elem()
```
