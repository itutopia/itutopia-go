package os

import (
	"fmt"
	"os"
)

func osExample() {
	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		fmt.Println("hi", os.Args[1])
	}

	// main函数不支持返回值,通过os.Exit(code): code为返回值;
	os.Exit(-1)
}
