// !bool
#[allow(unused)]
fn test_type1() {
	let x = true;
	let y: bool = !x;

	let z = x && y;
	let z = x & y;
	let z = x ^ y;
	fn logical_op(x: i32, y: i32) {
		let z: bool = x < y;
		if x >= y {} else {}
	}
}

// !char
// 描述任意一个unicode字符,占用四个字节
#[allow(unused)]
fn test_char1() {
	let love = '&';
	let c1 = '\n';
	let c2 = '\x7f';
}

// 使用一个字母b在字符或者字符串前面,代表这个字面量存储在u8类型数组中
#[allow(unused)]
fn test_char_u8() {
	let x: u8 = 1;
	let y: u8 = b'a';
	let s: &[u8; 5] = b"hello";
}

#[allow(unused)]
fn test_number() {
	let v1: i32 = 32;// 十进制
	let v2: i32 = 0xFF; // 十六进制
	let v3: i32 = 0o44; //以0o开头代表八进制表示
	let v4: i32 = 0b1001; // 以0b开头代表二进制表示
	let v5 = 0x_1234_abcd; // !可以使用下划线分割
}

#[allow(unused)]
fn test_number1() {
	let x = 4_i32;  // 后置类型
	let x = 400i32;  // 后置类型
	// let x:i32  = 4;
	println!("9 power 3 = {}", x.pow(3));
}

// 可能会溢出
// rustc -C overflow-checks=no test.rs
#[allow(unused)]
fn overflow(m: i8, n: i8) {
	println!("{:?}", m + n);
}

#[allow(unused)]
fn test_overflow() {
	let i = 100_i8;
	println!("{:?}", i.checked_add(i));
}

#[allow(unused)]
fn test_float() {
	let f1 = 123.0f64;
	let f2 = 123.4f64;
	let f3 = 123.4_f64;
	let f4: f64 = 2.;
	let f5 = 12E+99_f64; //科学计数法
}

#[allow(unused)]
fn test_point() {
	//? Box<T> 指向类型T的,具有所有权的指着,有权释放内存
	//? &T 指向类型T的借用指针,也称为引用,无权 释放内存,无权写数据
	//? &mut T 指向类型T的mut型借用指针,无权释放内存,有权写内存
	//? *const T 指向类型T的只读裸指针,没有生命周期信息,无权写数据
	//? *mut T 只想类型T的可读写裸指针,没有生命周期信息,有权写数据


	// Rc<T>  指向类型T的引用计数指针,共享所有权,线程不安全
	// Arc<T> 只想类型T的原子型引用计数指针,共享所有权,线程安全
	// Cow<'a,T> Clone-on-write, 写时复制指针,可能是借用指针,也可能是具有所有权的指针
}

#[allow(unused)]
fn type_change() {
	let v1: i8 = 41;
	//? 显示标记类型转换
	let v2: i16 = v1 as i16;
	println!("{:?}", v2);
	let i = 42;
	// 多次转换
	let p = &i as *const i32 as *mut i32;
	println!("{:p}", p);
}