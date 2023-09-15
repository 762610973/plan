fn main() {}
#[test]
#[allow(unused)]
fn test1() {
    let mut s = String::from("ownership");
    s.push_str("ownership");
    println!("s: {}", s);
    // 自动析构
}

/*#[test]
fn test2() {
    let s = String::from("hello");
    let s1 = s; // 变量s1的生命周期从生命开始到move给s1已经结束了, 字符串本身并不会销毁再创建
    println!("{:?}", s);
    // 由s拥有的字符串已经转移给了s1这个变量, 后面继续使用s是不对的
}*/

#[allow(unused)]
mod remove {
    fn create() -> String {
        let s = String::from("hello");
        return s; // 所有权转移,从函数内部移动到外部, 同时生命周期结束
    }
    fn consume(s: String) {
        // 所有权转移,从函数外部移动到内部
        println!("{}", s);
    }
    fn main() {
        let s = create();
        consume(s);
    }
}

#[allow(unused)]
mod copy {
    // !凡是实现了std::marker::Copy trait的类型, 都会执行copy语义
    struct Foo {
        data: i32,
    }

    impl Clone for Foo {
        fn clone(&self) -> Self {
            Foo { data: self.data }
        }
    }

    impl Copy for Foo {}

    #[test]
    fn test() {
        let v1 = Foo { data: 0 };
        let v2 = v1;
        println!("{}", v1.data);
    }
    #[derive(Copy, Clone)]
    struct Today {
        data: i32,
    }
    //只要一个类型的所有成员都具有Clone trait，我们就可以使用 #[derive(Copy, Clone)]来让编译器帮我们实现Clone trait
}
/*
!ownership
* 每个值在一个时间点上只有一个管理者
* 当变量所在的作用域结束的时候, 变量以及它代表的值将会被销毁
* 变量从出生到死亡的整个阶段, 叫做变量的"生命周期"
*/

/*
* 一个变量可以把它拥有的值转移给另外一个变量, 称为"所有权转移"
* 赋值语句、函数调用、函数返回等, 都有可能导致所有权转移
* Rust中所有权转移的重要特点是, 它是所有类型的默认语义
* Rust中的变量绑定操作, 默认是move语义
* 语义不代表最终的执行效率, 只规定了什么样的代码是编译器可以接受的, 以及它执行后的效果可以用怎样的思维模型去理解
* 编译器有权在不改变语义的情况下做任何有利于执行效率的优化
*/

mod box_type {
    // !Box类型是一种指针类型, 代表"拥有所有权的指针"
    // !Box类型永远执行的是move语义, 不能是copy语义
    // !对于Rust里面的所有变量，在使用前一定要合理初始化
    struct T {
        value: i32,
    }
    #[test]
    fn test() {
        let p: Box<T> = Box::new(T { value: 3 });
        println!("{}", p.value);
    }
}

mod copy_clone {
    /*
     * std::marker::Copy, 没有任何方法, 跟编译器紧密绑定, 只是为了给类型打标
     * 实现了Copy trait, 任何时候可以通过简单的内存复制实现该类型的复制
     * 对于自定义类型, 只有所有的成员都实现了Copy trait, 这个类型才有资格实现Copy trait
     */

    /*
     * std::clone::clone, 深拷贝, 但不强制
     */
    #[derive(Copy, Clone)]
    struct MyStruct(i32);
}

mod destructor_constructor {
    // 析构
    // 构造
    // RAII: 利用生命周期绑定资源的使用周期, 生命周期开始时申请资源, 结束时释放资源
    // Rust中的析构: std::ops::Drop
    // 先构造的后析构, 后构造的先析构, 先进后出
    // 用户手动调用析构函数是非法的
    // std::mem::drop 手动调用标准库函数, 提前终止生命周期
    // 如果你用下划线来绑定一个变量, 那么这个变量会当场执行析构, 而不是等到当前语句块结束的时候再执行
    // 带有析构函数的类型都是不能满足Copy语义的
    // 析构函数是在变量生命周期结束的时候被调用的, 调用析构函数是编译器自动插入的代码做的
}
