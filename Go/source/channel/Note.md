1. 对一个关闭的通道再发送值就会导致panic
2. 对一个关闭的通道进行接收会一直获取值直到通道为空
3. 对一个关闭的并且没有值的通道执行接收操作会得到该对应类型的零值
4. 关闭一个已经关闭的通道会导致panic
5. 向一个已经满了的channel发送数据, 会死锁
6. 尽量在channel的发送方关闭channel