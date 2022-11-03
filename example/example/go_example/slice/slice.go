package slice

import "fmt"

// slice(切片) -动态数组
// Go语言 切片(slice)是对数组的抽象。
// Go数组的长度不可改变，在特定场景中这样的集合就不太适用.
// Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，
// 可以追加元素，在追加时可能使切片的容量增大。

// 1. 声明一个未指定大小的数组来定义切片, 切片不需要说明长度。
//	var identifier []type

// 2. 使用make()函数来创建切片:
//	var slice1 []type = make([]type, len)
// 简写
// 	slice1 := make([]type, len)

// 3. 使用make()函数来创建切片指定容量，len是数组的长度并且也是切片的初始长度, 其中capacity为可选参数。
// slice2 := make([]T, length, capacity)

func SliceExample() {

	// 1. 直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	s := []int{1, 2, 3}

	s1 := s[1:2]
	fmt.Println(s, s1)

	// 2.定义指定容量和扩容的切片
	var number = make([]int, 3, 5)
	printSlice(number)

	// 3.空(nil)切片 : 一个切片在未初始化之前默认为 nil，长度为 0.
	var x []int
	printSlice(x)
	if x == nil {
		fmt.Printf("切片是空的")
	}

	// 4.创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含) */
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0,(不包含)索引3 */
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为len(s) ,索引4(不包含)开始 */
	fmt.Println("numbers[4:] ==", numbers[4:])

	// 定义并赋值切面number1
	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	// 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
	// append()函数
	// copy()函数
	var number4 []int
	printSlice(number4)

	/* 允许追加空切片 */
	number4 = append(number4, 0)
	printSlice(number4)

	/* 向切片添加一个元素 */
	number4 = append(number4, 1)
	printSlice(number4)

	/* 同时添加多个元素 */
	number4 = append(number4, 2, 3, 4)
	printSlice(number4)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers5 := make([]int, len(number4), (cap(number4))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers5, number4)
	printSlice(numbers5)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
