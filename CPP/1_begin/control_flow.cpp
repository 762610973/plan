#include <iostream>

void add();

void func_for();

void func_input();

void func_if();

int main()
{
	int sum = 0;
	int val = 1;
	while (val <= 10) {
		sum += val;
		++val;    // 前缀递增运算符
		//val++;
	}
	std::cout << "sum of 1 to 10 inclusive is " << sum << std::endl;
	add();
	//func_for();
	//func_input();
	func_if();
	return 0;
}

void add()
{
	int val = 1;
	std::cout << val << std::endl;
	val++;
	std::cout << val << std::endl;
	++val;
	std::cout << val << std::endl;
}

void func_for()
{
	int sum = 0;
	for (int val = 1; val <= 10; ++val) {
		sum += val;
	}
	std::cout << "sum of 1 to 10 inclusive is " << sum << std::endl;
}

void func_input()
{
	int sum = 0;
	int val = 0;
	// 遇到文件结束符停止, Windows: Ctrl+Z, Unix: Ctrl+D
	while (std::cin >> val) {
		sum += val;
	}
	std::cout << "sum is: " << sum << std::endl;
}

void func_if()
{
	int a = 1;
	if (a > 1) {
		std::cout << 1;
	} else if (a == 1) {
		std::cout << 2;
	} else {
		std::cout << 3;
	}
}