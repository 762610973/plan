fn main() {
    println!("Hello, world!");
    let  res = add(1, 2);
    println!("res: {}",res)

}

fn add(a:i32,b:i32) ->i32 {
    return a+b;
}