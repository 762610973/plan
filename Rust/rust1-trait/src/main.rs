fn main() {
    println!("Hello, world!");
}

// 成员方法
#[allow(unused)]
mod function {
    trait Shape {
        // ?所有的trait中都有一个隐藏的类型Self,代表当前这个实现了此trait的具体类型
        // ?trait中定义的函数，也可以称作关联函数
        fn area(&self) -> f64;
    }

    trait T {
        fn method1(self: Self);
        fn method2(self: &Self);
        fn method3(self: &mut Self);
    }

    trait TT {
        fn method1(self);
        fn method2(&self);
        fn method3(&mut self);
    }

    // T==TT
    struct Circle {
        radius: f64,
    }

    impl Shape for Circle {
        // Self 类型就是Circle
        // self 的类型就是&Self,即&Circle
        fn area(&self) -> f64 {
            std::f64::consts::PI * self.radius * self.radius
        }
    }

    fn main() {
        let c = Circle {
            radius: 2_f64,
        };
        println!("{:?}", c.area());
    }
}

/*
?方法: 具有receiver参数的函数(第一个参数是Self相关的类型,且命名为self)
?静态函数: 没有receiver参数,通过类型加双冒号::的方法来调用
* rust中,函数和方法没有本质区别
? 大写的Self是类型名
? 小写的self是变量名, self参数可以使Box指针类型self: Box<Self>

?内在方法(匿名实现):
impl Circle {
    fn()
}
*/

//? trait中可以包含方法的默认实现,如果这个方法已经在trait中有了方法体,那么针对具体类型实现
//? 可以选择不用重写, 当然也可以override
//? impl的对象甚至可以使trait

mod impl_trait {
    trait Shape {
        fn area(&self) -> f64;
    }

    trait Round {
        fn get_radius(&self) -> f64;
    }

    struct Circle {
        radius: f64,
    }

    impl Round for Circle {
        fn get_radius(&self) -> f64 {
            self.radius
        }
    }

    impl Circle {
        fn area(&self) -> f64 {
            std::f64::consts::PI * self.radius * self.radius
        }
    }

    impl Shape for dyn Round {
        fn area(&self) -> f64 {
            std::f64::consts::PI * self.get_radius() * self.get_radius()
        }
    }

    #[test]
    fn test1() {
        let b = Box::new(Circle { radius: 4_f64 }) as Box<dyn Round>;
        b.area();
    }

    #[test]
    fn test2() {
        let c = Circle {
            radius: 5_f64,
        };
        c.area();
    }
}

mod static_function {
    struct T(i32);

    impl T {
        // 即使类型是&Self, 但不是形参名不是self,调用只能通过::
        fn func1(this: &Self) {
            println!("value{:?}", this.0);
        }
    }

    fn main() {
        let x = T(42);
        T::func1(&x);
    }

    // ? trait中也可以定义静态函数
    pub trait Default {
        fn default() -> Self;
    }
}
