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

// 静态函数(go中的方法)
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

    // *trait中也可以定义静态函数
    pub trait Default {
        fn default() -> Self;
    }
}

// 孤儿规则: impl块要么与trait的声明在同一个的crate中，要么与类型的声明在同一个crate中
#[allow(unused)]
mod expend_function {
    trait Double {
        fn double(&self) -> Self;
    }
    impl Double for i32 {
        fn double(&self) -> i32 { *self * 2 }
    }
    fn main() {
    // 可以像成员方法一样调用
        let x : i32 = 10.double();
        println!("{}", x);
    }
}
#[allow(unused)]
mod fully_call_function {
    trait Cook {
        fn start(&self);
    }

    trait Wash {
        fn start(&self);
    }

    struct Chef;

    impl Cook for Chef {
        fn start(&self) {
            println!("Cook::start");
        }
    }

    impl Wash for Chef {
        fn start(&self) {
            println!("Wash::start");
        }
    }
    #[test]
    fn main() {
        let me = Chef;
        // me.start();  不能确定调用哪个
        <Chef as Wash>::start(&me);
        <dyn Cook>::start(&me);
    }
    struct T(usize);
    impl T {
        fn get1(&self) -> usize {self.0}
        fn get2(&self) -> usize {self.0}
    }
    fn get3(t: &T) -> usize { t.0 }
    fn check_type( _ : fn(&T)->usize ) {}

    #[test]
    fn main1() {
        check_type(T::get1);
        check_type(T::get2);
        check_type(get3);
    }
}

mod bind {
    // *泛型约束
    use std::fmt::Debug;
    // ?要求类型T实现Debug这个trait
    fn my_print<T : Debug>(x:T) {
        println!("{:?}", x);
    }
    fn my_print1<T>(x:T) where T:Debug {
        println!("{:?}", x);
    }
    // *trait允许继承
    trait Base{}
    // 满足derived,也必然满足base
    trait Derived: Base{}
}
#[allow(unused)]
mod derive {
    // *derive自动实现trait,编译器机械化得重复模板实现trait
    #[derive(Copy, Clone, Default, Debug, Hash, PartialEq, Eq, PartialOrd, Ord)]
    struct Foo {
        data : i32
    }
    fn main() {
        let v1 = Foo { data : 0 };
        let v2 = v1;
        println!("{:?}", v2);
    }
}
// * trait别名
mod trait_alas {
    use std::future::Future;
    pub trait Service {
        type Request;
        type Response;
        type Error;
        type Future: Future;
        fn call(&self, req: Self::Request) -> Self::Future;
    }
    // trait HttpService = Service<
    //     Response = http::Response,
    //     Error = http::Error, Future=(), Request=()>;
}

mod std_trait {
    /*
    * 只有实现了Display trait的类型,才能用{}格式控制打印出来
    * 只有实现了Debug trait的类型,才能用{:?}和{:#?}格式控制打印出来,编译器提供了自动derive的功能
    */
}