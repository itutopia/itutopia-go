package calc

import "testing"

// 1. 单元测试
/*
遵循以下几点：
1.测试文件以源文件名_test.go 命名
2.测试函数名称格式以 TestFuncName(t *testing.T) 命名
3.测试错误 t.Error / t.Fail / t.Errorf 记录
4.测试日志 t.Log 记录
*/

/*
*
点击执行 gu run
*/
func TestDec(t *testing.T) {
	a, b := 3, 1
	// 计算
	except := a - b
	// 函数计算
	fact := Dec(a, b)
	// 验证是否相等
	if except != fact {
		t.Errorf("%d - %d except %d but fact %d \n", a, b, except, fact)
	}
	t.Log("success")
}

/*
* 测试某个函数通常需要覆盖性的提供多组数据，使用测试表
 */
func TestTableDec(t *testing.T) {

	// 定义测试表
	tables := []struct {
		a   int
		b   int
		res int
	}{
		{1, 2, -1},
		{5, 4, 1},
		{24, 4, 20},
		{255, 171, 74},
	}

	for _, except := range tables {
		fact := Dec(except.a, except.b)
		if except.res != fact {
			t.Errorf("%d - %d except %d but fact %d \n", except.a, except.b, except.res, fact)
		}
	}
}

/**
1. 代码覆盖率: 代码覆盖率可以判断测试的完备性，即考虑到了各种可能的类型.即对指定单元执行测试覆盖率语法.
执行语法: go testes [文件路径] -cover
如: go testes /Users/junchao_lee/Documents/itutopia/itutopia_opensouce/itutopia-go/example/example/go_example/testes -cover
*/

func TestTableCover(t *testing.T) {
	// 定义测试表
	testTables := []struct {
		a   int
		b   int
		res int
	}{
		{1, 2, -1},
		{5, 4, 1},
		{24, 4, 20},
		{255, 171, 74},
	}
	for _, except := range testTables {
		fact := Dec(except.a, except.b)
		if except.res != fact {
			t.Errorf("%d - %d except %d but fact %d \n", except.a, except.b, except.res, fact)
		}
	}
}
