package constant

import (
	"fmt"
	"testing"
	"unsafe"
)

//常量是一个简单值的标识符，在程序运行时，不会被修改的量。
//常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

// 定义显示类型常量格式:
// const identifier [type] = value

const COMMUNITY string = "itutopia"
const CHINA = "中华人民共和国"

func ConstantExample() {

	// 1. 定义常量PAI
	const PAI = 3.1415926
	var radius float64 = 2
	circleArea(PAI, radius)

	// 2. 定义多个常量
	const a, b, c = 1, false, "北京"
	fmt.Println(a, b, c, COMMUNITY, CHINA)
}

func circleArea(PAI float64, r float64) {
	var area float64
	area = PAI * (r * r)
	fmt.Println(area)
}

// 2.全局常量. 做枚举, iota自增长,默认从0开始
const (
	UnKnow = iota
	Female
	Male
)

// Go 语言中 unsafe.Sizeof() 函数：
// 1. 对不同长度的字符串，unsafe.Sizeof() 函数的返回值都为 16，这是因为 string 类型对应一个结构体，该结构体有两个域，第一个域指向该字符串的指针，第二个域为字符串的长度，每个域占 8 个字节，但是并不包含指针指向的字符串的内容，这就解释了unsafe.Sizeof() 函数对 string 类型的返回值始终是16。
// 2. 对不同长度的数组，unsafe.Sizeof() 函数的返回值随着数组中的元素个数的增加而增加，这是因为unsafe.Sizeof() 函数总是在编译期就进行求值，而不是在运行时，这就意味着，unsafe.Sizeof() 函数的返回值可以赋值给常量，在编译期求值，意味着可以获得数组所占的内存大小，因为数组总是在编译期就指明自己的容量，并且在以后都是不可变的。
// 3. 对所含元素个数不同的切片，unsafe.Sizeof() 函数的返回值都为 24，这是因为对切片来说，unsafe.Sizeof() 函数返回的值对应的是切片的描述符，而不是切片所指向的内存的大小，因此都是24。
const (
	a = "go"
	b = len(a)
	// 字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
	c = unsafe.Sizeof(a)
)

// 3. 常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。 常量表达式中，函数必须是内置函数，否则编译不过.
func ConstantArithmetic(t *testing.T) {
	println("c", c)
	// 定义局部常量, iota默认为0开始
	const (
		Mon = iota + 1
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)
	println(a, b, c, UnKnow, Female, Male)

	t.Log(Mon, Tue, Wed, Thu, Fri, Sat, Sun)
	println(Mon, Tue, Wed, Thu, Fri, Sat, Sun)
}

// 定义读写常量 +  iota和表达式
type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

// 第一行分别为1,2,3; 第二行分别加1,2,3.即2,3,4
const (
	Apple, Banana, Cherimoya = iota + 1, iota + 2, iota + 3
	Durian, Elderberry, Fig
)

func IotaConstantExample() {
	fmt.Println(IgEggs | IgChocolate | IgNuts | IgStrawberries | IgShellfish)
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)

	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConst(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
