package pointer

import (
	"fmt"
	"reflect"
)

// 001 指针*- 取值示例
func PointerGetValue() {
	// go可在变量的初始化时省略变量的类型而由系统自动推断. a := 1 等同于 var a int = 1
	// 声明+赋值(:=)  声明变量a并赋值1,  &: 取地址:获取变量在计算机内存中的地址。
	a := 1
	fmt.Println("a:", a)
	fmt.Println("a的内存地址&a取值为:", &a)            // a的内存地址
	fmt.Println("a的数据类型为:", reflect.TypeOf(a)) // 检查变量类型
	fmt.Println("a的指针取值为:", *&a)               // 变量值不可以定位为指针.
	//fmt.Println("a的内存地址", *a) // 不可以直接取值.

	// 赋值(=) &: 取地址给变量p *p *: 取值
	var b = a                                 // 将a的值重新赋值给d变量
	fmt.Println("b:", b)                      // b为a的内存地址
	fmt.Println("b的数据类型为", reflect.TypeOf(b)) // 内存地址
	fmt.Println("b的内存地址", &b)                 // 0xc000100020
	fmt.Println("b内存地址和值", *&b)               // 0xc000100020

	var c = &a
	fmt.Println("c:", c)                      // c为a的赋值
	fmt.Println("c的数据类型为", reflect.TypeOf(c)) // 内存地址
	fmt.Println("c的内存地址", &c)                 // 0xc000100020
	fmt.Println("c内存地址和值", *c)                // 0xc000100020
}

// 002. 在函数中通过形参改变实参的值时,需要使用指针类型的参数.
func pointerUpdateValue() {
	age := 10
	modifyAge(&age)
	fmt.Println("age的值为:", age)
}
func modifyAge(age *int) {
	fmt.Println("age的传参为:", age)
	*age = 20
}

// 002. 参数传参的作用域范围
func pointerUpdateValueScope() {
	age := 18
	modifyAgeScope(age)
	fmt.Println("age的值为:", age)
}
func modifyAgeScope(age int) {
	fmt.Println("age的传参为:", age)
	age = 20 // age的作用域只在modifyAgeError函数内
	fmt.Println("age赋值后的值为:", age)
}
