fn main() {
    println!("Rust的类型 系统是一种代数类型系统😇");
    println!("一个类型的所有取值的可能性叫做这个类型的`基数`");
    println!("bool类型的基数是2,uint()的基数就是1");
    /*
    - 如果两个类型的基数是一样的,则携带的信息量是一样的,则它们是同构的
    -  tuple,struct,tuple struct实质上是同样的内存布局
    - enum类型可以类比为代数运算中的求和
    - tuple、struct可以类比为代数运算中的求积
    - 数组可以类比为代数运算中的乘方
    */
    let unicode = "\u{1F970}";
    println!("{}", unicode);
}

enum Never {
    // 无法构造变量
}

/*
- never类型是Rust类型系统中不可缺少的一部分
* never类型 -> !,没有值,永远不会有结果,!类型表达式可以强转为任何其他类型
*/

/*
如果从逻辑上说，我们需要一个变量确实是可空的，那么就应该显式标明其类型为Option<T>，否则应该直接声明为T类型
相对于裸指针，使用Option包装的指针类型的执行效率不会降低，这是“零开销抽象”

*/