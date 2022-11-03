package calc

import "fmt"

//4.案例测试
/*
1. 案例测试遵循以下几点：
	1. 测试文件以源文件_example_test.go 命名
	2.测试函数名称格式以 ExampleFuncName() 命名
	3.测试结果打印 以下列格式注释：
		// Output:
		// 结果1
 		// 结果2

2. 执行语法: go testes calc
如:go testes /Users/junchao_lee/Documents/itutopia/itutopia_opensouce/itutopia-go/example/example/go_example/testes
*/
func ExampleDec() {
	tables := []struct {
		a int
		b int
	}{
		{5, 4},
		{24, 10},
	}
	for _, except := range tables {
		fmt.Println(Dec(except.a, except.b))
	}

	// Output:
	// 1
	// 14
}
