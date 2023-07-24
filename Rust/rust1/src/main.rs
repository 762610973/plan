// use std::prelude::*;
// crate: 完整的编译单元，生成一个lib或者exe可执行文件
// mod: crate内部，即namespace
// println！ 宏 --> 更好地编译器检查
fn main() {
    println!("Hello, world!");
    let  res = add(1, 2);
    println!("res: {}",res);
    println!("{}",1);
    println!("{:o}",9);
    println!("{:x}",255);
    println!("{:X}",255);
    println!("{:p}",&0);
    println!("二进制：{:b}",8);
    println!("{:?}","debug");
    println!("{:#?}",("带有缩进","带有换行的debug"));
    println!("{a} {b}",a=1,b=2);

}

fn add(a:i32,b:i32) ->i32 {
    return a+b;
}