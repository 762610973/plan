fn main() {
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