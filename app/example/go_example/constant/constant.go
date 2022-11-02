package constant

import (
	"fmt"
	"testing"
)

// 定义星期常量
const (
	Mon = iota + 1
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

// 定义读写常量
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Mon, Tue, Wed, Thu, Fri, Sat, Sun)
	fmt.Println(Mon, Tue, Wed, Thu, Fri, Sat, Sun)
}

func TestConst(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

}
