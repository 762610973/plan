// *一个表达式总是会产生一个值,因此必然有类型
// *语句不产生值,类型永远是"()"
// ?表达式加上分好,就变成了一个语句;
// ?；如果把语句放到一个语句块中包起来，那么它就可以被当成一个表达式使用
/*
* 表达式分为左值和右值
    * 左值: 这个表达式可以表达一个内存地址,可以放到赋值运算符左边使用
*/

/*
* 运算表达式
*赋值表达式,赋值表达式的类型为unit,即空的tuple()
*语句块表达式
*/

#[allow(unused)]
fn main() {
    let x = 1;
    let mut y = 2;
    // z是表达式,(y=x)是将x赋值给y,只是一个表达式
    let z = (y = x);
    println!("{:?}", z);
    let x = { println!("hello"); };
    println!("{:?}", x);
    let x = {
        println!("hello");
        5;// 加了分号就是一个语句,而不是一个表达式
    };
    println!("{:?}", x);
    let x = {
        println!("hello");
        5   //这里可以看到x的类型
    };
    println!("{:?}", x);
}


#[allow(unused)]
#[test]
fn test_if_else() {
    let b: i32 = 1;
    if b == 1 {
        println!("{:?}", b);
    } else if b == 2 {
        println!("{:?}", b);
    } else {
        println!("{}", b);
    }
}

#[allow(unused)]
#[test]
fn test_loop() {
    let mut count = 0_i8;
    loop {
        if count > 0 {
            count += 1;
        }
        if count == 2 {
            continue;
        }
        if count == 3 {
            break;
        }
    }
}

#[allow(unused)]
#[test]
fn test_break_label() {
// A counter variable
    let mut m = 1;
    let n = 1;
    'a: loop {
        if m < 100 {
            m += 1;
            println!("{:?}", m);
        } else {
            'b: loop {
                if m + n > 50 {
                    println!("break");
                    break 'a;
                } else {
                    continue 'a;
                }
            }
        }
    }
    // *v是发散类型
    let v = loop {};
}

#[allow(unused)]
#[test]
fn test_while() {
    let mut n = 1;
    while n < 10 {
        n += 1;
        break;
    }
}

#[allow(unused)]
#[test]
fn test_loop_while() {
    let mut x;
    let b: bool = false;
    loop {
        x = 1;
        break;
    }
    while b {
        x = 1;
        break;
    }
    // 编译器会觉得while语句的执行跟条件表达式在运行阶段的值有关,因此它不确定x是否一定会初始化
    // loop和while true语句在运行时没有什么区别,它们主要是会影响编译器内部的静态分析结果
}

#[allow(unused)]
#[test]
fn test_for() {
    let array = &[1, 2, 3, 4];
    for i in array.iter() {
        println!("{}", i)
    }
    for i in array {
        println!("{}", i)
    }
    for i in 0..5 {
        print!("{}", i)
    }
}