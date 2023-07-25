use std::sync::atomic::AtomicBool;

fn main() {
    grow(1,1);
    // 语法歧义少,语法分析更容易
    // 方便引入类型推导功能
    // 方便 模式解构
    // let variable:i32 = 100;
    // 将mut a1看做一个组合
    // let mut a1 = 1;
    // let (mut a2,mut a3) = (2,3);
    //  不能使用未初始化的变量,编译器会做执行路径的静态分析
    // let unused :i32;
    // println!("{:?}", unused);
    // used binding `unused` isn't initialized
}

fn test(condition:bool) {
    let x:i32;
    if condition {
        x = 1;  // 初始化x不需要x是mut的,因为是初始化,不是修改
        println!("{:?}", 1);
    }
    // 如果条件不满足,x没有被初始化
    // 不使用x就可以
}
// 类型没有"默认构造函数",变量没有"默认值"

fn shadow() {
    let x = "hello, world";
    println!("{:?}", x);
    let x = 5;
    println!("{:?}", x);
}

fn shadow1() {
    let v = Vec::new();
    let mut v = v;
    v.push(1);
    println!("{:?}", v);
}

fn get_type() {
    let elem = 5u8;
    let mut vec = Vec::new();
    vec.push(elem);
    println!("{:?}",vec);
}

// 只允许局部变量/全局变量实现类型推导
fn get_type2() {
    let play = [
        ("jack",20),("jane",23)
    ];
    let players:Vec<_> = play.iter()
        .map(|&(player,_score)| {
        player
    })
    .collect();
    println!("{:?}",players);
}
// 类型别名,Go中的类型别名: type Age = uint32
type Age = u32;
fn grow(age:Age,year: u32) -> Age {
    if age == year {
        println!("{}", "true");
    } else {
        println!("{:?}", "false");

    }
    return age+year;
}
// 静态变量
// 全局变量的初始化必须是编译器可以确定的常量
// 带有mut修饰的全局变量,在使用的时候必须使用unsafe关键字
// Rust禁止在声明static变量的时候调用普通函数，或者利用语句块调用其他非const代码
// 允许调用const fn
static GLOBAL: i32 = 0;

fn global() {
    println!("{:?}", GLOBAL);
    use std::sync::atomic::AtomicBool;
    static FLAG:AtomicBool = AtomicBool::new(true);
}
// const声明常量,不具备类似语句的模式匹配功能
// 可能会被内联优化
const GLOBAL1:i32 = 0;