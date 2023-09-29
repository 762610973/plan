#include <iostream>

void init_val();

int i = 42;

int main()
{
	int a = 20;
	int b = 024;
	int c = 0x14;
	short d = 20;
	long e = 20;
	long long f = 20;
	std::cout << sizeof(a) << std::endl;
	std::cout << sizeof(d) << std::endl;
	std::cout << sizeof(e) << std::endl;
	std::cout << sizeof(f) << std::endl;
	std::cout << "\'" << std::endl;
	std::string str = "string";
	std::cout << str << std::endl;
	std::cout << str.append("-string") << std::endl;
	char c1 = 'c';
	char16_t c2 = 'c';
	char32_t c3 = 'c';
	std::cout << sizeof(c1) << std::endl;
	std::cout << sizeof(c2) << std::endl;
	std::cout << sizeof(c3) << std::endl;
	init_val();
	int t;
	// 此处的t不是零值
	std::cout << t << std::endl;
	return 0;
}
// 基本内置类型
// 算数类型: arithmetic type
// 空类型: void
// char, signed char, unsigned char
// nullptr是指针字面值


/*
 * 明确知晓数值不可能为负时, 选用无符号类型
 * 执行浮点运算选用double
 * 在c++中, 初始化和赋值是两个完全不同的操作
 * 如果内置类型的变量未被显示初始化, 它的值由定义的位置决定. 定义与任何函数体之外的变量被初始化为0
 * 建议初始化每一个内置类型的变量
 * 未初始化的变量含有一个不确定的值, 使用未初始化变量的值是一种错误的编程行为并且很难调试
 * 作用域中一旦声明了某个名字, 它所嵌套着的所有作用域中都能访问该名字
 * */

void init_val()
{
	int t;
	// 此处的t是一个int类型的零值
	std::cout << t << std::endl;
	std::cout << "四种初始化方式" << std::endl;
	int a = 11;
	//c++ 11之后, 一律使用{...}语法初始化类型, 统一初始化的优点是可以阻止它窄化.
	int b = {11};
	int c{11};
	int d(11);
	int e = (11);
	std::cout << a << std::endl;
	std::cout << b << std::endl;
	std::cout << c << std::endl;
	std::cout << d << std::endl;
	std::cout << e << std::endl;
}

// 变量能且只能被定义一次, 但是可以被多次声明
extern int i;     // 声明i而非定义i
int j;            // 声明并定义j
extern int k = 3; // 定义