#[allow(unused)]
mod string {
    /*
    * &str: 字符串切片类型
    str使用utf-8作为内部默认编码格式
    * &str是对一块字符串区间的借用,对指向的内存空间没有所有权
    ? String类型在堆上动态申请了一块内存空间,有权对这块内存空间扩容
    ? 实现了Deref<Target=str>的trait。所以在很多情况下,&String类型可以被编译器自动转换为&str类型
    */
    #[test]
    fn t1() {
        let greeting: &str = "hello";
        let greeting1 = "hello";
        let mut s = String::from("hello");
        s.make_ascii_uppercase();
        println!("{:?}", s);
        println!("{:?}", greeting.chars().nth(1));
    }
}