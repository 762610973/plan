#[test]
#[allow(unused)]
fn test_tuple() {
	let a = (1_i32, false);
	let b = ("a", (1_i32, 3.4_f64));
	let c = (0, );  //? 增加逗号,区分表达式和元组
	let d = (0);      //? 括号表达式
	let empty = ();     //? 单元类,内存为0
}

#[test]
#[allow(unused)]
fn access_tuple() {
	let a = (2_i32, 3_f64);
	let (x1, x2) = a;   //? 模式匹配 pattern destructuring
	println!("{:?}", a);
	let x1 = a.0;
	let x2 = a.1;
}

#[test]
#[allow(unused)]
fn test_mem_type() {
	println!("size of i8 {}", std::mem::size_of::<i8>());
	println!("size of char {}", std::mem::size_of::<char>());
	println!("size of '()' {}", std::mem::size_of::<()>());
}

#[test]
#[allow(unused)]
fn test_struct() {
	#[derive(Debug)]
	struct Point {
		x: i32,
		y: i32,
	}
	let p: Point = Point {
		x: 0,
		y: 111,
	};
	println!("point  is at {},{}", p.x, p.y);
	println!("point is {:?}", p);
	println!("point is {:#?}", p);

	let x = 10;
	let y = 20;
	let p = Point { x, y }; //?变量名和字段名字一样,可以省略字段名
}

#[test]
#[allow(unused)]
fn test_access_struct() {
	#[derive(Debug)]
	struct Point {
		x: i32,
		y: i32,
	}
	let p = Point { x: 20, y: 10 };
	let Point { x: px, y: py } = p;
	println!("point is at {} {}", px, py);
	// *在模式匹配的时候,如果新的变量名刚好和成员名字相同,可以使用简写方式
	let Point { x, y } = p;
	println!("Point is at {} {}", x, y);
}

#[allow(unused)]
#[derive(Debug)]
struct Point3d {
	x: i32,
	y: i32,
	z: i32,
}

fn default_value() -> Point3d {
	return Point3d {
		x: 0,
		y: 0,
		z: 0,
	};
}

#[test]
#[allow(unused)]
fn test_set_value() {
	let origin = Point3d {
		x: 100,
		y: 0,
		// *语法糖,允许使用一种简化的语法赋值使用另外医德struct的部分成员
		..default_value()
	};
}

// *内部没有成员

struct Empty;

struct Empty1();

struct Empty2 {}

#[allow(unused)]
fn test_empty_struct() {
	let e1 = Empty;
	let e2 = Empty1;
	let e3 = Empty2 {};
}

#[allow(unused)]
// tuple struct,成员没有名字
#[derive(Debug)]
struct Color(i32, i32, i32);

#[test]
#[allow(unused)]
fn test_tuple_struct() {
	let c = Color(1, 2, 3);
	let i1 = c.0;
	let i2 = c.2;
	// *可以通过下标赋值
	let c1 = Color {
		0: 1,
		1: 2,
		2: 12341,
	};
	println!("{:#?}", c1);
	println!("{:?}", c1);
}

#[allow(unused)]
#[derive(Debug)]
enum Number {
	Int(i32),
	Float(f32, f64),
	Empty,
	Struct { x: i32, y: i32 },
}

#[test]
#[allow(unused)]
fn test_enum() {
	let n: Number = Number::Empty;
	println!("{:}", std::mem::size_of::<Number>());
	enum Animal {
		Dog = 1,
		Cat = 200,
		Tiger,
	}
	let d = Animal::Dog as i8;
	println!("{:?}", d);
}

#[allow(unused)]
fn test_option() {
	// *Some是一个函数类型
	let b = Some(3_i32);
}

// *递归引用
struct Recursive {
	data: i32,
	rec: Box<Recursive>,
}