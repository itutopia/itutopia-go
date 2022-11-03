package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func GoRoutineFuncExample() {
	for i := 0; i <= 100; i++ {
		go goFun(i)
	}
	time.Sleep(time.Second)

}

func goFun(i int) {
	fmt.Println("goroutine", i, "...")
}

func GoDeferExample() {
	defer fmt.Println("A.defer")

	// 匿名内部类
	func() {
		defer fmt.Println("B.defer")
		runtime.Goexit() // 终止当前 goroutine, import "runtime"
		fmt.Println("B") // 不会执行
	}()
	fmt.Println("A") // 不会执行
}
