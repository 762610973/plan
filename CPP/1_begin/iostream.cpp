#include <iostream>

// 声明函数
void work();

int main()
{
	//标准库中定义的所有名字都在命名空间std中
	(std::cout << "enter two numbers: ") << std::endl;
	int v1 = 0, v2 = 0;
	std::cin >> v1 >> v2;
	std::cout << "the sum of " << v1 << " and " << v2 << " is " << v1 + v2 << std::endl;
	work();

	return 0;
}

void work()
{
	std::clog << "test std::clog" << std::endl;
	std::cerr << "test std::cerr" << std::endl;
}
