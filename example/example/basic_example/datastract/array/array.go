package array

import "fmt"

// 数组
func ArrayExample() {
	books := []string{
		"三国演义",
		"红楼梦",
		"西游记",
		"水浒传",
	}

	name := books[len(books)]
	fmt.Printf("name: %s", name)
}

func ArrayExtend() {
	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array)

}
