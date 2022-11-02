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
