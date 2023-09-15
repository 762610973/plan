fn main() {
    const C1: &str = "一般来说. CPU有专门的指令可以用于入栈或者出栈的操作";
    const C2: &str = "堆是为动态分配预留的内存空间, 用户可以在任何时候分配和释放它";
    const C3: &str = "堆是在内存中动态分配的内存, 是无序的. 每个线程都有一个栈, 但是每一个应用程序通常都只有一个堆";
    const C4: &str = "有许多定制的堆分配策略用来为不同的使用模式下调整堆的性能";
    const C5: &str = "操作系统提供了在堆上分配和释放内存的系统调用, 但是用户不是直接使用这个系统调用，而是使用封装的更好的“内存分配器";
    const C6: &str = "·栈上保存的局部变量在退出当前作用域的时候会自动释放";
    const C7: &str = "栈有一个确定的最大长度, 超过了这个长度会产生“栈溢出”(stack overflow)";
    const C8: &str = "堆上分配的空间没有作用域，需要手动释放";
    const C9: &str =
        "堆的空间一般要更大一些, 堆上的内存耗尽了, 就会产生“内存分配不足”(out of memory)";
    const C10: &str = "core dump是程序失控之后, 触发了操作系统的保护机制而被动退出";
    println!("{:?}", C1);
    println!("{:?}", C2);
    println!("{:?}", C3);
    println!("{:?}", C4);
    println!("{:?}", C5);
    println!("{:?}", C6);
    println!("{:?}", C7);
    println!("{:?}", C8);
    println!("{:?}", C9);
    println!("{:?}", C10);
}

/*
代码段: 编译后的机器码存在的区域,一般只读的\
bss段: 存放未初始化的全局变量和静态变量的区域\
数据段: 存放有初始化的全局变量和静态变量的区域
函数调用栈: 存放函数参数, 局部变量以及其他函数调用相关信息的区域
堆: 存放动态分配内存的区域
*/
/*
segfault是这样形成的: 进程空间中的每个段通过硬件MMU映射到真正的物理空间;
在这个映射过程中, 我们还可以给不同的段设置不同的访问权限, 比如代码段就是只能读不能写；进程在执行过程中
如果违反了这些权限, CPU会直接产生一个硬件异常；硬件异常会被操作系统内核处理，一般内核会向对应的进程发送一条信号
如果没有实现自己特殊的信号处理函数, 默认情况下, 这个进程会直接非正常退出;
如果操作系统打开了core dump功能, 在进程退出的时候操作系统会把它当时的内存状态、寄存器状态以及各种相关信息保存到一个文件中，供用户以后调试使用
*/

/*
!内存不安全
*空指针: 解引用空指针
*野指针: 未被初始化, 它的值取决于这个位置以前遗留下来的是什么值
*悬空指针:悬空指针指的是内存空间在被释放了之后, 继续使用
*使用未初始化内存
*非法释放:内存分配和释放要配对. 如果对同一个指针释放两次, 会制造出内存错误
*缓冲区溢出:指针访问越界了, 结果也是类似于野指针, 会读取或者修改临近内存空间的值, 造成危险
*执行非法函数指针:如果一个函数指针不是准确地指向一个函数地址, 那么调用这个函数指针会导致一段随机数据被当成指令来执行, 是非常危险的
*数据竞争:并发场景下, 针对同一块内存同时读写, 且没有同步措施
*/
