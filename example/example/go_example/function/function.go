package function

import "fmt"

// 1.函数返回多个值
// 格式: 传参是同类型
// func funcName(x,y [type]) ([type] [type]){ return  y, x}

// 格式: 传参
// func funcName(x [type] , y [type])  ([type] [type]) { return  y, x}

func FuncSwapExample() {
	a, b := swapPosition("Mahesh", "Kumar")
	fmt.Println(a, b)
}

func swapPosition(x, y string) (string, string) {
	return y, x
}

// 注意事项: init函数与import
// init 函数可在package main中，可在其他package中，可在同一个package中出现多次。
// main 函数只能在package main中。

func init() {
	fmt.Println("init")
}
