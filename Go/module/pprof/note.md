# `net/http/pprof`
- `allocs`: 查看过去所有内存分配的样本
- `block`: 查看导致阻塞同步的堆栈跟踪
- `cmdline`: 查看当前程序命令行的完整调用路径
- `goroutine`: 查看当前所有运行的goroutines堆栈跟踪
- `mutex`: 查看导致互斥锁的竞争持有者的堆栈跟踪,
- `profile`: 默认进行30s的CPU profiling, 会得到一个分析用的profile文件
- `threadcreate`:查看创建新OS线程的堆栈跟踪
# 通过交互式终端

- `flat`: 函数自身的运行耗时
- `flat%`: 函数自身占CPU运行总耗时的比例
- `sum%`: 函数自身累积使用占CPU运行总耗时比例
- `cum`: 函数自身及其调用函数的运行总耗时
- `cum%`: 函数自身及其调用函数占CPU运行总耗时的比例
- `Name`: 函数名

- `wget http://localhost:6060/debug/pprof/profile`
- `go tool pprof -http=:6001 profile`
- `go tool pprof profile(交互终端中输入web)`
- `go tool pprof -http=:6001 http://localhost:6060/debug/pprof/profile`