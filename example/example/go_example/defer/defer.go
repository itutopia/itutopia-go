package _defer

import "fmt"

// defer (延迟,推迟) 语句被用于预定对一个函数的调用。可以把这类被defer语句调用的函数称为延迟函数。
//defer作用：
// 释放占用的资源
// 捕捉处理异常
// 输出日志

// 如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。

func DeferExample() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}

// 错误拦截
//运行时panic异常一旦被引发就会导致程序崩溃。
//Go语言提供了专用于“拦截”运行时panic的内建函数“recover”。它可以是当前的程序从运行时panic的状态中恢复并重新获得流程控制权。
//注意：recover只有在defer调用的函数中有效。

func DeferRecoverExample() {
	DeferRecover(10)
}

func DeferRecover(i int) {

	// 定义10个int元素的数组
	var arr [10]int

	// 错误拦截要在产生错误前设置
	defer func() {
		//设置recover拦截错误信息
		err := recover()

		//产生panic异常  打印错误信息
		if err != nil {
			fmt.Println(err)
		}
		// 方法内可继续执行,出方法后就不在执行.
		fmt.Println("runtime")
		finalOperation()
	}()
	//根据函数参数为数组元素赋值
	//如果i的值超过数组下标 会报错误：数组下标越界
	arr[i] = 10

	// 使用延迟函数defer定义的方法, 执行完后退出后,本线程后续方法不在执行,主线程main继续执行.
	fmt.Println("runtime")
}

func finalOperation() {
	fmt.Println("finalOperation")
}
