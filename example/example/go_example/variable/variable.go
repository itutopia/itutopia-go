package variable

import (
	"fmt"
	"reflect"
)

// 声明变量示例
func VariableExample() {

	//第一种: 使用默认值
	var i int
	fmt.Printf("a = %d\n", i)
	var b bool
	fmt.Printf("b = %d\n", b)

	//第二种 指定数据类型并赋值
	var c int = 10
	fmt.Printf("c = %d\n", c)

	//第三种 省略后面的数据类型,系统自动匹配类型
	var x = 20
	fmt.Printf("x = %d\n", x)

	//第四种 省略var关键字
	d := 3.14
	fmt.Println(reflect.TypeOf(d))
	fmt.Printf("d = %f\n", d)
}

// 定义全局变量
var x, y int

// 分解定义写法,一般用于声明全局变量
var (
	a int
	b bool
)
var c, d int = 1, 2
var e, f = 123, "itutopia"

// 多变量声明
func VariableMultiExample() {

	g, h := 123, "需要在func函数体内实现"

	//实际上7的赋值被废弃，变量 _ 不具备读特性
	_, value := 7, 5
	fmt.Println(x, y, a, b, c, d, e, f, g, h, value)

}
