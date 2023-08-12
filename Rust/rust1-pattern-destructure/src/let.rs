#[allow(unused, dead_code)]
mod if_let {
    #[test]
    fn test1() {
        let x = Option::Some(123);
        // - match匹配
        match x {
            Some(x) => {}
            _ => {}
        }
        // - 第二种方法:使用Option的方法 ,但是在运行期判断了两次x是否有值,降低了执行效率
        if x.is_some() {
            let v = x.unwrap();  //取出内部的数据
            // do some thing with (x)
        }
        // - 第三种方法: 使用 `if let`,if let 只匹配感兴趣的某个特定的分支
        if let Some(val) = x {
            // do some thing with (x)
        }
    }

    #[test]
    fn test2() {
        enum E {
            A(i32),
            B,
            C,
            D,
        }
        let x = E::A(3);
        let r =
            if let
                E::A(x1) = x {
                x1
            } else {
                2
            };
        let r = match x {
            E::C | E::D => 1,
            E::A(x) => x,
            _ => 2,
        };
    }
}