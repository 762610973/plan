fn main() {
    println!("Hello, world!");
}


#[allow(unused)]
fn test1(t: (i32, i32)) {
    println!("{:?}", t.0 + t.1);
}

#[allow(unused)]
fn test2((x, y): (i32, i32)) -> i32 {
    x + y
}

#[allow(unused)]
fn test3(t: (i32, i32)) -> i32 {
    t.0 + t.1
}

#[test]
#[allow(unused)]
// * 函数可以被当成头等公民,被复制到一个值中
fn test4() {
    let p = (1, 2);
    let func = test2;
    println!("{}", func(p));
    let func1 = test1(p);   //?此时func1是一个unit类型,因为调用了test
    println!("{:?}", func1);
}


// ? Rust的函数体内也允许定义其他item,包括静态变量,常量,func,trait,type,mod,可以避免污染外部的命名空间
#[allow(unused)]
mod change {
    use crate::{test2, test3};

    // * 通过as进行类型转换
    #[test]
    fn test5_change() {
        let p = (1, 2);
        // func = test3;    参数一样,返回值一样,转换失败
        let mut func = test2 as fn((i32, i32)) -> i32;
        func = test3;
        println!("{:?}", func(p));
    }

    // * 初始化时指定类型
    #[test]
    fn test6_change() {
        let p = (1, 2);
        let mut func: fn((i32, i32)) -> i32 = test2;
        // func = test3;    参数一样,返回值一样,转换失败
        func = test3;
        println!("{:?}", func(p));
    }
}

// ?发散函数
#[allow(unused)]
mod diverging_func {
    // *发散函数
    #[test]
    fn test1() -> ! {
        panic!("this function never returns!");
    }

    fn test2() {
        let x = true;
        let p = if x {
            panic!("error");    // ? 返回!类型,与任意类型相容
        } else {
            100
        };
    }
    // ? 发散函数
    // * panic!以及基于它实现的各种函数/宏
    // * 死循环loop{}
    // * 进程退出函数std;
}

mod const_fn {
    // *函数可以用const关键字修饰，这样的函数可以在编译阶段被编译器执行
    // * 返回值也被视为编译期常量
    const fn cube(num: usize) -> usize {
        num.pow(3)
    }

    #[test]
    fn main() {
        const DIM: usize = cube(2);
        const ARR: [i32; DIM] = [0; DIM];
        let p = [1, 2, 3, 4];
        let p2 = [5; 5];
        println!("{:?}", ARR);
        println!("{:?}", p);
        println!("{:?}", p2);
    }
}