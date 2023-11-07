# Go调度的本质是一个生产-消费流程
- 生产数的goroutine放到可运行队列中
  - runnext: 只能只想一个goroutine, 是一个特殊的队列.
  - local queue: 大小为256的数组, 实际上用head和tail指针把它当成一个环形数组使用.
  - global queue
- 如果runnext为空, goroutine顺利放入runnext, 以最高优先级得到运行, 优先被消费.
- Go程序启动创建P, 创建初始的m0, m0启动一个调度循环, 不断地找g, 运行, 再找g
- 随程序运行, m更多地被创建出来, 生产者不断地生产g, m的调度循环不断地消费g
# 迷惑的goroutine执行顺序
```go
func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
// 顺序是9,0,1,2,3,4,5,6,7,8
```
- 本地只有一个p,for循环中生产出的goroutine都会进入到p的runnext和local queue
  - i从1开始, runnext已经有goroutine在了, 这是会把old goroutine移到p的本地队列中, 再把new goroutine放到runnext中, 重复这个过程.
  - 最后i为9, 新goroutine被放到runnext, 其余goroutine都在本地队列.
- go1.13的time包生产一个timerproc的goroutine用于唤醒挂在timer上的时间未到期的goroutine.
- go1.14去掉了这个用于唤醒的goroutine, 取而代之在调度循环的各个地方, sysmon里都是唤醒timer的代码, timer唤醒更及时.