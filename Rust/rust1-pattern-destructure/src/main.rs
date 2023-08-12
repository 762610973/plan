mod r#let;

fn main() {
    println!("Hello, world!");
}

#[allow(unused)]
mod test {
    struct T1(i32, char);

    struct T2 {
        item1: T1,
        item2: bool,
    }

    #[test]
    fn t1() {
        /*
        ! 怎么组合,怎么拆解
        */
        let x = T2 {
            item1: T1(0, 'a'),
            item2: false,
        };
        let T2 {
            item1: T1(value1, value2),
            item2: value3,
        } = x;
        println!("{:?} {} {}", value1, value2, value3);
    }
}

#[allow(unused, dead_code)]
mod test_match {
    enum Direction {
        East,
        West,
        South,
        North,
    }

    fn print(x: Direction) {
        match x {
            Direction::North => {
                println!("east");
            }
            Direction::West => {
                println!("west");
            }
            Direction::South => {
                println!("south");
            }
            Direction::East => {
                println!("east");
            }
        }
        /* match x {
             Direction::East => {
                 println!("east");
             }
             _ => {
                 println!("other");
             }
         }*/
    }

    #[test]
    fn print_direction() {
        let x = Direction::East;
        print(x);
        let x = Direction::West;
        print(x);
        let x = Direction::North;
        print(x);
        let x = Direction::South;
        print(x);
    }

    struct P(f32, f32, f32);

    fn calc(arg: P) -> f32 {
        let P(x, _, _) = arg;
        x
    }

    fn calc1(P(x, y, z): P) -> f32 {
        x + y + z
    }

    #[test]
    fn f2() {
        let x = (1, 2, 3);
        // 使用..来忽略其他的
        let (a, ..) = x;
        println!("{:?}", a);
        let (a, .., b, c) = x;
        println!("{}{}{}", a, b, c);
    }
}

#[allow(unused, dead_code)]
mod match_expression {
    enum Direction {
        East,
        West,
        South,
        North,
    }

    fn direction_to_int(x: Direction) -> i32
    {
        let res = match x {
            Direction::East => 10,
            Direction::West => 20,
            Direction::South => 30,
            // *可以写多个相同的条件,但只会匹配第一个
            Direction::North => 40,
            Direction::North => 40,
        };

        return res;
    }

    #[test]
    fn main() {
        let x = Direction::East;
        let s = direction_to_int(x);
        println!("{}", s);
    }

    // * rust的match还可以匹配值
    #[test]
    fn test1() {
        let x = 100;
        match x {
            100 => println!("100"),
            // 可以使用运算符来匹配 多个条件
            2 | 3 => println!("2"),
            // 还可以使用范围匹配
            4..=10 => {
                println!("111");
            }
            _ => {
                println!("200");
            }
        }
    }
}

#[allow(unused, dead_code)]
mod guards {
    enum Optional {
        Value(i32),
        Missing,
    }

    #[test]
    fn t1() {
        let x = Optional::Value(3);
        match x {
            // 可以写多个相同的match,会依次匹配
            Optional::Value(i) if i > 5 => println!("no"),
            Optional::Value(..) => println!("another"),
            Optional::Missing => println!("missing"),
        }
    }
    /*
    - 编译器会保证match的所有分支合起来一定覆盖了目标的所有可能取值范围
    - 但是并不会保证各个分支是否会有重叠的情况
    */
}

#[allow(unused, dead_code)]
mod bind {
    #![feature(exclusive_range_pattern)]

    #[test]
    fn test() {
        let x = 2;
        match x {
            // - 使用@符号绑定变量
            e @ 1..=5 => println!("{}", e),
            _ => (),
        }
    }

    fn deep_match(v: Option<Option<i32>>) -> Option<i32> {
        match v {
            // r 绑定到的是第一层 Option 内部,r 的类型是 Option<i32>
            // 与这种写法含义不一样：Some(Some(r)) if (1..10).contains(r)
            Some(r @ Some(1..=10)) => r,
            _ => None,
        }
    }

    #[test]
    fn main() {
        let x = Some(Some(5));
        println!("{:?}", deep_match(x));
        let y = Some(Some(100));
        println!("{:?}", deep_match(y));
    }
}

#[allow(unused, dead_code)]
mod ref_mut {
    use std::any::type_name;

    fn print_type_name<T>(_: T) {
        println!("{}", type_name::<T>())
    }

    // - 需要绑定的是被匹配对象的引用,使用ref关键字
    // - 模式匹配有时候可能发生变量的所有权转移
    // - mut修饰变量绑定
    // - mut可以修饰指针(引用)
    #[test]
    fn t1() {
        let x = 5_i32;
        match x {
            // !这里可以直接使用一个变量
            ref r => println!("got a reference to {}", r)
        }
        // ref 是模式的一部分
        let ref x = 5_i32;
        println!("{:?}", x);
        let y = &5_i32;
        println!("{}", x == y);
        print_type_name(x);
        print_type_name(y);
    }
}