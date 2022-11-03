package structs

import (
	"fmt"
	"time"
)

//type Book struct {
//	id      int
//	title   string
//	author  string
//	subject string
//	alias   string
//	date    time.Time
//}

type Book struct {
	id        int    `json:"id"`
	title     string `json:"title"`
	author    string `json:"author"`
	subject   string `json:"subject"`
	alias     string `json:"alias,omitempty"` // omitempty: 不存在,则忽略,存在则返回
	status    bool
	price     float64
	create_at time.Time `json:"-"` // -:永久忽略
	update_at time.Time `json:"-"` // -:永久忽略
}

func BookStructExample() {
	/* 声明 Book1, Book2  为 Books 类型 */
	book := Book{}
	book.id = 1
	book.title = "Go语言电子书"
	book.author = "李俊超"
	book.price = 99.00
	book.status = true
	book.create_at = time.Now()
	book.update_at = time.Now()

	/* 打印 Book 信息 */
	marshalToJson(book)
}

func marshalToJson(book Book) {
	fmt.Println("book对象值:", book)

	fmt.Println("book 日期时间戳:", book.create_at)
}
