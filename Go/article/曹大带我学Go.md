# Go调度的本质是一个生产-消费流程
- 生产数的goroutine放到可运行队列中
  - runnext: 只能只想一个goroutine, 是一个特殊的队列.
  - local queue: 大小为256的数组, 实际上用head和tail指针把它当成一个环形数组使用.
  - global queue
- 如果runnext为空, goroutine顺利放入runnext, 以最高优先级得到运行, 优先被消费.
- Go程序启动创建P, 创建初始的m0, m0启动一个调度循环, 不断地找g, 运行, 再找g
- 随程序运行, m更多地被创建出来, 生产者不断地生产g, m的调度循环不断地消费g
# 迷惑的goroutine执行顺序
```go
func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
// 顺序是9,0,1,2,3,4,5,6,7,8
```
- 本地只有一个p,for循环中生产出的goroutine都会进入到p的runnext和local queue
  - i从1开始, runnext已经有goroutine在了, 这是会把old goroutine移到p的本地队列中, 再把new goroutine放到runnext中, 重复这个过程.
  - 最后i为9, 新goroutine被放到runnext, 其余goroutine都在本地队列.
- go1.13的time包生产一个timerproc的goroutine用于唤醒挂在timer上的时间未到期的goroutine.
- go1.14去掉了这个用于唤醒的goroutine, 取而代之在调度循环的各个地方, sysmon里都是唤醒timer的代码, timer唤醒更及时.

# 初识AST的威力
> abstract syntax tree
- ![](../../images/go/规则二叉树.png)
```go
func main() {
	m := map[string]int{"orders": 10000, "driving_years": 18}
	rule := "orders > 1000 && driving_years > 5"
	fmt.Println(Eval(m, rule))
}

func Eval(m map[string]int, expr string) (bool, error) {
	// 解析表达式
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		return false, err
	}
	fileSet := token.NewFileSet()
	// 打印 ast
	if err = ast.Print(fileSet, exprAst); err != nil {
		return false, err
	}

	return judge(exprAst, m), nil
}

// dfs
func judge(bop ast.Node, m map[string]int) bool {
	// 叶子结点
	if isLeaf(bop) {
		// 断言成二元表达式
		expr := bop.(*ast.BinaryExpr)
		x := expr.X.(*ast.Ident)    // 左边
		y := expr.Y.(*ast.BasicLit) // 右边

		// 如果是 ">" 符号
		if expr.Op == token.GTR {
			left := m[x.Name]
			right, _ := strconv.Atoi(y.Value)
			return left > right
		}
		return false
	}

	// 不是叶子节点那么一定是 binary expression（我们目前只处理二元表达式）
	expr, ok := bop.(*ast.BinaryExpr)
	if !ok {
		println("this cannot be true")
		return false
	}

	// 递归地计算左节点和右节点的值
	switch expr.Op {
	case token.LAND:
		return judge(expr.X, m) && judge(expr.Y, m)
	case token.LOR:
		return judge(expr.X, m) || judge(expr.Y, m)
	}

	println("unsupported operator")
	return false
}

// 判断是否是叶子节点
func isLeaf(bop ast.Node) bool {
	expr, ok := bop.(*ast.BinaryExpr)
	if !ok {
		return false
	}

	// 二元表达式的最小单位，左节点是标识符，右节点是值
	_, okL := expr.X.(*ast.Ident)
	_, okR := expr.Y.(*ast.BasicLit)
	if okL && okR {
		return true
	}

	return false
}
```
# 哪里来的goexit