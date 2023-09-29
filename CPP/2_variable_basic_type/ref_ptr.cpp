#include <iostream>

void pointer();

void reference();

int main()
{
	//reference();
	pointer();

	return 0;
}

void reference()
{
	int val = 1;
	// 引用必须被初始化
	int &ref_val = val;
	int &r2 = val;
	// ref_val指向val, 是val的另一个名字, 引用即别名
	// 引用不是对象, 是为一个已经存在的对象所起的另外一个名字
	// 为引用赋值, 实际上是把值赋给了与引用绑定的对象, 不能指向引用的引用
	// 不能含有指向字面量的引用
	//对象是具有某种数据类型的内存空间
	std::cout << val << std::endl;
	std::cout << ref_val << std::endl;
	ref_val = 200;
	std::cout << ref_val << std::endl;
	// val也发生了变化, 另一个引用也发生了变化
	std::cout << val << std::endl;
	std::cout << r2 << std::endl;
}

void pointer()
{
	/*
	 * 指针本身就是一个对象, 允许对指针赋值和拷贝,而且在指针的生命周期内它可以先后指向几个不同的对象
	 * 指针无需在定义时赋初值
	 * 指针存放某个对象的地址
	 * 不能定义指向引用的指针
	 * */
	int val = 1;
	int val2 = 2;
	// p存放val的地址/p是指向变量val的指针
	int *p = &val;    // 等价于: int *p; p = &val;
	std::cout << p << std::endl;
	// 改变变量val的值, p不会发生变化
	val = 2;
	std::cout << p << std::endl;
	*p = 3;
	std::cout << p << "\tchange val by *p: " << val << *p << std::endl;
	std::cout << "int *p; *紧随类型名出现, 是声明的一部分, p是一个指针" << std::endl;

	// 改变p的指向, p的地址会发生变化
	p = &val2;
	std::cout << p << "\t" << *p << std::endl;
	if (p != nullptr) {
		std::cout << "p is not nullptr"<<std::endl;
	}
	// void* 是一种特殊的指针类型, 可用于存放任意对象的地址. 
	int *p1, p2;
}