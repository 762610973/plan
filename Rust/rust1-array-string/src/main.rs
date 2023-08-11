mod string;

fn main() {
    println!("Hello, world!");
}

#[allow(unused)]
mod array {
    #[test]
    fn t1() {
        let a1: [i32; 5] = [1, 2, 3, 4, 5];
        // 所有的元素,初始化为同样的数据
        let a2: [i32; 500] = [0; 500];
    }

    #[test]
    fn t2() {
        let mut a1: [i32; 5] = [1, 2, 3, 4, 5];
        let a2: [i32; 5] = [6, 7, 8, 9, 10];
        // 同类型的数组之间可以互相赋值
        // *同类型: 元素类型和元素个数都完全相同,与go一样
        a1 = a2;
        println!("new array{:?}", a1);
        let add_res = a2[0] + a2[1];
        println!("{:?}", add_res);
        for i in &a1 {
            println!("{:?}", i);
        }
    }

    #[test]
    fn t3() {
        let v1 = [1, 2, 3];
        let v2 = [1, 2, 3];
        println!("{:?}", v1 == v2);
        let v3 = [0_i32; 10];
        for i in v3 {
            print!("{:?}", i);
        }
        for i in v3.iter() {
            println!("{:?}", i);
        }
    }

    #[test]
    fn t4() {
        let v = [[0, 0], [1, 2]];
        for i in v {
            println!("{:?}", i);
        }
    }

    use std::mem;
    use crate::main;

    #[test]
    fn array_slice() {
        /*
        * 数组切片是指向一个数组的指针,不止包含指向数组的指针,切片本身还带有长度信息
        */
        fn mut_array(a: &mut [i32]) {
            if a.len() >= 3 {
                a[2] = 5;
            }
        }
        // 数组指针大小是8
        println!("size of &[i32;10] : {:?}", mem::size_of::<&[i32; 10]>());

        println!("size of &[i32] : {:?}", mem::size_of::<&[i32]>());

        let mut v: [i32; 3] = [1, 2, 3];
        {
            let s: &mut [i32; 3] = &mut v;
            println!("{:?}", mem::size_of::<&mut [i32; 3]>());

            mut_array(s);
        }
        println!("{:?}", v);
    }
}

#[allow(unused)]
mod fat_pointer {
    use std::mem;


    /*
不定长数组类型[T]在编译阶段是无法判断该类型占用空间的大小的
目前我们不能在栈上声明一个不定长大小数组的变量实例
也不能用它作为函数的参数、返回值。但是，指向不定长数组的胖指针的大小是确定的
&[T]类型可以用做变量实例、函数参数、返回值*/
    // *&[T]类型占用了两个指针大小的内存空间
    fn raw_slice(arr: &[i32]) {
        unsafe {
            let (val1, val2): (usize, usize) = mem::transmute(arr);
            println!("Value in raw pointer:");
            println!("value1: {:x}", val1);
            println!("value2: {:x}", val2);
        }
    }

    fn t1() {
        fn size<T>() {}
        size::<i32>()
    }

    #[test]
    fn t5() {
        let arr = [1, 2, 3, 4, 5];
        let address = &arr;
        println!("address of arr: {:p}", address);
        raw_slice(address as &[i32]);
    }
    /*
    ? 对于DST类型，Rust有如下限制
    * 只能通过指针来间接创建和操作DST类型，&[T]Box<[T]>可以，[T]不可以
    * 局部变量和函数参数的类型不能是DST类型
    * enum中不能包含DST类型，struct中只有最后一个元素可以是DST，其他地方不行，如果包含有DST类型，那么这个结构体也就成了DST类型
    *
    */
}

#[allow(unused)]
mod range {
    #[test]
    fn t1() {
        use std::iter::Iterator;
        let r = (1_i32..11).rev().map(|i| i * 10);
        for i in r {
            print!("{}\t", i)
        }
    }

    #[test]
    fn t2() {
        let arr: [i32; 5] = [1, 2, 3, 4, 5];
        // range to
        let s1 = &arr[..2];
        println!("{:?}", s1);

        // range from
        let s2 = &arr[2..];
        println!("{:?}", s2);

        // full range
        let s3 = &arr[..];
        println!("{:?}", s3);
        for i in 3..=14 {}

        let option = arr.get(10);
        match option {
            Some(x) => {
                x;
            }
            None => {
                println!("none");
            }
        };
    }
}

/*
一般情况下，Rust不鼓励大量使用“索引”操作
*/