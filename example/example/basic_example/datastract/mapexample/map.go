package mapexample

import "fmt"

// map和slice类似，只不过是数据结构不同。

func MapExample() {
	//第一种: 声明map: kv对, 指定分配空间大小10.
	var map1 map[string]string
	//在使用map前，需要先make，make的作用就是给map分配数据空间
	map1 = make(map[string]string, 10)
	map1["one"] = "php"
	map1["two"] = "golang"
	map1["three"] = "java"
	printMap(map1)
	fmt.Println(map1) //mapexample[two:golang three:java one:php]

	//第二种: 声明并赋值.
	test2 := make(map[string]string)
	test2["one"] = "php"
	test2["two"] = "golang"
	test2["three"] = "java"
	fmt.Println(test2) //mapexample[one:php two:golang three:java]

	//第三种声明
	test3 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "java",
	}
	fmt.Println(test3) //mapexample[one:php two:golang three:java]

	// 2. 嵌套Map
	language := make(map[string]map[string]string)
	language["php"] = make(map[string]string, 2)
	language["php"]["id"] = "1"
	language["php"]["desc"] = "php是世界上最美的语言"
	language["golang"] = make(map[string]string, 2)
	language["golang"]["id"] = "2"
	language["golang"]["desc"] = "golang抗并发非常good"

	fmt.Println(language) //mapexample[php:mapexample[id:1 desc:php是世界上最美的语言] golang:mapexample[id:2 desc:golang抗并发非常good]]

	//增删改查
	val, key := language["php"] //查找是否有php这个子元素
	if key {
		fmt.Printf("%v", val)
	} else {
		fmt.Printf("no")
	}

	language["php"]["id"] = "3"         //修改了php子元素的id值
	language["php"]["nickname"] = "啪啪啪" //增加php元素里的nickname值
	delete(language, "php")             //删除了php子元素

	fmt.Println(language)
}

func printMap(x map[string]string) {
	fmt.Printf("len=%d  map=%v\n", len(x), x)

	var value = x["one"]
	fmt.Printf("map指定取值:", value)
}
